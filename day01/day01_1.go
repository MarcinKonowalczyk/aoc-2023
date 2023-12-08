package day01

import (
	"aoc2023/utils"
	"errors"
	"fmt"
)

func Main(part int, lines []string) (n int, err error) {
	if part == 1 {
		return main_1(lines)
	} else if part == 2 {
		return main_2(lines)
	} else {
		return -1, errors.New("invalid part")
	}
}

func main_1(lines []string) (n int, err error) {
	numbers := make([]int, len(lines))

	for i, line := range lines {
		num, err := number_from_line(line)
		if err != nil {
			return -1, fmt.Errorf("line %d: %s", i, err)
		}
		numbers[i] = num
	}

	n = utils.ArrayReduce(numbers, 0, func(a, b int) int { return a + b })

	return n, nil
}

func number_from_line(line string) (int, error) {
	var first int = -1
	var last int = -1
	for _, char := range line {
		if char >= '0' && char <= '9' {
			v := int(char - 48)
			last = v
			if first == -1 {
				first = v
			}
		}
	}
	if first == -1 {
		return -1, errors.New("no first number")
	}
	if last == -1 {
		return -1, errors.New("no last number")
	}

	return first*10 + last, nil
}
