package day04

import (
	"aoc2023/utils"
	"fmt"
	"regexp"
)

func Main(part int, lines []string) (n int, err error) {
	if part == 1 {
		return main_1(lines)
	} else if part == 2 {
		return main_2(lines)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string) (n int, err error) {

	overall_score := 0
	for _, line := range lines {
		// fmt.Println(line)
		nums, err := lineToNumbers(line)
		// fmt.Println(nums)
		if err != nil {
			return -1, err
		}
		intersection := utils.ArrayArrayIntersection(nums.winning, nums.yours)
		// fmt.Println(intersection)
		score := intersectionToScore(intersection)
		// fmt.Println(score)
		overall_score += score
	}
	return overall_score, nil
}

type numbers struct {
	winning []int
	yours   []int
}

var lineRe = regexp.MustCompile(`Card +(?P<Card>\d+): +(?P<Winning>[^\|]+) ?\| ?(?P<Yours>[^\|]+)(?P<Rest>.*)`)

func lineToNumbers(line string) (numbers, error) {

	match := utils.GetParamsCompiledRe(lineRe, line)
	if match["Rest"] != "" {
		return numbers{}, fmt.Errorf("invalid line %s", line)
	}

	winning_string := match["Winning"]
	winning, err := utils.StringOfNumbersToNumbers(winning_string)
	if err != nil {
		return numbers{}, err
	}

	yours_string := match["Yours"]
	yours, err := utils.StringOfNumbersToNumbers(yours_string)
	if err != nil {
		return numbers{}, err
	}

	return numbers{
		winning: winning,
		yours:   yours,
	}, nil
}

func intersectionToScore(intersection []int) int {
	if len(intersection) == 0 {
		return 0
	} else {
		return 1 << (len(intersection) - 1)
	}
}

// 14139 too low
