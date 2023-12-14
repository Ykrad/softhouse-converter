package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register the handler function for a specific route
	http.HandleFunc("/", indexHandler)

	// Set the port for the server to listen on
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)

	// Define command line arguments
	https := flag.Bool("https", false, "Sets the webserver in https only mode")
	darkMode := flag.Bool("dark-mode-off", false, "Sets css styling to normal non darkmode")

	// Parse command line arguments
	flag.Parse()

	darkModeOff = *darkMode

	if *https {
		fmt.Printf("Starting the server in https mode. Go to https://localhost:%d\n", port)

		// Specify paths to TLS certificate and private key files
		certFile := "./security/cert.pem"
		keyFile := "./security/key.pem"

		// Start the web server with HTTPS and listen on the specified port
		error := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil)
		if error != nil {
			fmt.Println("Error starting the server:", error)
		}
	} else {
		fmt.Printf("Starting the server in http mode. Go to http://localhost:%d\n", port)

		// Start the web server and listen on the specified port
		error := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if error != nil {
			fmt.Println("Error starting the server:", error)
		}
	}
}
