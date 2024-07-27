package utils

import (
	"testing"
)

func TestMapMap(t *testing.T) {
	m := map[int]int{
		1: 2,
		3: 4,
		5: 6,
	}
	e := map[int]int{
		1: 4,
		3: 8,
		5: 12,
	}
	result := MapMap(m, func(n int) int { return n * 2 })
	AssertEqualMaps(t, result, e)
}
