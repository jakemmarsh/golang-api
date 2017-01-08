package main

import (
	"github.com/gorilla/mux"
	"github.com/jakemmarsh/viral-api/controllers"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/ping", controllers.GetPingHandler).Methods("GET")
}
