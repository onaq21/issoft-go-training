package main

import (
	"fmt"
	"regexp"
)

type Person struct {
	Name  string
	Phone string
	Email string
}

func main() {
	p := &Person{}

	for {
		var name string
		fmt.Print("Enter your name (max 20 chars, first capital): ")
		fmt.Scanln(&name)
		pattern := regexp.MustCompile(`^[A-Z][a-z]{0,19}$`)
		if pattern.MatchString(name) {
			p.Name = name
			break
		}
		fmt.Println("Invalid name, try again")
	}

	for {
		var phone string
		fmt.Print("Enter your phone (start with + and 12 digits): ")
		fmt.Scanln(&phone)
		pattern := regexp.MustCompile(`^\+\d{12}$`)
		if pattern.MatchString(phone) {
			p.Phone = phone
			break
		}
		fmt.Println("Invalid phone, try again")
	}

	for {
		var email string
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		pattern := regexp.MustCompile(`^[\w.+-]+@[a-z]{2,8}\.[a-z-]{2,8}$`)
		if pattern.MatchString(email) {
			p.Email = email
			break
		}
		fmt.Println("Invalid email, try again")
	}

	fmt.Printf("Name: %s\nPhone: %s\nEmail: %s\n", p.Name, p.Phone, p.Email)
}
