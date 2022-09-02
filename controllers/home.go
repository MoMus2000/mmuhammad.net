package controllers

import (
	"mustafa_m/views"
	"net/http"
)

type Home struct {
	HomePage      *views.View
	HomePageAdmin *views.View
}

func NewHomeController() *Home {
	return &Home{
		HomePage:      views.NewView("bootstrap", "home/home.gohtml"),
		HomePageAdmin: views.NewView("bootstrap", "home/home.gohtml"),
	}
}

func (home *Home) GetHomePage(w http.ResponseWriter, r *http.Request) {
	if !validateJWT(r) {
		home.HomePage.Render(w, nil)
	} else {
		type Data struct {
			LoggedIn string
		}
		data := &Data{LoggedIn: "true"}
		home.HomePageAdmin.Render(w, data)
	}
}
