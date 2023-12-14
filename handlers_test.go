package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRequest(t *testing.T, method string, path string) *http.Request {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestIndexHandlerOKResponse(t *testing.T) {
	// Given
	req := createRequest(t, "GET", "/")
	rr := httptest.NewRecorder()

	// When
	indexHandler(rr, req)

	// Then
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestIndexHandler404Response(t *testing.T) {
	// Given
	req := createRequest(t, "GET", "/notavalidpath")
	rr := httptest.NewRecorder()

	// Call the handler function with the fake request and response
	indexHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestIndexHandlerMethodNotAllowed(t *testing.T) {
	// Given
	req := createRequest(t, "PUT", "/")
	rr := httptest.NewRecorder()

	// Call the handler function with the fake request and response
	indexHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
