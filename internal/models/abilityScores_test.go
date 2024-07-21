package models

import (
	"strconv"
	"testing"
)

func TestMod(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: -5,
		},
		{
			input:    2,
			expected: -4,
		},
		{
			input:    3,
			expected: -4,
		},
		{
			input:    10,
			expected: 0,
		},
		{
			input:    12,
			expected: 1,
		},
		{
			input:    30,
			expected: 10,
		},
	}

	for _, tc := range testCases {

		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {

			actual := AbilityScoreMod(tc.input)
			if actual != tc.expected {
				t.Errorf("got %d, want %d", actual, tc.expected)
			}
		})
	}

}
