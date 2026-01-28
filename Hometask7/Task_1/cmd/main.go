package main

import (
	"Task_1/internal/db"
	"Task_1/internal/errors"
	"log"
)

func main() {
	errorsList, err := db.SetupDatabase()
	if err != nil {
		log.Fatalf("Setup db error: %s", err)
	}
	log.Println("Setup db success")

	if len(errorsList) != 0 {
		if err = errors.SaveErrors(errorsList); err != nil {
			log.Fatalf("Save errors list error: %s", err)
		}
		log.Printf("Save errors success")
	}
}