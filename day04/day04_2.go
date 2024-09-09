package day04

import (
	"aoc2023/utils"
)

func main_2(lines []string, verbose bool) (n int, err error) {
	var N_intersections []int = make([]int, len(lines))
	for i, line := range lines {
		nums, err := lineToNumbers(line)
		if err != nil {
			return -1, err
		}
		intersection := utils.ArrayArrayIntersection(nums.winning, nums.yours)
		N_intersections[i] = len(intersection)
	}

	copies := make([]int, len(N_intersections))
	for i := range copies {
		copies[i] = 1
	}

	for card_index, n := range N_intersections {
		for copy_index := 0; copy_index < copies[card_index]; copy_index++ {
			for int_index := 0; int_index < n; int_index++ {
				I := card_index + 1 + int_index
				if I <= len(copies) {
					copies[I] += 1
				} else {
					// Ignore indices that are out of range
				}
			}
		}
	}

	total_copies := utils.ArraySum(copies)

	return total_copies, nil
}
