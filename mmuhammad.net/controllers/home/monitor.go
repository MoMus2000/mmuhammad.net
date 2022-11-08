package home

import (
	"encoding/json"
	"fmt"
	"mustafa_m/controllers"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
)

type Monitor struct {
	monitorService *models.MonitorService
	MonitorPage    *views.View
}

func NewMonitorController(monitorService *models.MonitorService) *Monitor {
	return &Monitor{
		monitorService: monitorService,
		MonitorPage:    views.NewView("bootstrap", "monitoring/marketWatch.gohtml"),
	}
}

func (monitor *Monitor) GetUsdToPkr(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.UsdToPkr()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetSteelRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.MetalPrices()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetOilRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.OilPrices()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetBasementRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.BasementRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetApartmentRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.ApartmentRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetSpyRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.SPYRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func AddMonitorRoutes(r *mux.Router, monC *Monitor) {
	r.Handle("/market", monC.MonitorPage).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/usopen", monC.GetUsdToPkr).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/steel", monC.GetSteelRates).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/oil", monC.GetOilRates).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/basement", monC.GetBasementRates).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/apartment", monC.GetApartmentRates).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/spy", monC.GetSpyRates).Methods("GET")
}
