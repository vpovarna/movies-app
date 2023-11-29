package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 18081

type application struct {
	domain string
}

func main() {
	// set application config
	var app application

	// read from command line (flags)

	// connect to the database

	app.domain = "example.com"
	
	log.Printf("Starting application on port: %d", port)

	// start the web server
	mux := http.NewServeMux()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)

	if err != nil {
		log.Fatal("Unable to start the http server")
	}

}
