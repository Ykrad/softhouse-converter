package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

// Page is a struct to represent the data that will be rendered in the HTML template
type Page struct {
	Title               string
	UploadedFileContent string
	FileContentInXML    string
	DarkModeOff         bool
}

// Handles path "/download"
func downloadHandler(writer http.ResponseWriter, request *http.Request) {
	// Set the Content-Disposition header to trigger a download prompt
	writer.Header().Set("Content-Disposition", "attachment; filename=converted.xml")
	// Set the Content-Type header to indicate the file type (text/plain in this case)
	writer.Header().Set("Content-Type", "text/plain")

	// Write the text content to the response writer
	io.WriteString(writer, currentXML)
}

// Handles path "/"
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	// Validate
	if request.URL.Path != "/" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" && request.Method != "POST" {
		http.Error(writer, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	// Initialize page struct
	page := Page{
		Title:               "PSV to XML converter",
		UploadedFileContent: "",
		FileContentInXML:    "",
		DarkModeOff:         darkModeOff,
	}

	if request.Method == http.MethodPost {
		// Parse the form data
		error := request.ParseMultipartForm(10 << 20) // 10 MB limit
		if error != nil {
			http.Error(writer, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the file from the form
		file, handler, error := request.FormFile("file")
		if error != nil {
			http.Error(writer, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Println("Received file: ", handler.Filename)

		// Read the content of the uploaded file
		content, error := io.ReadAll(file)
		if error != nil {
			http.Error(writer, "Unable to read file content", http.StatusInternalServerError)
			return
		}

		// Add content to page template data
		page.UploadedFileContent = string(content)
		page.FileContentInXML, error = convertToXml(bytes.NewReader(content))
		if error != nil {
			http.Error(writer, "Unable to convert file content to XML", http.StatusInternalServerError)
			return
		}

		currentXML = page.FileContentInXML
	}

	// Parse the HTML template file
	indexTemplate, error := template.ParseFiles("./index.html")
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the data and write the result to the response writer
	error = indexTemplate.Execute(writer, page)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
	}
}

type People struct {
	XMLName xml.Name `xml:"people"`
	People  []Person `xml:"person"`
}

func convertToXml(reader io.Reader) (string, error) {
	// Create a Scanner to read the file line by line
	scanner := bufio.NewScanner(reader)
	people := People{
		People: []Person{},
	}

	// Get all lines
	lines := [][]string{}
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "|")
		lines = append(lines, splitLine)
	}

	// Parse all lines
	for i := 0; i < len(lines); i++ {
		if lines[i][0] == "P" {
			personLines, _ := getUnitLines(i, lines, personValidation)
			people.People = append(people.People, parsePerson(personLines))
		}
	}

	if error := scanner.Err(); error != nil {
		return "", error
	}

	xmlBytes, _ := xml.MarshalIndent(people, "", "    ")
	return string(xmlBytes), nil
}

func getUnitLines(index int, lines [][]string, validation func(string) bool) ([][]string, int) {
	unitLines := [][]string{}
	stoppedAtIndex := len(lines)
	for i := index; i < len(lines); i++ {
		unitLines = append(unitLines, lines[i])

		if i+1 != len(lines) && validation(lines[i+1][0]) {
			stoppedAtIndex = i
			break
		}
	}

	return unitLines, stoppedAtIndex
}
