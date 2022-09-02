package controllers

import (
	"fmt"
	"io/ioutil"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
)

type Admin struct {
	LoginPage       *views.View
	AdminService    *models.AdminService
	PostService     *models.PostService
	CategoryService *models.CategoryService
	BlogForm        *views.View
	DeleteForm      *views.View
	EditForm        *views.View
	CategoryForm    *views.View
}

func NewAdminController(adminService *models.AdminService, ps *models.PostService,
	cs *models.CategoryService) *Admin {
	return &Admin{
		LoginPage:       views.NewView("bootstrap", "admin/login.gohtml"),
		AdminService:    adminService,
		PostService:     ps,
		CategoryService: cs,
		BlogForm:        views.NewView("bootstrap", "admin/blogForm.gohtml"),
		DeleteForm:      views.NewView("bootstrap", "admin/deleteForm.gohtml"),
		EditForm:        views.NewView("bootstrap", "admin/editForm.gohtml"),
		CategoryForm:    views.NewView("bootstrap", "admin/categoryForm.gohtml"),
	}
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type BlogForm struct {
	Topic      string `schema:"Topic"`
	Summary    string `schema:"Summary"`
	Imgur_URL  string `schema:"Imgur"`
	Content    string
	CategoryId string `schema:"CID"`
}

type DeleteForm struct {
	Id string `schema:"Id"`
}

type EditForm struct {
	Id        string `schema:"ID"`
	Topic     string `schema:"Topic"`
	Summary   string `schema:"Summary"`
	Imgur_URL string `schema:"Imgur"`
}

type CategoryForm struct {
	Category  string `schema:"Cat"`
	Summary   string `schema:"Summary"`
	Imgur_URL string `schema:"Imgur"`
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (admin *Admin) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request recieved")
	internalServerError := InternalServerError()
	form := LoginForm{}
	parseForm(r, &form)
	fmt.Println(form)
	adminTemp := models.Admin{Email: form.Email, Password: form.Password}
	result, err := admin.AdminService.ByEmail(&adminTemp)
	if err != nil {
		fmt.Println(err)
		internalServerError.Render(w, nil)
	}
	fmt.Println(result)

	createJWT(w, &adminTemp)

	http.Redirect(w, r, "/admin/create", http.StatusFound)
}

func (admin *Admin) SubmitBlogPost(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	internalServerError := InternalServerError()
	form := BlogForm{}
	content, err := parseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	form.Content = content
	// TODO: Change to take in the IMGUR URL and the uploaded file
	// Save the file content into the database
	post := models.Post{Topic: form.Topic, Content: form.Content, Summary: form.Summary,
		Imgur_URL: form.Imgur_URL, Date: time.Now().String(), CategoryId: form.CategoryId}
	err = admin.PostService.CreatePost(&post)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	http.Redirect(w, r, "/admin/create", http.StatusFound)
}

func (admin *Admin) SubmitDeleteRequest(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	form := DeleteForm{}
	internalServerError := InternalServerError()
	_, err := parseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	fmt.Println(form)
	idToUint, err := strconv.ParseUint(form.Id, 0, 64)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	err = admin.PostService.DeletePost(uint(idToUint))
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	http.Redirect(w, r, "/admin/delete", http.StatusFound)
}

func (admin *Admin) SubmitEditRequest(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	form := EditForm{}
	internalServerError := InternalServerError()
	_, err := parseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
		return
	}
	post := models.Post{Topic: form.Topic, Summary: form.Summary,
		Imgur_URL: form.Imgur_URL}
	err = admin.AdminService.UpdateChangesFromEdit(&post, form.Id)
	if err != nil {
		fmt.Println("Error I got back ", err)
		internalServerError.Render(w, nil)
		return
	}
	http.Redirect(w, r, "/admin/edit", http.StatusFound)
}

func (admin *Admin) SubmitCategoryFrom(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	internalServerError := InternalServerError()
	form := CategoryForm{}
	_, err := parseForm(r, &form)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	cat := models.Category{CategoryName: form.Category, Imgur_URL: form.Imgur_URL,
		CategorySummary: form.Summary, CreationDate: time.Now().String()}
	err = admin.CategoryService.Create(&cat)
	if err != nil {
		internalServerError.Render(w, nil)
	}
	http.Redirect(w, r, "/admin/category", http.StatusFound)
}

func (admin *Admin) GetBlogForm(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	admin.BlogForm.Render(w, nil)
}

func (admin *Admin) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	if validateJWT(r) {
		http.Redirect(w, r, "/admin/create", http.StatusFound)
		return
	}
	admin.LoginPage.Render(w, nil)
}

func (admin *Admin) GetDeletePage(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	admin.DeleteForm.Render(w, nil)
}

func (admin *Admin) GetEditPage(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}
	admin.EditForm.Render(w, nil)
}

func (admin *Admin) GetCategoryPage(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		ForbiddenError().Render(w, nil)
		return
	}

	admin.CategoryForm.Render(w, nil)
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
			fmt.Println(err)
			return "", err
		}
		files := r.MultipartForm.File["File"]
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
	}

	decoder := schema.NewDecoder()

	err := decoder.Decode(f, r.PostForm)

	fmt.Println(r.PostForm)

	if err != nil {
		return "", err
	}

	return text, nil
}
