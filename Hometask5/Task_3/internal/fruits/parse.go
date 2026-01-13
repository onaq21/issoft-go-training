package fruits

import (
	"strings"
	"strconv"
	"fmt"
)

func ParseFruit(str string) (fruit Fruit, err error) {
	data := strings.Split(str, ",")
	if len(data) != 3 {
		return fruit, fmt.Errorf("Parse error: wrong data")
	}

	fruit.Name = data[0]

	fruit.X, err = strconv.ParseFloat(data[1], 64)
	if err != nil {
		return fruit, fmt.Errorf("Can't parse X: %s", err)
	}

	fruit.Y, err = strconv.ParseFloat(data[2], 64)
	if err != nil {
		return fruit, fmt.Errorf("Can't parse Y: %s", err)
	}

	return fruit, nil
}