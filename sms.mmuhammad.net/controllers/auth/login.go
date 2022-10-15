package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/views"
)

type Login struct {
	LoginPage *views.View
}

func NewLoginPageController() *Login {
	return &Login{
		views.NewView("landingPage", "/auth/login.gohtml"),
	}
}

func (login *Login) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	login.LoginPage.Render(w, nil)
}

func AddLoginRoutes(r *mux.Router, landC *Login) {
	r.HandleFunc("/login", landC.GetLoginPage).Methods("GET")
}
