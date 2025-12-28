package main

import "fmt"

type Matrix struct {
	rows int
	columns int
	data []float64
}

func newMatrix(rows, columns int) *Matrix {
	if rows < 1 || columns < 1 {
		return &Matrix{1, 1, make([]float64, 1)}
	}
	return &Matrix{rows, columns, make([]float64, rows * columns)}
}

func (matrix *Matrix) Get(i, j int) float64 {
	if i < 0 || i >= matrix.rows || j < 0 || j >= matrix.columns { return 0 }
	return matrix.data[i * matrix.columns + j]
}
func (matrix *Matrix) Set(i, j int, value float64) {
	if i < 0 || i >= matrix.rows || j < 0 || j >= matrix.columns { return }
  matrix.data[i * matrix.columns + j] = value
}

func (matrix *Matrix) Print() {
	rows, columns := matrix.rows, matrix.columns
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%.1f ", matrix.data[i * columns + j])
		}
		fmt.Println()
	}
}

func main() {
	var matrix *Matrix = newMatrix(3, 3)
	matrix.Set(0, 0, 10)
	matrix.Set(0, 1, 20)
	matrix.Set(0, 2, 30)
	matrix.Set(1, 0, 40)
	matrix.Set(1, 1, 50)
	matrix.Set(1, 2, 60)
	matrix.Set(2, 0, 70)
	matrix.Set(2, 1, 80)
	matrix.Set(2, 2, 90)
	matrix.Print()

	fmt.Println()
	matrix = newMatrix(0, 10)
	matrix.Print()
}
