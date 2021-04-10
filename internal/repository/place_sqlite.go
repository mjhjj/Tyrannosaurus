package repository

import (
	"database/sql"
	"errors"

	"github.com/mjhjj/Tyrannosaurus/internal/domain"
)

// placesQLite ...
type PlacesSQLite struct {
	db *sql.DB
}

// NewplacesQLite ...
func NewPlaceSQLite(db *sql.DB) *PlacesSQLite {
	return &PlacesSQLite{db: db}
}

// SelectAllPlaces select wall places
func (d *PlacesSQLite) SelectAllPlaces() ([]domain.Place, error) {
	var places []domain.Place
	query := "SELECT id, x, y, name, address, about, bio, link, sity, image, linkName FROM places;"
	rows, err := d.db.Query(query)
	if err != nil {
		return []domain.Place{}, err
	}

	var counter int
	for rows.Next() {
		var place domain.Place
		err := rows.Scan(&place.Id, &place.PositionX, &place.PositionY, &place.Name, &place.Address, &place.About, &place.Bio, &place.PanoramLink, &place.Sity, &place.Image, &place.LinkName)
		if err != nil {
			continue
		}
		places = append(places, place)
		counter++
	}
	if counter < 1 {
		return []domain.Place{}, errors.New("Not found")
	}
	return places, err

}

// Insert new place
func (d *PlacesSQLite) Insert(place domain.Place) error {
	query := "insert into places (x, y, name, address, about, bio, link, sity, image, linkName) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	_, err := d.db.Exec(query, place.PositionX, place.PositionY, place.Name, place.Address, place.About, place.Bio, place.PanoramLink, place.Sity, place.Image, place.LinkName)

	return err
}
