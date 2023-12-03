package utils

import (
	"cmp"
	"errors"
)

func ReduceArray[T any, V any](arr []T, initial V, reduce func(V, T) V) V {
	result := initial
	for i := 0; i < len(arr); i++ {
		result = reduce(result, arr[i])
	}
	return result
}

func MinArray[T cmp.Ordered](arr []T) (T, int, error) {
	if len(arr) == 0 {
		var zero T
		return zero, -1, errors.New("array is empty")
	}
	var min T = arr[0]
	var min_idx int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			min_idx = i
		}
	}
	return min, min_idx, nil
}

func MaxArray[T cmp.Ordered](arr []T) (T, int, error) {
	if len(arr) == 0 {
		var zero T
		return zero, -1, errors.New("array is empty")
	}
	var max T = arr[0]
	var max_idx int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			max_idx = i
		}
	}
	return max, max_idx, nil
}

// https://stackoverflow.com/a/37563128/2531987
// Filter an array using a test. Elements passing the test are kept while those
// failing it are rejected.
func FilterArray[T any](arr []T, test func(T) bool) (ret []T) {
	for _, a := range arr {
		if test(a) {
			ret = append(ret, a)
		}
	}
	return
}