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
	expected := "<people>\n" +
		"    <person>\n" +
		"        <firstname>Carl Gustaf</firstname>\n" +
		"        <lastname>Bernadotte</lastname>\n" +
		"        <phone>\n" +
		"            <mobile>0768-101801</mobile>\n" +
		"            <landline>08-101801</landline>\n" +
		"        </phone>\n" +
		"        <address>\n" +
		"            <street>Drottningholms slott</street>\n" +
		"            <city>Stockholm</city>\n" +
		"            <areaCode>10001</areaCode>\n" +
		"        </address>\n" +
		"    </person>\n" +
		"</people>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: \n%s\nBut was: \n%s", expected, xml)
	}
}

func TestConvertToXMLWorksForPersonWithoutAddress(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithoutaddress")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<people>\n" +
		"    <person>\n" +
		"        <firstname>Carl Gustaf</firstname>\n" +
		"        <lastname>Bernadotte</lastname>\n" +
		"        <phone>\n" +
		"            <mobile>0768-101801</mobile>\n" +
		"            <landline>08-101801</landline>\n" +
		"        </phone>\n" +
		"    </person>\n" +
		"</people>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: \n%s\nBut was: \n%s", expected, xml)
	}
}

func TestConvertToXMLWorksForPersonWithoutPhone(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithoutphone")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<people>\n" +
		"    <person>\n" +
		"        <firstname>Carl Gustaf</firstname>\n" +
		"        <lastname>Bernadotte</lastname>\n" +
		"        <address>\n" +
		"            <street>Drottningholms slott</street>\n" +
		"            <city>Stockholm</city>\n" +
		"            <areaCode>10001</areaCode>\n" +
		"        </address>\n" +
		"    </person>\n" +
		"</people>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: \n%s\nBut was: \n%s", expected, xml)
	}
}

func TestConvertToXMLWorksForPersonWithAddressWithoutAreaCode(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithaddresswithoutareacode")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<people>\n" +
		"    <person>\n" +
		"        <firstname>Barack</firstname>\n" +
		"        <lastname>Obama</lastname>\n" +
		"        <address>\n" +
		"            <street>1600 Pennsylvania Avenue</street>\n" +
		"            <city>Washington, D.C</city>\n" +
		"        </address>\n" +
		"    </person>\n" +
		"</people>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: \n%s\nBut was: \n%s", expected, xml)
	}
}

func TestConvertToXMLWorksForTwoPersons(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/twopersons")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<people>\n" +
		"    <person>\n" +
		"        <firstname>Carl Gustaf</firstname>\n" +
		"        <lastname>Bernadotte</lastname>\n" +
		"        <phone>\n" +
		"            <mobile>0768-101801</mobile>\n" +
		"            <landline>08-101801</landline>\n" +
		"        </phone>\n" +
		"        <address>\n" +
		"            <street>Drottningholms slott</street>\n" +
		"            <city>Stockholm</city>\n" +
		"            <areaCode>10001</areaCode>\n" +
		"        </address>\n" +
		"    </person>\n" +
		"    <person>\n" +
		"        <firstname>Barack</firstname>\n" +
		"        <lastname>Obama</lastname>\n" +
		"        <address>\n" +
		"            <street>1600 Pennsylvania Avenue</street>\n" +
		"            <city>Washington, D.C</city>\n" +
		"        </address>\n" +
		"    </person>\n" +
		"</people>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: \n%s\nBut was: \n%s", expected, xml)
	}
}

func TestConvertToXMLForPersonWithFamily(t *testing.T) {
	// Given
	testdata, error := os.Open("testdata/personwithfamily")
	if error != nil {
		t.Errorf("Failed to read test data")
	}
	expected := "<people>\n" +
		"    <person>\n" +
		"        <firstname>Carl Gustaf</firstname>\n" +
		"        <lastname>Bernadotte</lastname>\n" +
		"        <phone>\n" +
		"            <mobile>0768-101801</mobile>\n" +
		"            <landline>08-101801</landline>\n" +
		"        </phone>\n" +
		"        <address>\n" +
		"            <street>Drottningholms slott</street>\n" +
		"            <city>Stockholm</city>\n" +
		"            <areaCode>10001</areaCode>\n" +
		"        </address>\n" +
		"        <family>\n" +
		"            <name>Victoria</name>\n" +
		"            <born>1977</born>\n" +
		"            <address>\n" +
		"                <street>Haga Slott</street>\n" +
		"                <city>Stockholm</city>\n" +
		"                <areaCode>10002</areaCode>\n" +
		"            </address>\n" +
		"        </family>\n" +
		"        <family>\n" +
		"            <name>Carl Philip</name>\n" +
		"            <born>1979</born>\n" +
		"            <phone>\n" +
		"                <mobile>0768-101802</mobile>\n" +
		"                <landline>08-101802</landline>\n" +
		"            </phone>\n" +
		"        </family>\n" +
		"    </person>\n" +
		"</people>"

	// When
	xml, error := convertToXml(testdata)
	if error != nil {
		t.Errorf("Failed to convert")
	}

	// Then
	if xml != expected {
		t.Errorf("Did not return correct string: \nExpected: \n%s\nBut was: \n%s", expected, xml)
	}
}
