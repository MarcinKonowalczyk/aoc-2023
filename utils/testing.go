package utils

import (
	"runtime"
	"testing"
)

func Assert(t *testing.T, predicate bool, msg string) {
	if !predicate {
		parent, _, _, _ := runtime.Caller(1)
		info := runtime.FuncForPC(parent)
		file, line := info.FileLine(parent)
		t.Errorf(msg+" in %s:%d", file, line)
	}
}

func AssertEqual[T comparable](t *testing.T, a T, b T) {
	if a != b {
		parent, _, _, _ := runtime.Caller(1)
		info := runtime.FuncForPC(parent)
		file, line := info.FileLine(parent)
		t.Errorf("Expected %v == %v in %s:%d", a, b, file, line)
	}
}
