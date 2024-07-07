package utils

import (
	"testing"
)

func TestCsprintf(t *testing.T) {
	out := Csprintf(Black, "Black")
	AssertEqual(t, out, "\033[30mBlack\033[0m")
	out = Csprintf(Red, "Red")
	AssertEqual(t, out, "\033[31mRed\033[0m")
	out = Csprintf(Green, "Green")
	AssertEqual(t, out, "\033[32mGreen\033[0m")
}
