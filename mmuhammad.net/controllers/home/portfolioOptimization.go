package home

import (
	"fmt"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
)

type PortfolioOptimization struct {
	PoPage *views.View
}

type PortfolioOptimizationForm struct {
	Tickers     string
	Amount      float32
	TimeHorizon string
}

func NewPortfolioOptimization() *PortfolioOptimization {
	return &PortfolioOptimization{
		PoPage: views.NewView("bootstrap", "portfolioOptimization/portfolioOptimization.gohtml"),
	}
}

func (po *PortfolioOptimization) SubmitPortfolioOptimizationValues(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Req got")
}

func AddPORoutes(r *mux.Router, poC *PortfolioOptimization) {
	r.Handle("/optimize", poC.PoPage)
	r.HandleFunc("/api/v1/optimize/crunch", poC.SubmitPortfolioOptimizationValues).Methods("POST")
}
