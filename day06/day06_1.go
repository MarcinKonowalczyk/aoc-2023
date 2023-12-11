package day06

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
	times, distances, err := parseLines(lines)
	if err != nil {
		return -1, err
	}
	fmt.Printf("times: %v\n", times)
	fmt.Printf("distances: %v\n", distances)
	return 0, nil
}

var TIMES_LINE_RE = regexp.MustCompile(`Time:(?P<times> *(\d+ *)+)`)

func parseTimesLine(line string) (times []int, err error) {
	matches := utils.GetNamedSubexpsCompiledRe(TIMES_LINE_RE, line)
	return utils.StringOfNumbersToNumbers(matches["times"])
}

var DISTANCES_LINE_RE = regexp.MustCompile(`Distance:(?P<distances> *(\d+ *)+)`)

func parseDistancesLine(line string) (distances []int, err error) {
	matches := utils.GetNamedSubexpsCompiledRe(DISTANCES_LINE_RE, line)
	return utils.StringOfNumbersToNumbers(matches["distances"])
}

func parseLines(lines []string) (times []int, distances []int, err error) {
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("invalid input")
	}
	times, err = parseTimesLine(lines[0])
	if err != nil {
		return nil, nil, err
	}
	distances, err = parseDistancesLine(lines[1])
	if err != nil {
		return nil, nil, err
	}
	return times, distances, nil
}
