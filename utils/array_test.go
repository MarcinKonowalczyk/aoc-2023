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

func TestArrayDiff(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	diff_0, err := ArrayDiff(arr1, 0)
	AssertNoError(t, err)
	AssertEqualWithComparator(t, diff_0, []int{1, 2, 3, 4, 5}, CompareArrays)
	diff_1, err := ArrayDiff(arr1, 1)
	AssertNoError(t, err)
	AssertEqualWithComparator(t, diff_1, []int{1, 1, 1, 1}, CompareArrays)
	diff_2, err := ArrayDiff(arr1, 2)
	AssertNoError(t, err)
	AssertEqualWithComparator(t, diff_2, []int{0, 0, 0}, CompareArrays)
}

func TestArrayAll(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	result := ArrayAll(arr1, func(a int) bool { return a < 6 })
	AssertEqual(t, result, true)
	result = ArrayAll(arr1, func(a int) bool { return a < 5 })
	AssertEqual(t, result, false)
}

func TestArrayAny(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	result := ArrayAny(arr1, func(a int) bool { return a < 6 })
	AssertEqual(t, result, true)
	result = ArrayAny(arr1, func(a int) bool { return a < 0 })
	AssertEqual(t, result, false)
}
