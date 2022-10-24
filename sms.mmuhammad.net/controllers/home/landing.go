package home

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"sms.mmuhammad.net/controllers/helper"
	"sms.mmuhammad.net/models/landing"
	"sms.mmuhammad.net/views"
)

type Landing struct {
	LandingPage    *views.View
	PrivacyPolicy  *views.View
	LandingService *landing.LandingService
}

type LandingContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Time    time.Time
}

func NewLandingPageController(ls *landing.LandingService) *Landing {
	return &Landing{
		LandingPage:    views.NewView("landingPage", "/home/landingPageContent.gohtml"),
		PrivacyPolicy:  views.NewView("privacy", "/home/privacy.gohtml"),
		LandingService: ls,
	}
}

func (landing *Landing) GetLandingPage(w http.ResponseWriter, r *http.Request) {
	landing.LandingPage.Render(w, nil)
}

func (landing *Landing) GetPrivacyPolicyPage(w http.ResponseWriter, r *http.Request) {
	landing.PrivacyPolicy.Render(w, nil)
}

func (landing *Landing) SubmitContactForm(w http.ResponseWriter, r *http.Request) {
	lcf := LandingContactForm{}
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(payload, &lcf)
	if err != nil {
		fmt.Println(err)
	}
	lcf.Time = helper.GetEstTime()

	err = landing.LandingService.SaveContactInfo(lcf)
	if err != nil {
		fmt.Println(err)
	}
}

func AddHomePageRoutes(r *mux.Router, landC *Landing) {
	r.HandleFunc("/", WrapIPHandler(landC.GetLandingPage)).Methods("GET")
	r.HandleFunc("/policy", landC.GetPrivacyPolicyPage).Methods("GET")
	r.HandleFunc("/api/v1/landing/contact", landC.SubmitContactForm).Methods("POST")
}
