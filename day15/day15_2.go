package day15

import "fmt"

func main_2(lines []string) (n int, err error) {
	fmt.Println("Hello from main_2")
	fmt.Printf("Got %d lines\n", len(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
	return 0, nil
}
