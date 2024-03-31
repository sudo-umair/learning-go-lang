package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat64s(strings []string) ([]float64, error) {
	var float64s []float64

	for _, str := range strings {
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error parsing float: ", err)
			return nil, errors.New("error parsing float")
		}
		float64s = append(float64s, floatValue)
	}

	return float64s, nil
}
