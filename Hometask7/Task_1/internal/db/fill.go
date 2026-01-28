package db

import (
	"Task_1/internal/users"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

func FillTable(db *sql.DB) ([]string, error) {
	file, err := os.Open("data/persons.txt")
	if err != nil {
		return nil, fmt.Errorf("Open file error: %s", err)
	}
	defer file.Close()

	query := `
			INSERT INTO Users
			(Email, Hash, Name, IsActive)
			VALUES (?, ?, ?, 1)
	`

	cmd, err := db.Prepare(query)
	if err != nil {
    return nil, fmt.Errorf("Prepare query error: %s", err)
	}
	defer cmd.Close()

	errorList := make([]string, 0)
	scanner := bufio.NewScanner(file)
	lineNum := 1


	for scanner.Scan() {
		person, err := users.ParseUser(scanner.Text())
		if err != nil {
			errorList = append(errorList, fmt.Sprintf("%d %s: %s", lineNum, scanner.Text(), err))
		} else {
			if _, err = cmd.Exec(person.Email, person.Hash, person.Name); err != nil {
				errorList = append(errorList, fmt.Sprintf("%d %s: %s", lineNum, person.Name, err))
			}
		}
		lineNum++
	}

	return errorList, scanner.Err()
}