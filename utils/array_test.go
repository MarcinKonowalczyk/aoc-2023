package utils

import (
	"testing"
)

func TestArrayReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := ArrayReduce(arr, 0, func(a, b int) int { return a + b })
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

func TestArrayArrayIntersection(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}
	result := ArrayArrayIntersection(arr1, arr2)
	AssertEqualWithComparator(t, result, []int{3, 4, 5}, CompareArrays)
}

func TestYieldPairwise(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	output := make(chan [2]int)
	go YieldPairwise(arr, output)
	pairs := make([][2]int, 0)
	for {
		pair, ok := <-output
		if !ok {
			break
		}
		pairs = append(pairs, pair)
	}
	AssertEqualWithComparator(t, pairs, [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, CompareArrays)
}
