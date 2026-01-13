package goods

import (
	"errors"
	"strings"
	"strconv"
	"fmt"
)

func ParseGoods(str string) (good Good, err error) {
	data := strings.Split(str, ",")

	if len(data) != 3 {
		return good, errors.New("Invalid data")
	}

	good.Name = strings.Trim(data[0], `"`)

	good.Price, err = strconv.ParseFloat(data[1], 64)
	if err != nil {
		return good, fmt.Errorf("Can't parse price: %s", data[1])
	}

	good.Units, err = strconv.Atoi(data[2])
	if err != nil {
		return good, fmt.Errorf("can't parse units: %s", data[2])
	}
	
	return good, nil
}