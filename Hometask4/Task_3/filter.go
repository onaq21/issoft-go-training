package main

import (
	"fmt" 
	"time"
)

func whereInParallel[T any](list []T, filter func(T) bool) []T {
	size := len(list)
	goroutines := min(size, 8)
	chunk := size / goroutines
	ch := make(chan []T, goroutines)
	remainder := size % goroutines

	for start, end := 0, 0; start < size; start = end {
		end += chunk
		if remainder > 0 {
			end++
			remainder--
		}
		go func(start, end int) {
			res := make([]T, 0, end - start)
			for _, val := range list[start:end] {
				if filter(val) {
					res = append(res, val)
				}
			}
			ch <- res
		}(start, end)
	}

	result := make([]T, 0, size)
	for i := 0; i < goroutines; i++ {
		result = append(result, <-ch...)
	}

	return result
}

func where(list []int, f func(int) bool) []int {
	result := make([]int, 0, len(list))
	for _, val := range list {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

func timeDecorator[T any](f func([]T, func(T) bool)[]T) func([]T, func(T) bool)[]T {
	return func(list []T, filter func(T) bool) []T {
		start := time.Now()
		result := f(list, filter)
		fmt.Printf("Execution time: %s\n", time.Since(start))
		return result
	}
}

func main() {
  size := 10_000_000
  list := make([]int, size)
  for i := 0; i < size; i++ {
    list[i] = i
  }

  timedWhere := timeDecorator(where)
  result1 := timedWhere(list, func(x int) bool {
		return x & 1 == 0
	})
	//Execution time: ~32ms
  fmt.Println("where result length:", len(result1))

  timedWhereParallel := timeDecorator(whereInParallel[int])
  result2 := timedWhereParallel(list, func(x int) bool {
		return x & 1 == 0
	})
	//Execution time: ~16ms
  fmt.Println("whereInParallel result length:", len(result2))
}
