package main

import (
	"flag"
	"fmt"
	"log"
	"movies-app-be/internal/repository"
	"movies-app-be/internal/repository/dbrepo"
	"net/http"
)

const port = 18081

type application struct {
	DNS    string
	domain string
	DB     repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// read from command line (flags)
	flag.StringVar(&app.DNS, "dns", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "POSTGRES_CONNECTION_STRING")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.domain = "example.com"
	log.Printf("Starting application on port: %d", port)

	// start the web server
	routes := app.routes()
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), routes)
	if err != nil {
		log.Fatal("Unable to start the http server")
	}

}
