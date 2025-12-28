package main

import "fmt"

type Color [3]byte

func (color *Color) Print() {
	fmt.Printf("Red: %d, Green: %d, Blue: %d\n", color[0], color[1], color[2])
}

func (color *Color) GetR() byte{
	return color[0]
}
func (color *Color) GetG() byte{
	return color[1]
}
func (color *Color) GetB() byte{
	return color[2]
}

func (color *Color) SetR(val byte) {
	color[0] = val
}
func (color *Color) SetG(val byte) {
	color[1] = val
}
func (color *Color) SetB(val byte) {
	color[2] = val
}

func (color *Color) GetBrightness() float64 {
	return 0.2126 * float64(color[0]) + 0.7152 * float64(color[1]) + 0.0722 * float64(color[2])
}

func maxBrightness(colors []Color) *Color {
	if len(colors) == 0 {
		return nil
	}
	var (
		max float64
		maxIndex int
	)

	for i, color := range colors {
		if val := color.GetBrightness(); val > max {
			max = val
			maxIndex = i
		}
	}

	return &colors[maxIndex]
}

func main() {
	colors := []Color{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{255, 255, 255},
		{128, 128, 128},
	}

	fmt.Println("All colors:")
	for _, c := range colors {
		c.Print()
		fmt.Printf("Brightness: %.1f\n\n", c.GetBrightness())
	}

	brightest := maxBrightness(colors)
	fmt.Println("The color with the highest brightness:")
	brightest.Print()
	fmt.Printf("Brightness: %.2f\n", brightest.GetBrightness())
}
