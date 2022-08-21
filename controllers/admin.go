package controllers

import (
	"fmt"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/schema"
)

type Admin struct {
	LoginPage *views.View
}

func NewAdminController() *Admin {
	return &Admin{
		LoginPage: views.NewView("bootstrap", "admin/login.gohtml"),
	}
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (admin *Admin) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request recieved")
	form := LoginForm{}
	parseForm(r, &form)
	fmt.Println(form)
}

func parseForm(r *http.Request, f *LoginForm) {
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
