package main

import (
	"Task_1/internal"
	"flag"
	"path/filepath"
	"fmt"
	"os"
)

func main() {
	var fileName string
	var wordsNum int
	flag.StringVar(&fileName, "file", "", "input text file")
	flag.StringVar(&fileName, "f", "", "shortcut for -file")
	flag.IntVar(&wordsNum, "top", 10, "number of words to display")
	flag.IntVar(&wordsNum, "t", 10, "shortcut for -top")

	flag.Parse()

	if fileName == "" {
		fmt.Println("Flag -file or -f is required")
		os.Exit(1)
	}
	if wordsNum < 0 {
		fmt.Println("Flag -top (-t) must be a positive integer")
		os.Exit(1)
	}

	filePath := filepath.Join("data", fileName)

	data, err := internal.LoadData(filePath)
	if err != nil {
		fmt.Printf("Load data error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Load data: success")

	words := internal.MapToSlice(internal.ProcessData(data))
	internal.SortSlice(words)

	if wordsNum > len(words) { wordsNum = len(words) }

	fmt.Printf("Top %d most popular words in your file:\n", wordsNum)
	for i := range wordsNum {
		fmt.Printf("%d. %s - %d\n", i + 1, words[i].Word, words[i].Count)
	}
}