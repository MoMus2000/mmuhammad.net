package terminal

import (
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/controllers/auth"
	"sms.mmuhammad.net/views"
)

type SmsTerminal struct {
	MainPage  *views.View
	Dashboard *views.View
	HelpPage  *views.View
}

func NewSmsTerminal() *SmsTerminal {
	return &SmsTerminal{
		views.NewView("smsLayout", "/sms/smsMain.gohtml"),
		views.NewView("smsLayout", "/sms/smsDashboard.gohtml"),
		views.NewView("smsLayout", "/sms/smsHelp.gohtml"),
	}
}

func (sms *SmsTerminal) GetSmsMainPage(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sms.MainPage.Render(w, nil)
}

func (sms *SmsTerminal) GetSmsDashboard(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sms.Dashboard.Render(w, nil)
}

func (sms *SmsTerminal) SmsTerminalSignOut(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.SignoutJWT(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (sms *SmsTerminal) GetSmsHelp(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sms.HelpPage.Render(w, nil)
}

func AddTerminalRoutes(r *mux.Router, smsC *SmsTerminal) {
	r.HandleFunc("/sms", smsC.GetSmsMainPage).Methods("GET")
	r.HandleFunc("/sms/dash", smsC.GetSmsDashboard).Methods("GET")
	r.HandleFunc("/sms/help", smsC.GetSmsHelp).Methods("GET")
	r.HandleFunc("/signout", smsC.SmsTerminalSignOut).Methods("GET")
}
