package internal

import (
	"strings"
	"regexp"
)

func ProcessData(data []byte) map[string]int {
	counter := make(map[string]int)
	text := strings.ToLower(string(data))

	re := regexp.MustCompile(`[^a-zA-Zа-яА-Я\d\s]`)
	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)

	for _, word := range words {
		counter[word]++
	}
	return counter
}