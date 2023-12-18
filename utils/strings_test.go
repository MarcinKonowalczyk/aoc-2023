package utils

import (
	"testing"
)

func TestStringOfNumbersToInts(t *testing.T) {
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
		actual, err := StringOfNumbersToInts(s)
		AssertNoError(t, err)
		AssertEqualWithComparator(t, actual, expected, CompareArrays)
	}
}

func TestStringOfNumbersToInts_2(t *testing.T) {
	s := "   8234 -4234928394 1982312389081  0 "
	expected := []int{8234, -4234928394, 1982312389081, 0}
	actual, err := StringOfNumbersToInts(s)
	AssertNoError(t, err)
	AssertEqualWithComparator(t, actual, expected, CompareArrays)
}
