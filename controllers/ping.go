package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/jakemmarsh/viral-api/models"
)

func GetPingHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(models.Ping{Message: "pong"})
}
