package file

import (
	"Task_1/internal/employee"
	"bufio"
	"fmt"
	"os"
)

func LoadData() (employees []employee.Employee, errorList []string, err error) {
	file, err := os.Open("data/data.csv")
	if err != nil {
		return nil, nil, fmt.Errorf("Open file error: %s", err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		employee, err := employee.ParseEmployee(scanner.Text())
		if err != nil {
			errorList = append(errorList, fmt.Sprintf("%d %s", lineNum, scanner.Text()))
		} else {
			employees = append(employees, employee)
		}
		lineNum++
	}

	err = scanner.Err()
	
	return 
}