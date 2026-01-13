package file

import (
	"Task_1/internal/employee"
	"fmt"
)

func ProcessData(employees []employee.Employee) []string {
	sortedEmployees := make([]string, 0, len(employees))

	employee.SortByVacationDuration(employees)

	for _, emp := range employees {
		days := int(emp.EndDate.Sub(emp.StartDate).Hours() / 24)
		sortedEmployees = append(sortedEmployees, fmt.Sprintf("%s %s: %d", emp.FirstName, emp.LastName, days))
	}
	return sortedEmployees
}