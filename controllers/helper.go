package controllers

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
)

var IsProduction *bool

func InternalServerError() *views.View {
	return NewStaticController().InternalServerError
}

func ForbiddenError() *views.View {
	return NewStaticController().ForbiddenError
}

func GetIP(r *http.Request) {
	ip := r.RemoteAddr
	xforward := r.Header.Get("X-Forwarded-For")
	ipAddr := fmt.Sprintf("IP: %s", ip)
	forwardFor := fmt.Sprintf("X-Forwarded-For : %s", xforward)
	file, err := os.OpenFile("visitors.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening the file")
	}
	_, err = file.Write([]byte(ipAddr + "\n" + forwardFor + "\n"))
	if err != nil {
		fmt.Println("Error writing to the file")
	}
}

func WrapIPHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call any pre handler functions here
		go GetIP(r)
		f(w, r)
	}
}

func createJWT(w http.ResponseWriter, admin *models.Admin) error {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return nil
}

func validateJWT(r *http.Request) bool {
	token, err := r.Cookie("token")
	if err != nil {
		return false
	}

	tokenString := token.Value

	claims := &Claims{}

	result, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !result.Valid {
		return false
	}

	return true
}

func refreshJWT(w http.ResponseWriter, r *http.Request) error {
	// TODO :
	return nil
}

func (admin *Admin) SignoutJWT(w http.ResponseWriter, r *http.Request) {
	internalServerError := InternalServerError()
	if !validateJWT(r) {
		internalServerError.Render(w, nil)
	}
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}

func parseForm(r *http.Request, f interface{}) (string, error) {
	var text string
	encoding := r.Header.Get("Content-Type")
	if encoding == "application/x-www-form-urlencoded" {
		r.ParseForm()
	} else {
		err := r.ParseMultipartForm(10000000000)
		if err != nil {
			return "", err
		}
		files := r.MultipartForm.File["File"]
		if len(files) > 0 {
			for _, file := range files {
				fileContent, err := file.Open()
				if err != nil {
					fmt.Println("Error opening the file")
					return "", err
				}
				byteContainer, err := ioutil.ReadAll(fileContent)
				if err != nil {
					fmt.Println("Error reading the file")
					return "", err
				}
				defer fileContent.Close()
				text = string(byteContainer)
			}
		} else {
			fmt.Println("HERE")
			r.ParseForm()
		}
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	fmt.Println(r.PostForm)
	err := decoder.Decode(f, r.PostForm)

	fmt.Println(r.PostForm)

	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR NUMBER 3")
		return "", err
	}

	return text, nil
}

func ScriptFetcher(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Split(r.URL.Path, "/")[4:][0]
	var filePath string
	if *IsProduction == true {
		filePath = filepath.Join("./views", "js", "prod", fileName)
	} else {
		filePath = filepath.Join("./views", "js", "develop", fileName)
	}

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := template.New("").Parse(
		fmt.Sprintf("<script>%s</script>", string(contents)),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var (
		result = buf.Bytes()
		start  = len([]byte("<script>"))
		end    = len(result) - len([]byte("</script>"))
	)
	if _, err := w.Write(result[start:end]); err != nil {
		fmt.Printf("error writing response: %s\n", err.Error())
	}
}

func CheckProduction() *bool {
	return flag.Bool("prod", false, "Specify if server in production")
}
