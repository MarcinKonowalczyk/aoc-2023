package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

func FileToLines(file string) ([]string, error) {
	data_bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read input file '%s'. error: %s", file, err)
	}

	if len(data_bytes) == 0 {
		return nil, fmt.Errorf("input file '%s' is empty", file)
	}

	data := string(data_bytes[:])

	if !IsAscii(data) {
		return nil, fmt.Errorf("input file '%s' contains non-ASCII characters", file)
	}

	// Split into lines
	lines := strings.Split(data, "\n")

	// Remove empty lines
	// lines = filter(lines, func(s string) bool { return len(s) > 0 })

	// Remove comments
	// lines = filter(lines, func(s string) bool { return s[:2] != "//" })

	// for _, line := range lines {
	// 	println(line)
	// }

	return lines, nil
}

func ResolvePath(file string) (string, error) {
	file = path.Clean(file)
	fs, err := os.Stat(file)
	if err == nil {
		if fs.IsDir() {
			return file, fmt.Errorf("file '%s' is a directory", file)
		}
		// Otherwise file exists and is a file
	} else if errors.Is(err, os.ErrNotExist) {
		return file, fmt.Errorf("file '%s' does not exists", file)
	} else {
		// https://stackoverflow.com/a/12518877/2531987
		return file, fmt.Errorf("file '%s' is a Schrodinger file", file)
	}
	return file, nil
}

// https://stackoverflow.com/a/53069799/2531987
func IsAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func Stopf(format string, a ...any) {
	if format[len(format)-1] != byte('\n') {
		format = format + "\n"
	}
	_, err := fmt.Printf(format, a...)
	if err != nil {
		panic(err)
	}
	os.Exit(1)
}
