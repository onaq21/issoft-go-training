package employee

import "sort"

func SortByVacationDuration(employees []Employee) {
	sort.SliceStable(employees, func(i, j int) bool {
		dur1 := employees[i].EndDate.Sub(employees[i].StartDate)
		dur2 := employees[j].EndDate.Sub(employees[j].StartDate)
		if dur1 != dur2 {
			return dur1 > dur2
		}
		return employees[i].LastName < employees[j].LastName
	})
}