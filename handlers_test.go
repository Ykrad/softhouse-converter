package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRequest(t *testing.T, method string, path string) *http.Request {
	request, error := http.NewRequest(method, path, nil)
	if error != nil {
		t.Fatal(error)
	}
	return request
}

func TestIndexHandlerOKResponse(t *testing.T) {
	// Given
	request := createRequest(t, "GET", "/")
	responseRecorder := httptest.NewRecorder()

	// When
	indexHandler(responseRecorder, request)

	// Then
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestIndexHandler404Response(t *testing.T) {
	// Given
	request := createRequest(t, "GET", "/notavalidpath")
	responseRecorder := httptest.NewRecorder()

	// Call the handler function with the fake request and response
	indexHandler(responseRecorder, request)

	// Check the response status code
	if status := responseRecorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestIndexHandlerMethodNotAllowed(t *testing.T) {
	// Given
	request := createRequest(t, "PUT", "/")
	responseRecorder := httptest.NewRecorder()

	// Call the handler function with the fake request and response
	indexHandler(responseRecorder, request)

	// Check the response status code
	if status := responseRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
