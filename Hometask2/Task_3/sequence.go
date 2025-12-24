package main

import "fmt"

func sequence(list ...int) []int {
	switch n := len(list); n {
		case 1:
			a := list[0]
			result := make([]int, 0, max(-a, a) + 1)
			if a > 0 {
				for i := 0; i <= a; i++ {
					result = append(result, i)
				}
			}	else {
				for i := a; i <= 0; i++ {
					result = append(result, i)
				}
			}
			return result
		case 2:
			a, b := list[0], list[1]
			if a > b {
				a, b = b, a
			}
			result := make([]int, 0, b - a + 1)

			for ; a <= b; a++ {
				result = append(result, a)
			}
			return result
	}
	return list
}

func main() {
	fmt.Println(sequence())							//[]
	fmt.Println(sequence(5))						//0 1 2 3 4 5
	fmt.Println(sequence(-3))						//-3 -2 -1 0
	fmt.Println(sequence(1,2,23,4,99))	//1 2 23 4 99
}