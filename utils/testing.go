package utils

import (
	"runtime"
	"testing"
)

func getParentInfo() (string, int) {
	parent, _, _, _ := runtime.Caller(2)
	info := runtime.FuncForPC(parent)
	file, line := info.FileLine(parent)
	return file, line
}

func Assert(t *testing.T, predicate bool, msg string) {
	if !predicate {
		file, line := getParentInfo()
		t.Errorf(msg+" in %s:%d", file, line)
	}
}

func AssertEqual[T comparable](t *testing.T, a T, b T) {
	if a != b {
		file, line := getParentInfo()
		t.Errorf("Expected %v == %v in %s:%d", a, b, file, line)
	}
}

func AssertNotEqual[T comparable](t *testing.T, a T, b T) {
	if a == b {
		file, line := getParentInfo()
		t.Errorf("Expected %v != %v in %s:%d", a, b, file, line)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		file, line := getParentInfo()
		t.Errorf("Expected no error, got '%v' in %s:%d", err, file, line)
	}
}

func AssertError(t *testing.T, err error) {
	if err == nil {
		file, line := getParentInfo()
		t.Errorf("Expected error, got '%v' in %s:%d", err, file, line)
	}
}

func AssertEqualWithComparator[T any](t *testing.T, a T, b T, comparator func(T, T) bool) {
	if !comparator(a, b) {
		file, line := getParentInfo()
		t.Errorf("Expected %v == %v in %s:%d", a, b, file, line)
	}
}

func CompareArrays[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareMaps[T comparable, V comparable](a map[T]V, b map[T]V) bool {
	if len(a) != len(b) {
		return false
	}
	var vb V
	var ok bool

	// NOTE: the range on a map is in random order
	for k, va := range a {
		// Check if key exists in b
		if vb, ok = b[k]; !ok {
			return false
		}
		// Check if value is the same
		if va != vb {
			return false
		}
	}

	// All keys of a exist in b, and a and b have the same length, hence they
	// must have the same keys
	return true
}
