package day03

import "fmt"

func Main_1(lines []string) (n int, err error) {
	for _, line := range lines {
		fmt.Println(line)
	}
	return 0, nil
}
