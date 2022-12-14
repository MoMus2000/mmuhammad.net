package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/models/model_auth"
	"sms.mmuhammad.net/views"
)

type Login struct {
	LoginPage    *views.View
	LoginService *model_auth.AuthService
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginPageController(LoginService *model_auth.AuthService) *Login {
	return &Login{
		views.NewView("landingPage", "/auth/login.gohtml"),
		LoginService,
	}
}

func (login *Login) SubmitLoginPage(w http.ResponseWriter, r *http.Request) {
	lf := LoginForm{}
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(payload, &lf)
	_, err = login.LoginService.ByEmail(lf.Email, lf.Password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Cookie needs to be set before writing the headers
	err = createJWT(w, lf.Email)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
}

func (login *Login) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	if ValidateJWT(r) {
		http.Redirect(w, r, "/sms", http.StatusFound)
	}
	login.LoginPage.Render(w, nil)
}

func AddLoginRoutes(r *mux.Router, landC *Login) {
	r.HandleFunc("/login", landC.GetLoginPage).Methods("GET")
	r.HandleFunc("/login", landC.SubmitLoginPage).Methods("POST")
}
