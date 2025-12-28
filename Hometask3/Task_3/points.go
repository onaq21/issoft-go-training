package main

import "fmt" 

type Pointer interface {
	Set(x, y float64)
	Get() (float64, float64)
}

type Point struct {
	X float64
	Y float64
}

type PointLabeled struct {
	Point
	Label string
}

func (point *Point) Set(x, y float64) {
	point.X = x
	point.Y = y
}
func (point *Point) Get() (float64, float64) {
	return point.X, point.Y
}

func normalize(points []Pointer) {
	maxX, minX, maxY, minY := GetMinMax(points)

	for _, point := range points {
		x, y := point.Get()
		var normX, normY float64

		if maxX != minX {
			normX = (x - minX) / (maxX - minX)
		} else {
			normX = 0.0
		}
		if maxY != minY {
				normY = (y - minY) / (maxY - minY)
		} else {
			normY = 0.0
		}

		point.Set(normX, normY)
	}
}

func GetMinMax(points []Pointer) (maxX, minX, maxY, minY float64) {
	maxX, maxY = points[0].Get()
	minX, minY = maxX, maxY

	for _, p := range points[1:] {
    x, y := p.Get()
    if x > maxX { maxX = x }
    if x < minX { minX = x }
    if y > maxY { maxY = y }
    if y < minY { minY = y }
}
	return
}

func main() {
	points := []Pointer {
		&Point{-20.0, 10.0},
		&PointLabeled{Point{0.0, 20.0}, "A"},
		&PointLabeled{Point{100.2, 5.4}, "B"},
		&Point{2.0, 3.4},
	}
	normalize(points)
	for _, point := range points {
		x, y := point.Get()
		fmt.Printf("X:%f, Y:%f\n", x, y)
	}
}
