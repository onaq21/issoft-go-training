package main 

import (
	"Task_3/internal/file"
	"Task_3/internal/fruits"
	"fmt"
	"os"
)

func main() {
	items, err := file.LoadData() 
	if err != nil {
		fmt.Printf("Load data error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Load data: success")

	var (
		x, y float64
		k int
		choice string
	)

	for {
		fmt.Print("Enter fruit coordinates x, y and number of nearest neighbors k (the larger k, the more accurate prediction): ")

		if _, err := fmt.Scan(&x, &y, &k); err != nil {
			fmt.Println("Invalid input, try again")
			continue
		}

		fmt.Print(fruits.FindNearestFruit(items, x, y, k))

		fmt.Print("Do you want to continue? y/n: ")
		fmt.Scan(&choice)
		if len(choice) > 0 && (choice[0] == 'n' || choice[0] == 'N') { break }
	}
}