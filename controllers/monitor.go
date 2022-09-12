package controllers

import (
	"encoding/json"
	"fmt"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
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
	internalServerError := InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetSteelRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.MetalPrices()
	internalServerError := InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (monitor *Monitor) GetOilRates(w http.ResponseWriter, r *http.Request) {
	monitors, err := monitor.monitorService.OilPrices()
	internalServerError := InternalServerError()
	if err != nil {
		internalServerError.Render(w, nil)
	}
	fmt.Println(monitors)
	jsonEncoding, err := json.Marshal(monitors)
	fmt.Fprintln(w, string(jsonEncoding))
}
