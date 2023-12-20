package utils

import (
	"testing"
)

func TestArrayReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := ArrayReduce(arr, 0, func(a, b int) int { return a + b })
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

func TestArrayArrayIntersection(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}
	result := ArrayArrayIntersection(arr1, arr2)
	AssertEqualWithComparator(t, result, []int{3, 4, 5}, CompareArrays)
}

func TestArrayPartition(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	left, right := ArrayPartition(arr, func(a int) bool { return a < 3 })
	AssertEqualWithComparator(t, left, []int{1, 2}, CompareArrays)
	AssertEqualWithComparator(t, right, []int{3, 4, 5}, CompareArrays)
}
