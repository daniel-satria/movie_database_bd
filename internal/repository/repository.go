package repository

import (
	"database/sql"
	"movie_database_be/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
