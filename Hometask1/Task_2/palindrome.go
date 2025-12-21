package main

import "fmt"

const base = 3

func main() {
	var number, reverse int
	fmt.Print("Enter your positive number: ")
	fmt.Scanln(&number)

	if number <= 0 {
		fmt.Println("Input error")
	} else {
		for temp := number; temp > 0; temp /= base {
			reverse = reverse * base + temp % base
		}

		if number == reverse {
			fmt.Println("Your number is palindrom in ternary system")
		} else {
			fmt.Println("Your number isn't palindrom in ternary system")
		}
	}
}