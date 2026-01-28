package main

import (
	"net/http"
	"log"
	"database/sql"
	_ "modernc.org/sqlite"
	"fmt"
)

func main() {
	db, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := Server{db}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", s.GetUsers)
	mux.HandleFunc("GET /users/{id}", s.GetUserByID)
	mux.HandleFunc("DELETE /users/{id}", s.DeleteUser)

	if err = http.ListenAndServe("localhost:5001", mux); err != nil {
		log.Fatalf("Listen and serve server error: %s", err)
	}
}

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "db/users.db")
	if err != nil {
		return nil, fmt.Errorf("Open db error: %s", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Ping db error: %s", err)
	}

	return db, nil
}