package employee

import "time"

type Employee struct {
    FirstName string
    LastName  string
    StartDate time.Time
    EndDate   time.Time
}