package fruits

import (
	"sort"
	"fmt"
)

type neighbor struct {
	Name string 
	Distance float64
}

func FindNearestFruit(fruits []Fruit, x, y float64, k int) string {
	if k <= 0 || k > len(fruits) {
		return fmt.Sprintf("Invalid k, use k from %d to %d\n", 1, len(fruits))
	}

	neighbors := make([]neighbor, 0, len(fruits))
	for _, fruit := range fruits {
		neighbors = append(neighbors, neighbor{fruit.Name, (x - fruit.X) * (x - fruit.X) + (y - fruit.Y) * (y - fruit.Y)})
	}

	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})

	fruitsCount := map[string]int{}

	for i := range k {
		fruitsCount[neighbors[i].Name]++
	}

	leaders := []string{}
	maxCount := 0

	for name, count := range fruitsCount {
		if count > maxCount {
			leaders = []string{name}
			maxCount = count
		} else if count == maxCount {
			leaders = append(leaders, name)
		}
	}

	if len(leaders) == 1 {
		return fmt.Sprintf("Your fruit with x: %.2f, y: %.2f is probably %s\n", x, y, leaders[0])
	} else {
		return "It's 50/50 bro\n"
	}
}