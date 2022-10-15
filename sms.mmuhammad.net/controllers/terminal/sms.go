package terminal

import (
	"net/http"

	"github.com/gorilla/mux"
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
	sms.MainPage.Render(w, nil)
}

func AddTerminalRoutes(r *mux.Router, smsC *SmsTerminal) {
	r.HandleFunc("/usr", smsC.GetSmsTerminal).Methods("GET")
}
