package main

type User struct {
	Id int					`json:"id"`
	Email string		`json:"email"`
	Hash string			`json:"-"`
	Name string			`json:"name"`
	IsActive bool		`json:"is_active"`
}