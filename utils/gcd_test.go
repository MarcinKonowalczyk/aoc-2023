package utils

import "testing"

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
	}

	for _, test := range tests {
		result := GCD(test.a, test.b)
		AssertEqual(t, result, test.expected)
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
