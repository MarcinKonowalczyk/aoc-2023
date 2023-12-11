package utils

import (
	"cmp"
	"errors"
)

// Reduce an array using a reduce function. The accumulator is initialized to
// the initial value and the reduce function is applied to each element of the
// array and the accumulator. The result is returned as the final accumulator.
func ArrayReduce[T any, V any](arr []T, initial V, reduce func(V, T) V) V {
	result := initial
	for i := 0; i < len(arr); i++ {
		result = reduce(result, arr[i])
	}
	return result
}

// Reduce an array using a reduce function. The accumulator is initialized to
// the initial value and the reduce function is applied to each element of the
// array and the accumulator. The result is returned as the final accumulator.
// This version of the function allows the reduce function to return an error.
// If the reduce function returns an error, the function returns the error, as
// well as the current accumulator, immediately.
func ArrayReduceWithError[T any, V any](arr []T, initial V, reduce func(V, T) (V, error)) (V, error) {
	result := initial
	for i := 0; i < len(arr); i++ {
		var err error
		result, err = reduce(result, arr[i])
		if err != nil {
			return result, err
		}
	}
	return result, nil
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
func ArrayFilter[T any](arr []T, test func(T) bool) (ret []T) {
	for _, a := range arr {
		if test(a) {
			ret = append(ret, a)
		}
	}
	return
}

// Filter an array using a test. Elements passing the test are kept while those
// failing it are rejected. This version of the function allows the test to
// return an error. If the test returns an error, the function returns the error
// immediately.
func ArrayFilterWithError[T any](arr []T, test func(T) (bool, error)) ([]T, error) {
	ret := make([]T, 0)
	for _, a := range arr {
		passes, err := test(a)
		if err != nil {
			return nil, err
		}
		if passes {
			ret = append(ret, a)
		}
	}
	return ret, nil
}

// Map an array using a map function. The map function is applied to each
// element of the array and the result is returned as a new array.
func ArrayMap[T any, V any](arr []T, map_func func(T) V) []V {
	result := make([]V, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = map_func(arr[i])
	}
	return result
}

// Map an array using a map function. The map function is applied to each
// element of the array and the result is returned as a new array. This version
// of the function allows the map function to return an error. If the map
// function returns an error, the function returns the error immediately.
func ArrayMapWithError[T any, V any](arr []T, map_func func(T) (V, error)) ([]V, error) {
	result := make([]V, len(arr))
	for i := 0; i < len(arr); i++ {
		mapped, err := map_func(arr[i])
		if err != nil {
			return nil, err
		}
		result[i] = mapped
	}
	return result, nil
}

// N^2 intersection of two arrays - compare each element of the first array to
// each element of the second array and return the intersection.
func ArrayArrayIntersection[T comparable](a, b []T) []T {
	intersection := make([]T, 0)
	for _, n := range a {
		for _, m := range b {
			if n == m {
				intersection = append(intersection, n)
			}
		}
	}
	return intersection
}

// Interface for numeric types
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

func ArraySum[T Numeric](arr []T) T {
	var sum T = 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func ArrayPairwise[T any](arr []T) [][2]T {
	pairwise := make([][2]T, 0)
	for i := 0; i < len(arr)-1; i++ {
		pairwise = append(pairwise, [2]T{arr[i], arr[i+1]})
	}
	return pairwise
}
