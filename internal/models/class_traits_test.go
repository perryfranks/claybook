package models

import (
	"testing"
)

func TestClassTraitsSetAdd(t *testing.T) {
	testCases := []struct {
		name        string
		ct          ClassTraits
		set         ClassTraitsSet
		expectedLen int
	}{
		{
			"Empty add valid Class",
			ClassTraits{
				ClassName: Pugilist,
				Traits:    []Trait{},
			},
			ClassTraitsSet{},
			1,
		},
		{
			"Empty wrong class",
			ClassTraits{
				ClassName: 90,
				Traits:    []Trait{},
			},
			ClassTraitsSet{},
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.set.Add(tc.ct)
			if len(tc.set.ClassTraits) != tc.expectedLen {
				t.Errorf("got %d, expected %d", len(tc.set.ClassTraits), tc.expectedLen)
			}
		})
	}
}

func TestClassString(t *testing.T) {
	testCases := []struct {
		name   string
		input  Class
		expect string
	}{
		{
			"Valid - Pugilist",
			Pugilist,
			"Pugilist",
		},
		{
			"Valid - Sorcrerer",
			Sorcerer,
			"Sorcrerer",
		},
		{
			"Invalid - 99",
			99,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.String()
			if actual != tc.expect {
				t.Errorf("expected %v, got %v", tc.expect, actual)
			}
		})
	}
}
