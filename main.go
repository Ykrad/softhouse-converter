package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// Page is a struct to represent the data that will be rendered in the HTML template
type Page struct {
	Title               string
	UploadedFileContent string
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" && request.Method != "POST" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
		return
	}

	page := Page{
		Title:               "PSV to XML converter",
		UploadedFileContent: "",
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

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register the handler function for a specific route
	http.HandleFunc("/", handler)

	// Set the port for the server to listen on
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)

	// Define command line arguments
	https := flag.Bool("https", false, "Sets the webserver in https only mode")

	// Parse command line arguments
	flag.Parse()

	if *https {
		fmt.Printf("Starting the server in https mode. Go to https://localhost:%d", port)

		// Specify paths to TLS certificate and private key files
		certFile := "./security/cert.pem"
		keyFile := "./security/key.pem"

		// Start the web server with HTTPS and listen on the specified port
		error := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil)
		if error != nil {
			fmt.Println("Error starting the server:", error)
		}
	} else {
		fmt.Printf("Starting the server in http mode. Go to http://localhost:%d", port)

		// Start the web server and listen on the specified port
		error := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if error != nil {
			fmt.Println("Error starting the server:", error)
		}
	}
}
