package day21

import (
	"fmt"
)

func Main(part int, lines []string, verbose bool) (n int, err error) {
	if part == 1 {
		return main_1(lines, verbose)
	} else if part == 2 {
		return main_2(lines, verbose)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string, verbose bool) (n int, err error) {
	g, err := parseLines(lines)
	if err != nil {
		return 0, err
	}

	if verbose {
		fmt.Println(g)
	}

	return 0, nil
}
