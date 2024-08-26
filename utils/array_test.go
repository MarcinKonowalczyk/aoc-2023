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
	result, idx, _ := ArrayMin(arr)
	AssertEqual(t, result, 1)
	AssertEqual(t, idx, 0)
}

func TestMaxArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result, idx, _ := ArrayMax(arr)
	AssertEqual(t, result, 5)
	AssertEqual(t, idx, 4)
}

func TestArrayArrayIntersection(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}
	result := ArrayArrayIntersection(arr1, arr2)
	AssertEqualArrays(t, result, []int{3, 4, 5})
}

func TestArrayPartition(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	left, right := ArrayPartition(arr, func(a int) bool { return a < 3 })
	AssertEqualArrays(t, left, []int{1, 2})
	AssertEqualArrays(t, right, []int{3, 4, 5})
}

func TestArrayDiff(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	diff_0, err := ArrayDiff(arr1, 0)
	AssertNoError(t, err)
	AssertEqualArrays(t, diff_0, []int{1, 2, 3, 4, 5})
	diff_1, err := ArrayDiff(arr1, 1)
	AssertNoError(t, err)
	AssertEqualArrays(t, diff_1, []int{1, 1, 1, 1})
	diff_2, err := ArrayDiff(arr1, 2)
	AssertNoError(t, err)
	AssertEqualArrays(t, diff_2, []int{0, 0, 0})
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

func TestArrayReverse(t *testing.T) {
	// Test odd
	arr1 := []int{1, 2, 3, 4, 5}
	result := ArrayReverse(arr1)
	AssertEqualArrays(t, result, []int{5, 4, 3, 2, 1})
	// Test even
	arr2 := []int{1, 2, 3, 4}
	result = ArrayReverse(arr2)
	AssertEqualArrays(t, result, []int{4, 3, 2, 1})
}

func TestArrayUnique(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	result := ArrayUnique(arr1)
	AssertEqualArrays(t, result, []int{1, 2, 3, 4, 5})
	arr2 := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	result = ArrayUnique(arr2)
	AssertEqualArrays(t, result, []int{1, 2, 3, 4, 5})
}

func TestArrayUniqueEmpty(t *testing.T) {
	arr1 := []int{}
	result := ArrayUnique(arr1)
	AssertEqualArrays(t, result, []int{})
}

func TestArray_RemoveIndices_Basic(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	to_remove := []int{0, 2, 4}
	result, n_removed := ArrayRemoveIndices(arr, to_remove...)
	AssertEqualArrays(t, result, []int{2, 4})
	AssertEqual(t, n_removed, 3)
}

func TestArray_RemoveIndices_IndicesOutOfBounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	to_remove := []int{0, 2, 10}
	result, n_removed := ArrayRemoveIndices(arr, to_remove...)
	AssertEqualArrays(t, result, []int{2, 4, 5})
	AssertEqual(t, n_removed, 2)
}

func TestArray_RemoveIndices_UnsortedIndices(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	to_remove := []int{4, 2, 0}
	result, _ := ArrayRemoveIndices(arr, to_remove...)
	AssertEqualArrays(t, result, []int{2, 4})
}

func TestArray_RemoveIndices_DuplicateIndices(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	to_remove := []int{0, 0, 2, 4, 4}
	result, _ := ArrayRemoveIndices(arr, to_remove...)
	AssertEqualArrays(t, result, []int{2, 4})
}

func TestArray_RemoveIndices_EmptyArray(t *testing.T) {
	arr := []int{}
	to_remove := []int{0, 2, 4}
	result, _ := ArrayRemoveIndices(arr, to_remove...)
	AssertEqualArrays(t, result, []int{})
}

func TestArray_RemoveIndices_EmptyToRemove(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result, _ := ArrayRemoveIndices(arr)
	AssertEqualArrays(t, result, []int{1, 2, 3, 4, 5})
}
