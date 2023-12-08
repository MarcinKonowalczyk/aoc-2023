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
