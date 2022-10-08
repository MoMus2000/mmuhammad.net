package fmb

import (
	"fmt"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"github.com/xuri/excelize/v2"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createJWTFmb(w http.ResponseWriter, fmb *models.Fmb) error {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: fmb.Email,
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
		Name:    "FMB",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return nil
}

func (tw *Twilio) SignoutJWTFmb(w http.ResponseWriter, r *http.Request) {
	if !validateJWTFmb(r) {
		controllers.InternalServerError().Render(w, nil)
		return
	}

	c := http.Cookie{
		Name:   "FMB",
		MaxAge: -1,
		Path:   "/",
	}

	http.SetCookie(w, &c)

	http.Redirect(w, r, "/fmb", http.StatusFound)
}

func validateJWTFmb(r *http.Request) bool {
	token, err := r.Cookie("FMB")
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

func parseExcelForm(r *http.Request, f interface{}) (*excelize.File, error) {
	err := r.ParseMultipartForm(10000000000)
	if err != nil {
		return nil, err
	}
	file, _, err := r.FormFile("File")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	excelFile, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	fmt.Println(r.PostForm)
	err = decoder.Decode(f, r.PostForm)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return excelFile, nil
}
