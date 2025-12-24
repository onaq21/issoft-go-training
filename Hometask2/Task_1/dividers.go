package main

import "fmt"

func countSimpleDividers(number int) (string, bool) {
	if number < 2 {
		return "", false
	}

	simpleDividers := countFactorsOfTwo(number)

	number >>= simpleDividers

	for i := 3; i * i <= number; i += 2 {
		for number % i == 0 {
			number /= i
			simpleDividers++
		}
	}

	if number > 1 {	simpleDividers++ }

	if simpleDividers & 1 == 0 {
		return "even", true
	} else {
		return "odd", true
	}
}

func countFactorsOfTwo(number int) (factor int) {
	for number & 1 == 0 {
		number >>= 1
		factor++
	}
	return
}

func main() {
	fmt.Println(countSimpleDividers(15))					//even true
	fmt.Println(countSimpleDividers(7))						//odd true
	fmt.Println(countSimpleDividers(-2))					//"" false
	fmt.Println(countSimpleDividers(2147483647))	//odd true
}