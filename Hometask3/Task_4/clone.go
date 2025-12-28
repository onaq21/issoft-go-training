package main

import "fmt"

type Cloner interface {
	Clone() Cloner
}

type Point struct {
	X int
	Y int
}
type Person struct {
	Name string
	Age int
}

func (p *Point) Clone() Cloner {
	return &Point{p.X, p.Y}
}
func (p *Point) String() string {
	return fmt.Sprintf("X: %d, Y: %d ", p.X, p.Y)
}

func (p *Person) Clone() Cloner {
	return &Person{p.Name, p.Age}
}
func (p *Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d ", p.Name, p.Age)
}

func sliceClone(list []interface{}) (listClone []interface{}) {
	listClone = make([]interface{}, 0, len(list))
	for _, el := range list {
		switch v := el.(type) {
			case int, float64, bool, string: 
				listClone = append(listClone, v)
			case Cloner:
				listClone = append(listClone, v.Clone())
		}
	}
	return
}

func main() {
	list := []interface{}{
		45,
		&Point{5, 10},
		&Person{"Kostya", 18},
		&Point{2, 2},
		false,
		"Top",
		23.5,
	}

	cloneList := sliceClone(list)

	fmt.Println(cloneList...)

	if val, ok := cloneList[1].(*Point); ok {
		val.X = 999
	}
	fmt.Println("Original Point:", list[1])

	cloneList[4] = true
	fmt.Println("Original bool:", list[4])
}