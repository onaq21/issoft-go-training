package users

import (
	"strings"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ParseUser(str string) (user User, err error) {
	data := strings.Fields(str)

	if len(data) < 2 {
		return user, fmt.Errorf("Too short")
	} else if len(data) > 2 {
		return user, fmt.Errorf("Too long")
	}

	for _, word := range data {
		for _, ch := range word {
			if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
				return user, fmt.Errorf("Invalid data")
			}
		}
	}

	user.Name = strings.Join(data, " ")
	user.Email = strings.Join(data, "") + "@coolcompany.com"

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Email), bcrypt.DefaultCost)
	if err != nil {
		return user, fmt.Errorf("Generate hash error: %s", err)
	}

	user.Hash = string(hashBytes)
	user.IsActive = true

	return user, nil
}