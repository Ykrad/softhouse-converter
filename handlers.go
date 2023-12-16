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
	people := []Person{}

	// Get all lines
	lines := [][]string{}
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "|")
		lines = append(lines, splitLine)
	}

	// Parse all lines
	for i := 0; i < len(lines); i++ {
		if lines[i][0] == "P" {
			personLines := getPersonLines(i, lines)
			people = append(people, parsePerson(personLines))
		}
	}

	if error := scanner.Err(); error != nil {
		return "", error
	}

	personsXML := ""
	for _, person := range people {
		personsXML += person.toXML()
	}

	return fmt.Sprintf("<people>%s</people>", personsXML), nil
}

func getPersonLines(index int, lines [][]string) [][]string {
	personLines := [][]string{}
	for i := index; i < len(lines); i++ {
		personLines = append(personLines, lines[i])

		if i+1 != len(lines) && lines[i+1][0] == "P" {
			break
		}
	}

	return personLines
}

func parsePerson(splitLines [][]string) Person {
	person := Person{}
	for _, splitLine := range splitLines {
		if splitLine[0] == "P" {
			person.firstName = splitLine[1]
			person.lastName = splitLine[2]
		} else if splitLine[0] == "T" {
			phone := Phone{
				mobile:   splitLine[1],
				landline: splitLine[2],
			}
			person.phone = phone
			person.phoneInitialized = true
		} else if splitLine[0] == "A" {
			address := Address{
				street: splitLine[1],
				city:   splitLine[2],
			}

			if len(splitLine) == 4 {
				address.areaCode = splitLine[3]
			}

			person.address = address
			person.addressInitialized = true
		}
	}

	return person
}
