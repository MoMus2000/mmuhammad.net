package terminal

import (
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
	f32, err := sms.SmsMetricsService.GetTotalMessages(userEmail)
	if err != nil {
		fmt.Println(err)
	}
	type Data struct {
		Data float32 `json:"Data"`
	}

	payload, err := json.Marshal(&Data{Data: f32})

	fmt.Fprintln(w, string(payload))
}

func (sms *SmsTerminal) GetTotalPriceSent(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username
	f32, err := sms.SmsMetricsService.GetTotalPrices(userEmail)
	if err != nil {
		fmt.Println(err)
	}

	type Data struct {
		Data float32 `json:"Data"`
	}

	payload, err := json.Marshal(&Data{Data: f32})

	fmt.Fprintln(w, string(payload))
}

func (sms *SmsTerminal) GetPriceSentToday(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username
	f32, err := sms.SmsMetricsService.GetTodayPrices(userEmail)
	if err != nil {
		fmt.Println(err)
	}

	type Data struct {
		Data float32 `json:"Data"`
	}

	payload, err := json.Marshal(&Data{Data: f32})

	fmt.Fprintln(w, string(payload))
}

func (sms *SmsTerminal) GetMsgSentToday(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateJWT(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := auth.GetEmailFromJwt(r).Username
	f32, err := sms.SmsMetricsService.GetTodayMessages(userEmail)
	if err != nil {
		fmt.Println(err)
	}

	type Data struct {
		Data float32 `json:"Data"`
	}

	payload, err := json.Marshal(&Data{Data: f32})

	fmt.Fprintln(w, string(payload))
}

func AddTerminalRoutes(r *mux.Router, smsC *SmsTerminal) {
	r.HandleFunc("/sms", smsC.GetSmsMainPage).Methods("GET")
	r.HandleFunc("/sms/dash", smsC.GetSmsDashboard).Methods("GET")
	r.HandleFunc("/sms/help", smsC.GetSmsHelp).Methods("GET")
	r.HandleFunc("/api/v1/sms/totalMsg", smsC.GetTotalMsgSent).Methods("GET")
	r.HandleFunc("/api/v1/sms/totalPrice", smsC.GetTotalPriceSent).Methods("GET")
	r.HandleFunc("/api/v1/sms/todayMsg", smsC.GetMsgSentToday).Methods("GET")
	r.HandleFunc("/api/v1/sms/todayPrice", smsC.GetPriceSentToday).Methods("GET")
	r.HandleFunc("/api/v1/sms/singleSms", smsC.SendSingleSms).Methods("POST")
	r.HandleFunc("/api/v1/sms/bulkSms", smsC.SendBulkSms).Methods("POST")
	r.HandleFunc("/signout", smsC.SmsTerminalSignOut).Methods("GET")
}
