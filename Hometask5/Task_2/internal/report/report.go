package report

type ItemReport struct {
	Name string
	Price float64
}

func MapToSlice(items map[string]float64) []ItemReport{
	result := make([]ItemReport, 0, len(items))

	for name, price := range items {
		result = append(result, ItemReport{name, price})
	}
	return result
}