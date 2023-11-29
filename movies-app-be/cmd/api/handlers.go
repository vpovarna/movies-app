package main

import (
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.GetAllMovies()
	if err != nil {
		app.errorJson(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	// read a json payload (username and password)

	// validate user against DB

	// check password

	// create jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	// generate token
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	log.Println(tokens.Token)
	w.Write([]byte(tokens.Token))
}
