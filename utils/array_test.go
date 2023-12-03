package utils

import (
	"testing"
)

func TestReduceArray(t *testing.T) {
	// Cast t to MyT
	arr := []int{1, 2, 3, 4, 5}
	result := ReduceArray(arr, 0, func(a, b int) int { return a + b })
	AssertEqual(t, result, 15)
	AssertEqual(t, result, 15)
}

func TestMinArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result, idx, _ := MinArray(arr)
	AssertEqual(t, result, 1)
	AssertEqual(t, idx, 0)
}

func TestMaxArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result, idx, _ := MaxArray(arr)
	AssertEqual(t, result, 5)
	AssertEqual(t, idx, 4)
}
