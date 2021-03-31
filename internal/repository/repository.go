package repository

import (
	"database/sql"

	"github.com/mjhjj/Tyrannosaurus/internal/domain"
)

// Workers ...
type Places interface {
	SelectAllPlaces() ([]domain.Place, error)
	Insert(place domain.Place) error
}

// Repositories ...
type Repositories struct {
	Places
}

// NewRepository ...
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Places: NewPlaceSQLite(db),
	}
}
