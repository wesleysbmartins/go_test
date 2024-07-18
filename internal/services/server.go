package server

import (
	"fmt"
	"go_tests/internal/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ServerApi struct{}

type IServerAPI interface {
	Run()
}

func (s *ServerApi) Run() {
	router := mux.NewRouter()
	router.Use()
	routes.Routes(router)
	http.Handle("/", router)

	port := "8003"
	allowedOrigins := "*"

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{allowedOrigins}))(router)

	fmt.Println(fmt.Printf("SERVER LISTENNING ON PORT: %s", port))
	panic(http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler))
}
