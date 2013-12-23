package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  )

func init() {
  p := &Page{Title: "foo" }
  p.save()
}

func TestHandleGetReturnsPage(t *testing.T) {
  b := "{\"title\":\"foo\"}"
  request, _ := http.NewRequest("GET", "/pages/foo", nil)
  response := httptest.NewRecorder()

  viewHandler(response, request, "foo")
  if response.Body.String() != b {
    t.Fatalf("Response body does not equal '%s'", response.Body.String())
  }
}
