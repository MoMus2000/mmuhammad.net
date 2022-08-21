package controllers

import (
	"fmt"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
)

type Admin struct {
	LoginPage    *views.View
	AdminService *models.AdminService
	PostService  *models.PostService
	BlogForm     *views.View
}

func NewAdminController(adminService *models.AdminService, ps *models.PostService) *Admin {
	return &Admin{
		LoginPage:    views.NewView("bootstrap", "admin/login.gohtml"),
		AdminService: adminService,
		PostService:  ps,
		BlogForm:     views.NewView("bootstrap", "admin/blogForm.gohtml"),
	}
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type BlogForm struct {
	Topic   string `schema:"Topic"`
	Summary string `schema:"Summary"`
	Content string `schema:"Content"`
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (admin *Admin) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request recieved")
	form := LoginForm{}
	parseForm(r, &form)
	fmt.Println(form)
	adminTemp := models.Admin{Email: form.Email, Password: form.Password}
	result, err := admin.AdminService.ByEmail(&adminTemp)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(result)

	createJWT(w, &adminTemp)

	http.Redirect(w, r, "/admin/create", http.StatusFound)
}

func (admin *Admin) SubmitBlogPost(w http.ResponseWriter, r *http.Request) {
	form := BlogForm{}
	parseForm(r, &form)
	post := models.Post{Topic: form.Topic, Content: form.Content, Summary: form.Summary}
	err := admin.PostService.Create(&post)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin/create", http.StatusFound)
}

func (admin *Admin) GetBlogForm(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	admin.BlogForm.Render(w, nil)
}

func createJWT(w http.ResponseWriter, admin *models.Admin) {
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
		panic(err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return
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

func refreshJWT() {
	//:TODO
}

func parseForm(r *http.Request, f interface{}) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	decoder := schema.NewDecoder()

	err = decoder.Decode(f, r.PostForm)

	if err != nil {
		panic(err)
	}

	return
}
