package main

import (
	"encoding/json"
	"fmt"
	"movies-app-be/internal/models"
	"net/http"
	"time"
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

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-01-02", "1986-03-07")
	highlanderMovie := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		RunTime:     116,
		MPAARating:  "R",
		Description: "A very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlanderMovie)

	rd, _ = time.Parse("2006-01-02", "1981-06-12")

	rotla := models.Movie{
		ID:          2,
		Title:       "Riders of the last Ark",
		ReleaseDate: rd,
		RunTime:     115,
		MPAARating:  "PG13",
		Description: "Another very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, rotla)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
