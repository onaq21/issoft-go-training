package employee

import (
	"errors"
	"strings"
	"time"
	"fmt"
)

func ParseEmployee(str string) (e Employee, err error) {
	data := strings.Split(str, ",")

	if len(data) < 3 {
		return e, errors.New("Too short")
	} else if len(data) > 3 {
		return e, errors.New("Too long")
	}

	parts := strings.Split(data[0], " ")

	if len(parts) != 2 {
		return e, errors.New("Invalid name")
	}

	e.FirstName = parts[0]
	e.LastName = parts[1]

	e.StartDate, err = time.Parse("1/2/2006", data[1])
	if err != nil {
		return e, fmt.Errorf("Invalid start date: %s", err)
	}

	e.EndDate, err = time.Parse("1/2/2006", data[2])
	if err != nil {
		return e, fmt.Errorf("Invalid end date: %s", err)
	}

	return e, nil
}