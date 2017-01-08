package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jakemmarsh/viral-api/utils"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func StartServer() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/v1").Subrouter()
	server := &Server{router}

	SetupRoutes(apiRouter)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", utils.GetEnvVariable("PORT", "3000")), server))
}

func (server *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		return
	}

	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, DELETE, PATCH")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Content-Type", "application/json")

	server.router.ServeHTTP(rw, req)
}
