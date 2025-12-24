package main

import "fmt"

func where(list []int, f func(int) bool) (result []int) {
	result = make([]int, 0, len(list))
	for _, val := range list {
		if f(val) {
			result = append(result, val)
		}
	}
	return
}

func foreach(list [] int, f func(int)) {
	for _, val := range list {
		f(val)
	}
}

func main() {
	list := []int{1,2,3,4,5,6,7,8,9}

	fmt.Println(where(list, func(n int) bool { return n & 1 == 0 }))
	foreach(list, func(n int) { if n & 1 == 1 { fmt.Printf("%d ", n) } })
	fmt.Println()
}