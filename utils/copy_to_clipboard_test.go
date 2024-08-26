package utils

import (
	"runtime"
	"testing"
)

func TestCopyToClipboard(t *testing.T) {

	err := CopyToClipboard("test")
	// Check if we're on macOS
	if runtime.GOOS == "darwin" {
		AssertNoError(t, err)
	} else {
		AssertError(t, err)
	}
}
