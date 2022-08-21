package controllers

import (
	"mustafa_m/views"
	"net/http"
)

type Admin struct {
	LoginPage views.View
}

func NewAdminController() *Admin {
	return &Admin{
		LoginPage: *views.NewView("bootstrap", "admin/login.gohtml"),
	}
}

func (admin *Admin) Login(w http.ResponseWriter, r *http.Request) {
	admin.LoginPage.Render(w, nil)
}
