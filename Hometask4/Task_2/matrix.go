package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Matrix[T Number] struct {
	rows int
	columns int
	data map[int]T
}

func newMatrix[T Number] (rows, columns int) *Matrix[T] {
	if rows <= 0 || columns <= 0 {
		panic("Wrong size")
	}
	return &Matrix[T]{rows, columns, make(map[int]T, rows * columns)}
}

func (matrix *Matrix[T]) Get(i, j int) T {
	if i < 0 || i >= matrix.rows || j < 0 || j >= matrix.columns {
		panic("Index out of bounds")
	}
	return matrix.data[i * matrix.columns + j]
}

func (matrix *Matrix[T]) Set(i, j int, data T) {
	if i < 0 || i >= matrix.rows || j < 0 || j >= matrix.columns {
		panic("Index out of bounds")
	}
	var zero T
	if data == zero {
		delete(matrix.data, i * matrix.columns + j)
	} else {
		matrix.data[i * matrix.columns + j] = data
	}
}

func main() {
  matrix := newMatrix[int](3, 3)

  matrix.Set(0, 0, 10)
  matrix.Set(1, 2, 20)

  for i := 0; i < matrix.rows; i++ {
    for j := 0; j < matrix.columns; j++ {
      fmt.Print(matrix.Get(i, j), " ")
    }
    fmt.Println()
  }

  matrix.Set(1, 2, 0)

  fmt.Println("After zeroing (1,2):")
  for i := 0; i < matrix.rows; i++ {
    for j := 0; j < matrix.columns; j++ {
      fmt.Print(matrix.Get(i, j), " ")
    }
    fmt.Println()
  }
}