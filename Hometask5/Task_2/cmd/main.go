package main

import (
	"text/template"
	"time"
	"Task_2/internal/file"
	"Task_2/internal/report"
	"os"
	"fmt"
)

func getData() string {
	t := time.Now()
	return fmt.Sprintf("%d %s, %d", t.Day(), t.Month(), t.Year())
}


func totalValue(items map[string]float64) (result float64) {
	for _, price := range items {
		result += price
	}
	return
}

func topItem(items map[string]float64) string {
	var highest float64
	var topName string
	for name, price := range items {
		if price > highest {
			highest = price
			topName = name
		}
	}
	return fmt.Sprintf("%s: %.2f", topName, highest)
}

func main() {
	goods, err := file.LoadData()
	if err != nil {
		fmt.Printf("Load data error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Load data: success")

	items := file.ProcessData(goods)
	fmt.Println("Process data: success")

	t := template.New("main")

	t.Funcs(map[string]any {
		"getData": func() string { return getData() },
		"totalValue": func() float64 { return totalValue(items) },
		"topItem": func() string { return topItem(items) },
		"add": func(a, b int) int { return a + b },
	})

	listItems := report.MapToSlice(items)

	data := map[string]any {
		"items": listItems,
	}

	t, err = t.ParseFiles("templates/report.md")
	if err != nil {
		fmt.Printf("Load templates error: %s\n", err)
	}

	err = file.SaveReport(t, data)
	if err != nil {
		fmt.Printf("Save report failed: %s\n", err)
    os.Exit(1)
	}
	fmt.Println("Save file: success")
}