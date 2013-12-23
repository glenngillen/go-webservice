package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  )

func init() {
}

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
  request, _ := http.NewRequest("GET", "/", nil)
  response := httptest.NewRecorder()

  viewHandler(response, request, "ada")
}
