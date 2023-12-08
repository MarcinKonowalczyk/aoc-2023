package utils

import (
	"testing"
)

func TestStringOfNumbersToNumbers(t *testing.T) {
	// s := "1 2 3"
	candidates := []string{
		"1 2 3",
		" 1 2 3",
		"1 2 3 ",
		" 1 2 3 ",
		" 1  2  3 ",
	}

	expected := []int{1, 2, 3}
	for _, s := range candidates {
		actual, err := StringOfNumbersToNumbers(s)
		AssertNoError(t, err)
		AssertEqualWithComparator(t, actual, expected, CompareArrays)
	}
}
