package conv

import (
	"fmt"
	"strconv"
)

func Ints(values ...string) ([]int, error) {
	numbers := make([]int, 0, len(values))

	for index, val := range values {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("could not convert value `%q` at index `%d`", val, index)
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}
