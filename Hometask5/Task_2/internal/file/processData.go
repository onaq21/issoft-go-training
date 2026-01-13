package file

import "Task_2/internal/goods"

func ProcessData(items []goods.Good) map[string]float64 {
	result := make(map[string]float64, len(items))

	for _, item := range items {
		result[item.Name] += item.Price * float64(item.Units)
	}

	return result
}