package internal

import (
	"fmt"
	"os"
)

func LoadData(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Read file error: %s", err)
	}
	return data, nil
}