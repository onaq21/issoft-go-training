package errors

import (
	"os"
	"fmt"
)


func SaveErrors(errorsList []string) error {
	file, err := os.Create("errors.txt")
	if err != nil {
		return fmt.Errorf("Create file error: %s", err)
	}
	defer file.Close()

	for _, person := range errorsList {
		fmt.Fprintln(file, person)
	}
	return nil
}