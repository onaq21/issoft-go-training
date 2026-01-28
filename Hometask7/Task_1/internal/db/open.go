package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"fmt"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "users.db")
	if err != nil {
		return nil, fmt.Errorf("Open db error: %s", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Ping db error: %s", err)
	}

	return db, nil
}