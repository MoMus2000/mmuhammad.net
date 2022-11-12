package home

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mustafa_m/views"
	"net/http"

	"github.com/gorilla/mux"
)

type PortfolioOptimization struct {
	PoPage *views.View
}

type PortfolioOptimizationForm struct {
	Tickers     string `json:"tickerValue"`
	Amount      string `json:"amountValue"`
	TimeHorizon string `json:"timeHorizon"`
}

func NewPortfolioOptimization() *PortfolioOptimization {
	return &PortfolioOptimization{
		PoPage: views.NewView("bootstrap", "portfolioOptimization/portfolioOptimization.gohtml"),
	}
}

func (po *PortfolioOptimization) SubmitPortfolioOptimizationValues(w http.ResponseWriter, r *http.Request) {
	pof := PortfolioOptimizationForm{}
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(payload, &pof)
	fmt.Println(pof)

	flask_payload := make(map[string]string)

	flask_payload["Tickers"] = pof.Tickers
	flask_payload["Amount"] = pof.Amount
	flask_payload["TimeHorizon"] = pof.TimeHorizon

	jsonString, err := json.Marshal(flask_payload)

	resp, err := http.Post("http://localhost:3001/api/v1/optimize/crunch", "application/json", bytes.NewBuffer(jsonString))

	if resp.StatusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseString, err := io.ReadAll(resp.Body)

	fmt.Println(string(responseString))

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(string(responseString))
}

func AddPORoutes(r *mux.Router, poC *PortfolioOptimization) {
	r.Handle("/optimize", poC.PoPage)
	r.HandleFunc("/api/v1/optimize/crunch", poC.SubmitPortfolioOptimizationValues).Methods("POST")
}
