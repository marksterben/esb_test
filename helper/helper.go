package helper

import (
	"fmt"
	"strconv"
)

func IncrementNumberString(input string) (*string, error) {
	number, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}

	number++
	output := fmt.Sprintf("%04d", number)

	return &output, nil
}
