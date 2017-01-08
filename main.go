package main

import (
	"fmt"
	"log"
	"net/http"
  "encoding/json"
  "github.com/gorilla/mux"
	"github.com/jakemmarsh/viral-api/utils"
)

type Ping struct {
	Message string `json:"message"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Ping{Message: "pong"})
}

func main() {
	utils.LoadConfiguration()

	r := mux.NewRouter()

	r.HandleFunc("/ping", pingHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", utils.GetEnvVariable("PORT", "3000")), r))
}
