// +build test

package controllers_test

import (
  "strings"
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/jakemmarsh/viral-api/controllers"
  "github.com/jakemmarsh/viral-api/test_utils"
)

func TestGetPingHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "/v1/ping", nil)

  if err != nil {
    t.Fatal(err)
  }

  responseRecorder := httptest.NewRecorder()
  handler := http.HandlerFunc(controllers.GetPingHandler)

  handler.ServeHTTP(responseRecorder, req)

  testutils.Equals(t, responseRecorder.Code, http.StatusOK)

  expectedBody := `{"message":"pong"}`
  testutils.Equals(t, strings.TrimSpace(responseRecorder.Body.String()), expectedBody)
}
