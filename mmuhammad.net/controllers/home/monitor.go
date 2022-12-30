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

func (monitor *Monitor) GetSpyProbs(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.SPYRegimeProbs()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetCADHousingProbs(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.CADHousingRegime()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetDurhamApartment(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.DurhanApartmentRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetDurhamBasement(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.DurhamBasementRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetWindsorApartment(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.WindsorApartmentRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetWindsorBasement(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.WindsorBasementRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetStCatharinesApartment(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.StCatharinesApartmentRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetStCatharinesBasement(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.StCatharinesBasementRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetHamiltonApartment(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.HamiltonApartmentRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetHamiltonBasement(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.HamiltonBasementRates()
	internalServerError := controllers.InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
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
	r.HandleFunc("/api/v1/monitoring/spy/regime", monC.GetSpyProbs).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/cad_housing/regime", monC.GetCADHousingProbs).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/stcatharines/apartment", monC.GetStCatharinesApartment).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/stcatharines/basement", monC.GetStCatharinesBasement).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/durham/apartment", monC.GetDurhamApartment).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/durham/basement", monC.GetDurhamBasement).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/windsor/apartment", monC.GetWindsorApartment).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/windsor/basement", monC.GetWindsorBasement).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/hamilton/apartment", monC.GetHamiltonApartment).Methods("GET")
	r.HandleFunc("/api/v1/monitoring/hamilton/basement", monC.GetHamiltonBasement).Methods("GET")
}
