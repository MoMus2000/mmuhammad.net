package terminal

import (
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/controllers/auth"
	"sms.mmuhammad.net/views"
)

type SmsTerminal struct {
	MainPage *views.View
}

func NewSmsTerminal() *SmsTerminal {
	return &SmsTerminal{
		views.NewView("smsLayout", "/sms/smsMain.gohtml"),
	}
}

func (sms *SmsTerminal) GetSmsTerminal(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sms.MainPage.Render(w, nil)
}

func (sms *SmsTerminal) SmsTerminalSignOut(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.SignoutJWT(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func AddTerminalRoutes(r *mux.Router, smsC *SmsTerminal) {
	r.HandleFunc("/usr", smsC.GetSmsTerminal).Methods("GET")
	r.HandleFunc("/signout", smsC.SmsTerminalSignOut).Methods("GET")
}
