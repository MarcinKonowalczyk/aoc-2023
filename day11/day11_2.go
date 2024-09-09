package day11

import (
	"aoc2023/utils"
)

func main_2(lines []string, verbose bool) (n int, err error) {
	u := parseLinesToUniverse(lines)

	distances := findDistances(u, 1_000_000)

	sum_distances := utils.MapReduceValues(distances, 0, func(a, b int) int {
		return a + b
	})

	return sum_distances, nil
}
