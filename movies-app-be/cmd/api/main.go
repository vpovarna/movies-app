package main

import (
	"flag"
	"fmt"
	"log"
	"movies-app-be/internal/repository"
	"movies-app-be/internal/repository/dbrepo"
	"net/http"
	"time"
)

const port = 18081

type application struct {
	DNS          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	// set application config
	var app application

	// read from command line (flags)
	flag.StringVar(&app.DNS, "dns", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "POSTGRES_CONNECTION_STRING")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verySecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-Domain", "localhost", "cookie Domain")
	flag.StringVar(&app.Domain, "Domain", "example.com", "Domain")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  app.CookieDomain,
		CookiePath:    "/",
		CookieName:    "__Host-refresh-token",
	}

	log.Printf("Starting application on port: %d", port)

	// start the web server
	routes := app.routes()
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), routes)
	if err != nil {
		log.Fatal("Unable to start the http server")
	}

}
