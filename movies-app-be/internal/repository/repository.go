package repository

import (
	"database/sql"
	"movies-app-be/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetAllMovies() ([]*models.Movie, error)
}
