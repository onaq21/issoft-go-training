package db

import (
	"fmt"
	"database/sql"
)

func CreateTable(db *sql.DB) error {
	createQuery := `
	CREATE TABLE IF NOT EXISTS Users (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Email TEXT NOT NULL UNIQUE,
		Hash TEXT NOT NULL,
		Name TEXT NOT NULL UNIQUE,
		IsActive INTEGER NOT NULL
	)`

	if _, err := db.Exec(createQuery); err != nil {
		return fmt.Errorf("Create table error: %s", err)
	}

	return nil
}