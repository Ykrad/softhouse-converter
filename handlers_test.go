package main

import (
	"net/http"
	"net/http/httptest"
	"os"
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

func TestConvertToXMLForPersonWithAddressAndPhone(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithaddressandphone")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<person>" +
		"<firstname>Carl Gustaf</firstname>" +
		"<lastname>Bernadotte</lastname>" +
		"<phone>" +
		"<mobile>0768-101801</mobile>" +
		"<landline>08-101801</landline>" +
		"</phone>" +
		"<address>" +
		"<street>Drottningholms slott</street>" +
		"<city>Stockholm</city>" +
		"<areaCode>10001</areaCode>" +
		"</address>" +
		"</person>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: %s\nBut was:%s", expected, xml)
	}
}

func TestConvertToXMLWorksForPersonWithoutAddress(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithoutaddress")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<person>" +
		"<firstname>Carl Gustaf</firstname>" +
		"<lastname>Bernadotte</lastname>" +
		"<phone>" +
		"<mobile>0768-101801</mobile>" +
		"<landline>08-101801</landline>" +
		"</phone>" +
		"</person>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: %s\nBut was:%s", expected, xml)
	}
}

func TestConvertToXMLWorksForPersonWithoutPhone(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithoutphone")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<person>" +
		"<firstname>Carl Gustaf</firstname>" +
		"<lastname>Bernadotte</lastname>" +
		"<address>" +
		"<street>Drottningholms slott</street>" +
		"<city>Stockholm</city>" +
		"<areaCode>10001</areaCode>" +
		"</address>" +
		"</person>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: %s\nBut was:%s", expected, xml)
	}
}

func TestConvertToXMLWorksForPersonWithAddressWithoutAreaCode(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithaddresswithoutareacode")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<person>" +
		"<firstname>Barack</firstname>" +
		"<lastname>Obama</lastname>" +
		"<address>" +
		"<street>1600 Pennsylvania Avenue</street>" +
		"<city>Washington, D.C</city>" +
		"</address>" +
		"</person>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: %s\nBut was:%s", expected, xml)
	}
}
