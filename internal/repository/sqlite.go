package repository

import (
	// sqlite3 driver
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// ping database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// create table workers if don`t exist
	query := "CREATE TABLE places (id INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,x TEXT  NOT NULL,y TEXT  NOT NULL,name TEXT  NOT NULL,address TEXT  NOT NULL,about TEXT  NOT NULL,bio TEXT  NOT NULL,link TEXT  NOT NULL);"
	db.Exec(query)

	return db, nil
}
