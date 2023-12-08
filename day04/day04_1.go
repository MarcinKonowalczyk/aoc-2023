package day04

import "fmt"

func Main_1(lines []string) (n int, err error) {
	fmt.Println("Hello from main_1")
	fmt.Printf("Got %d lines\n", len(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
	return 0, nil
}
