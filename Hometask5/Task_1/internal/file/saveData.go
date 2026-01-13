package file

import (
	"os"
	"fmt"
)

func SaveData(employees []string, errorList []string) error {
	employeesFile, err := os.Create("out.txt")
	if err != nil {
		return fmt.Errorf("Create file error: %s", err)
	}

	defer employeesFile.Close()

	for _, employee := range employees {
		fmt.Fprintln(employeesFile, employee)
	}

	if len(errorList) != 0 {
		employeesErrorFile, err := os.Create("errors.txt")
		if err != nil {
			return fmt.Errorf("Create file error: %s", err)
		}

		defer employeesErrorFile.Close()

		for _, employeeErr := range errorList {
			fmt.Fprintln(employeesErrorFile, employeeErr)
		}
	}

	return nil
}