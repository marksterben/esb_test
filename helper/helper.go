package helper

import (
	"fmt"
	"strconv"
)

func IncrementNumberString(input string) string {
	number, _ := strconv.Atoi(input)
	number++
	output := fmt.Sprintf("%04d", number)

	return output
}
