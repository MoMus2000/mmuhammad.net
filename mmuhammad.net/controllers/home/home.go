package home

import (
	"mustafa_m/controllers"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
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
	if !controllers.ValidateJWT(r) {
		home.HomePage.Render(w, nil)
	} else {
		data := &views.Data{LoggedIn: "true"}
		home.HomePageAdmin.Render(w, data)
	}
}

func AddHomeRoutes(r *mux.Router, homeC *Home) {
	r.HandleFunc("/", controllers.WrapIPHandler(homeC.GetHomePage)).Methods("GET")
}
