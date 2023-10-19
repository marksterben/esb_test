package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementNumberString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0000", "0001"},
		{"0009", "0010"},
		{"0099", "0100"},
		{"9999", "10000"}, // Test case for overflow
	}

	for _, tc := range testCases {
		t.Run("input:"+tc.input, func(t *testing.T) {
			result, err := IncrementNumberString(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
