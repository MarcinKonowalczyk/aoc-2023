package utils

import (
	"runtime"
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestCopyToClipboard(t *testing.T) {

	err := CopyToClipboard("test")
	// Check if we're on macOS
	if runtime.GOOS == "darwin" {
		assert.NoError(t, err)
	} else {
		assert.NotEqual(t, err, nil)
	}
}
