package file

import (
	"Task_2/internal/goods"
	"bufio"
	"fmt"
	"os"
)

func LoadData() (items []goods.Good, err error) {
	file, err := os.Open("data/sales.csv")
	if err != nil {
		return nil, fmt.Errorf("Open file error: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		item, err := goods.ParseGoods(scanner.Text())
		if err != nil {
			fmt.Printf("Parse error: %s", err)
		} else {
			items = append(items, item)
		}
	}
	
	return items, scanner.Err()
}