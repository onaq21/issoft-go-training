package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter first positive number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	if a < 0 {
		a = 0
	}
	if b < 0 {
		b = 0
	}
	if a > b {
		a, b = b, a
	}

	for ; a <= b; a++ {
		count := 0
		temp := a
		for temp > 0 {
			if temp&1 == 1 {
				count++
			}
			temp >>= 1
		}
		if count == 4 {
			fmt.Println(a)
		}
	}
}
