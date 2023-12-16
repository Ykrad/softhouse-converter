package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
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

func convertToXml(file multipart.File) (string, error) {
	// Create a Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	person := Person{}

	// Process each line
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "|")

		if splitLine[0] == "P" {
			person = Person{
				firstName: splitLine[1],
				lastName:  splitLine[2],
			}

			// Get the lines related to this person
			for i := 0; i < 2; i++ {
				if !scanner.Scan() {
					break
				}
				splitLine = strings.Split(scanner.Text(), "|")
				parseSubLine(&person, splitLine)
			}
		}
	}

	if error := scanner.Err(); error != nil {
		return "", error
	}

	return person.toXML(), nil
}

func parseSubLine(person *Person, splitLine []string) {
	if splitLine[0] == "T" {
		person.phone = Phone{
			mobile:   splitLine[1],
			landline: splitLine[2],
		}
		person.phoneInitialized = true
	} else if splitLine[0] == "A" {
		person.address = Address{
			street: splitLine[1],
			city:   splitLine[2],
		}

		if len(splitLine) == 4 {
			person.address.areaCode = splitLine[3]
		}

		person.addressInitialized = true
	}
}
