package utils

import (
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestCsprintf(t *testing.T) {
	out := Csprintf(Black, "Black")
	assert.Equal(t, out, "\033[0;30mBlack\033[0m")
	out = Csprintf(Red, "Red")
	assert.Equal(t, out, "\033[0;31mRed\033[0m")
	out = Csprintf(Green, "Green")
	assert.Equal(t, out, "\033[0;32mGreen\033[0m")
}
