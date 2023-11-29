package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 18081

type application struct {
	DNS    string
	domain string
}

func main() {
	// set application config
	var app application

	// read from command line (flags)
	flag.StringVar(&app.DNS, "dns", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "POSTGRES_CONNECTION_STRING")
	flag.Parse()

	// connect to the database

	app.domain = "example.com"

	log.Printf("Starting application on port: %d", port)

	// start the web server
	routes := app.routes()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), routes)
	if err != nil {
		log.Fatal("Unable to start the http server")
	}

}
