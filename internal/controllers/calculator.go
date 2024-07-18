package controllers

import (
	"encoding/json"
	"go_tests/internal/usecases"
	"net/http"
)

func CalculatorController(w http.ResponseWriter, r *http.Request) {
	dto := &usecases.CalculatorDTO{}

	json.NewDecoder(r.Body).Decode(&dto)

	result, err := usecases.Calculator(*dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
