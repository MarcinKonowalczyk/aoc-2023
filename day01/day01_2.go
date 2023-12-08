package day01

import (
	"aoc2023/utils"
	"strings"
)

func main_2(lines []string) (n int, err error) {
	numbers := make([]int, len(lines))

	for i, line := range lines {
		this_line_numbers := line_to_numbers(line)
		first := this_line_numbers[0]
		last := this_line_numbers[len(this_line_numbers)-1]
		num := first*10 + last
		numbers[i] = num
	}

	n = utils.ArrayReduce(numbers, 0, func(a, b int) int { return a + b })

	return n, nil
}

var num_strings = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// scan the line one character at a time
// if the character is a number, add it to te list of numbers
// if the character is a letter, see if it makes a number
func line_to_numbers(line string) (nums []int) {
	for i, char := range line {
		if char >= '0' && char <= '9' {
			v := int(char - 48)
			nums = append(nums, v)
		} else {
			// check if the next characters make a number
			for j, num_str := range num_strings {
				if strings.HasPrefix(line[i:], num_str) {
					nums = append(nums, j+1)
				}
			}
		}
	}
	return nums
}
