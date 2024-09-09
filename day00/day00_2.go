package day00

import "fmt"

func main_2(lines []string, verbose bool) (n int, err error) {
	if verbose {
		fmt.Println("Hello from main_2")
		fmt.Printf("Got %d lines\n", len(lines))
		for _, line := range lines {
			fmt.Println(line)
		}
	}
	return 0, nil
}
