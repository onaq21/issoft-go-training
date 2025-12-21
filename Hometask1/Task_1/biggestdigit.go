package main

import "fmt"

const limit int = 1_000_000_000

func main() {
	var number int
	fmt.Printf("Enter your natural number (less than %d): ", limit)
	fmt.Scanln(&number)

	if number < 1 || number >= limit {
		fmt.Println("Input error")
		return
	}

	maxDigit := 0

	for number > 0 {
		if number % 10 > maxDigit {
			maxDigit = number % 10
		}
		number /= 10
	}

	fmt.Printf("Max digit is %d\n", maxDigit)
}