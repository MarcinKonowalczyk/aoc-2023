package day00

import "fmt"

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
	if verbose {
		fmt.Println("Hello from main_1")
		fmt.Printf("Got %d lines\n", len(lines))
		for _, line := range lines {
			fmt.Println(line)
		}
	}
	return 0, nil
}
