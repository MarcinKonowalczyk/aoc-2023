package day06

import "fmt"

func main_2(lines []string) (n int, err error) {
	times, distances, err := parseLines(lines)
	if err != nil {
		return -1, err
	}
	fmt.Printf("times: %v\n", times)
	fmt.Printf("distances: %v\n", distances)
	return 0, nil
}
