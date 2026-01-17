package internal

import "sort"

type WordCount struct {
	Word string
	Count int
}

func MapToSlice(words map[string]int) []WordCount{
	list := make([]WordCount, 0, len(words))

	for word, count := range words {
		list = append(list, WordCount{word, count})
	}

	return list
}

func SortSlice (words []WordCount) {
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})
}