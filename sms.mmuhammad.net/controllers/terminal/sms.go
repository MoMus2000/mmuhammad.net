package terminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/controllers/auth"
	"sms.mmuhammad.net/models/model_sms"
	"sms.mmuhammad.net/views"
)

type SmsTerminal struct {
	MainPage          *views.View
	Dashboard         *views.View
	HelpPage          *views.View
	SmsMetricsService *model_sms.SmsMetricService
}

func NewSmsTerminal(sms_service *model_sms.SmsMetricService) *SmsTerminal {
	return &SmsTerminal{
		views.NewView("smsLayout", "/sms/smsMain.gohtml"),
		views.NewView("smsLayout", "/sms/smsDashboard.gohtml"),
		views.NewView("smsLayout", "/sms/smsHelp.gohtml"),
		sms_service,
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

func (sms *SmsTerminal) GetTotalMsgSent(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username

	flask_payload := make(map[string]string)

	flask_payload["email"] = userEmail

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/sms/total_messages", "application/json", bytes.NewBuffer(jsonString))

	if err != nil {
		fmt.Println(err)
	}

	var j interface{}

	err = json.NewDecoder(resp.Body).Decode(&j)

	respPayload, err := json.Marshal(j)
	fmt.Fprintln(w, string(respPayload))
}

func (sms *SmsTerminal) GetBalance(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username

	flask_payload := make(map[string]string)

	flask_payload["email"] = userEmail

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/sms/balance", "application/json", bytes.NewBuffer(jsonString))

	if err != nil {
		fmt.Println(err)
	}

	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	payload, err := json.Marshal(j)
	fmt.Fprintln(w, string(payload))
}

func (sms *SmsTerminal) GetMsgSentToday(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username

	flask_payload := make(map[string]string)

	flask_payload["email"] = userEmail

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/sms/total_messages_today", "application/json", bytes.NewBuffer(jsonString))

	if err != nil {
		fmt.Println(err)
	}

	var j interface{}

	err = json.NewDecoder(resp.Body).Decode(&j)

	respPayload, err := json.Marshal(j)

	fmt.Fprintln(w, string(respPayload))
}

func (sms *SmsTerminal) GetTotalCost(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username

	flask_payload := make(map[string]string)

	flask_payload["email"] = userEmail

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/sms/total_cost", "application/json", bytes.NewBuffer(jsonString))

	if err != nil {
		fmt.Println(err)
	}

	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	respPayload, err := json.Marshal(j)
	fmt.Fprintln(w, string(respPayload))
}

func AddTerminalRoutes(r *mux.Router, smsC *SmsTerminal) {
	r.HandleFunc("/sms", smsC.GetSmsMainPage).Methods("GET")
	r.HandleFunc("/sms/dash", smsC.GetSmsDashboard).Methods("GET")
	r.HandleFunc("/sms/help", smsC.GetSmsHelp).Methods("GET")
	r.HandleFunc("/api/v1/sms/totalMsg", smsC.GetTotalMsgSent).Methods("GET")
	r.HandleFunc("/api/v1/sms/balance", smsC.GetBalance).Methods("GET")
	r.HandleFunc("/api/v1/sms/todayMsg", smsC.GetMsgSentToday).Methods("GET")
	r.HandleFunc("/api/v1/sms/singleSms", smsC.SendSingleSms).Methods("POST")
	r.HandleFunc("/api/v1/sms/bulkSms", smsC.SendBulkSms).Methods("POST")
	r.HandleFunc("/api/v1/sms/totalCost", smsC.GetTotalCost).Methods("GET")
	r.HandleFunc("/signout", smsC.SmsTerminalSignOut).Methods("GET")
}
