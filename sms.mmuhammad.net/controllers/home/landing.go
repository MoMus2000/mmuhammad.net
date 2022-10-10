package home

import (
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/views"
)

type Landing struct {
	LandingPage *views.View
}

func NewLandingPageController() *Landing {
	return &Landing{
		LandingPage: views.NewView("landingPage", "/home/landingPageContent.gohtml"),
	}
}

func (landing *Landing) GetLandingPage(w http.ResponseWriter, r *http.Request) {
	landing.LandingPage.Render(w, nil)
}

func AddHomePageRoutes(r *mux.Router, landC *Landing) {
	r.HandleFunc("/", landC.GetLandingPage).Methods("GET")
	// Linking css files
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("views/layout/style/"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/",
		http.FileServer(http.Dir("views/layout/style/img/"))))
}
