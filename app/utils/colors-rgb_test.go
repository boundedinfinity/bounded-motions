package utils_test

import (
	"go-motions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRgb(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected utils.RgbColor
		err      error
	}{
		{
			name:     "contains a",
			input:    "rgb(0, 0, 0)",
			expected: utils.RgbColor{R: 0, B: 0, G: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := utils.ParseRgb(tc.input)

			assert.ErrorIs(tt, err, tc.err)
			assert.Equal(tt, tc.expected, actual)
		})
	}

}
