package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
)

type Server struct {
	DB *sql.DB
}

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM Users ORDER BY ID`

	rows, err := s.DB.Query(query)
	if err != nil {
		http.Error(w, "Select db error: " + err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Email, &u.Hash, &u.Name, &u.IsActive); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
    }
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Json encode error: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM Users WHERE ID = ?`

	row := s.DB.QueryRow(query, r.PathValue("id"))

	var u User

	if err := row.Scan(&u.Id, &u.Email, &u.Hash, &u.Name, &u.IsActive); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Query error: " + err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, "Json encode error: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	query := `DELETE FROM Users WHERE ID = ?`

	res, err := s.DB.Exec(query, r.PathValue("id"))

	if err != nil {
		http.Error(w, "Delete query error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "RowsAffected error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}