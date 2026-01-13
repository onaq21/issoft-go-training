package file 

import (
	"os"
	"fmt"
	"bufio"
	"Task_3/internal/fruits"
)

func LoadData() (items []fruits.Fruit, err error) {
	file, err := os.Open("data/data.txt")
	if err != nil {
		return nil, fmt.Errorf("Open file error: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fruit, err := fruits.ParseFruit(scanner.Text())
		if err != nil {
			fmt.Printf("Parse error: %s", err)
		} else {
			items = append(items, fruit)
		}
	}

	return items, scanner.Err()
}