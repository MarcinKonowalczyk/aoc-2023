package main

import (
	"aoc2023/day01"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

func stopf(format string, a ...any) {
	if format[len(format)-1] != byte('\n') {
		format = format + "\n"
	}
	_, err := fmt.Printf(format, a...)
	if err != nil {
		panic(err)
	}
	os.Exit(1)
}
func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 3 {
		stopf("Usage: aoc2023 <day> <part> <input-filename>")
	}

	day_string := flag.Arg(0)
	part_string := flag.Arg(1)
	file := flag.Arg(2)

	day, err := strconv.Atoi(day_string)
	if err != nil {
		stopf("Cannot convert string \"%s\" to integer (day)", day_string)
	}

	part, err := strconv.Atoi(part_string)
	if err != nil {
		stopf("Cannot convert string \"%s\" to integer (part)", part_string)
	}

	if day < 1 || day > 30 {
		stopf("Expected day in rage [0, 30]. Got: %d", day)
	}

	if part < 1 || part > 2 {
		stopf("Expected part in rage [1, 2]. Got: %d", part)
	}

	// Resolve path
	file = path.Clean(file)
	fs, err := os.Stat(file)
	if err == nil {
		if fs.IsDir() {
			stopf("Input file \"%s\" is a directory. Expected a file.", file)
		}
		// Otherwise file exists and is a file
	} else if errors.Is(err, os.ErrNotExist) {
		stopf("Input file \"%s\" does not exists.", file)
	} else {
		// https://stackoverflow.com/a/12518877/2531987
		stopf("Input file \"%s\" is a Schrodinger file", file)
	}

	fmt.Println("Running AOC 2023 code")
	fmt.Println("Got day:", day)
	fmt.Println("Got part:", part)
	fmt.Println("Got file:", file)

	// Read input file
	data_bytes, err := os.ReadFile(file)
	if err != nil {
		stopf("Cannot read input file \"%s\". Error: %s", file, err)
	}

	if len(data_bytes) == 0 {
		stopf("Input file \"%s\" is empty", file)
	}

	data := string(data_bytes[:])

	if !is_ascii(data) {
		stopf("Input file \"%s\" contains non-ASCII characters", file)
	}

	// Split into lines
	lines := strings.Split(data, "\n")
	// lines = array_filter(lines, func(s string) bool { return len(s) > 0 })

	switch day {
	case 1:
		if part == 1 {
			day01.Main_1(lines)
		} else {
			day01.Main_2(lines)
		}
	default:
		stopf("Day %d is not implemented yet", day)
	}
}

// https://stackoverflow.com/a/53069799/2531987
func is_ascii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// https://stackoverflow.com/a/37563128/2531987
// Filter an array using a test. Elements passing the test are kept while those
// failing it are rejected.
func filter[T any](A []T, test func(T) bool) (ret []T) {
	for _, a := range A {
		if test(a) {
			ret = append(ret, a)
		}
	}
	return
}