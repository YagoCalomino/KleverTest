package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRequest(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"example": "response"}`)
	}))
	defer testServer.Close()

	response := sendRequest(testServer.URL)
	if response == nil {
		t.Error("Expected non-nil response, got nil")
	}

	expectedResponse := `{"example": "response"}`
	if string(response) != expectedResponse {
		t.Errorf("Expected response: %s, got: %s", expectedResponse, string(response))
	}
}
