package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// Page is a struct to represent the data that will be rendered in the HTML template
type Page struct {
	Title               string
	UploadedFileContent string
	DarkModeOff         bool
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
		DarkModeOff:         darkModeOff,
	}

	if request.Method == http.MethodPost {
		// Parse the form data
		err := request.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			http.Error(writer, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the file from the form
		file, handler, err := request.FormFile("file")
		if err != nil {
			http.Error(writer, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Println("Received file: ", handler.Filename)

		// Read the content of the uploaded file
		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(writer, "Unable to read file content", http.StatusInternalServerError)
			return
		}

		// Add content to page template data
		page.UploadedFileContent = string(content)
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
