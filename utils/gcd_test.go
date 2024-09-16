package utils

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{48, 18, 6},
		{56, 98, 14},
		{101, 103, 1},
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		{7, 3, 1},
		{4, 5, 1},
		{21, 6, 3},
		{8, 9, 1},
		{123*5 + 13, 123, 1},
	}

	for _, test := range tests {
		result := GCD(test.a, test.b)
		AssertEqual(t, result, test.expected)
		// Also test the ExtendedGCD function
		result, x, y := ExtendedGCD(test.a, test.b)
		if testing.Verbose() {
			fmt.Printf("GCD(%d, %d) = %d*(%d) + %d*(%d) = %d\n", test.a, test.b, test.a, x, test.b, y, test.a*x+test.b*y)
		}
		AssertEqual(t, result, test.expected)
		AssertEqual(t, test.a*x+test.b*y, result)
	}
}
func TestLCM(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{4, 5, 20},
		{21, 6, 42},
		{8, 9, 72},
		{0, 5, 0},
		{5, 0, 0},
		{7, 3, 21},
	}

	for _, test := range tests {
		result := LCM(test.a, test.b)
		AssertEqual(t, result, test.expected)
	}
}
