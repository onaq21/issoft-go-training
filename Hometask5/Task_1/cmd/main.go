package main

import (
	"Task_1/internal/file"
	"fmt"
	"os"
)

func main() {
	employees, errorList, err := file.LoadData()
	if err != nil {
		fmt.Printf("Load data error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Load data: success")

	result := file.ProcessData(employees)
	fmt.Println("Process data: success")

	if err := file.SaveData(result, errorList); err != nil {
		fmt.Printf("Save data error: %s", err)
		os.Exit(1)
	}
	fmt.Println("Save data: success")
}