package main

import "fmt"

func group(m map[byte]string) map[byte][]string {
	result := make(map[byte][]string, 10)

	for key, val := range m {
		digit := key % 10
		result[digit] = append(result[digit], val)
	}

	return result
}

func main() {
	data := map[byte]string{
		11: "red",
		51: "green",
		22: "blue",
		5:  "yellow",
		30: "black",
		44: "white",
	}

	fmt.Println(group(data))	/*
															0: [black]
															1: [red green]
															2: [blue]
															4: [white]
															5: [yellow] 
														*/
}