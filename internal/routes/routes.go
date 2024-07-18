package routes

import (
	"go_tests/internal/controllers"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/calculator", controllers.CalculatorController).Methods("POST")
}
