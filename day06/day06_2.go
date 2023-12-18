package day06

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

func main_2(lines []string) (n int, err error) {
	time, distance, err := parseLinesRemoveSpaces(lines)
	if err != nil {
		return -1, err
	}
	fmt.Printf("time: %v\n", time)
	fmt.Printf("distance: %v\n", distance)

	min_time, max_time, err := calcHoldTimeRange(time, distance)
	if err != nil {
		return -1, err
	}
	N_times := max_time - min_time + 1

	return N_times, nil
}

func parseTimesLineRemoveSpaces(line string) (time int, err error) {
	matches := utils.GetNamedSubexpsCompiledRe(TIMES_LINE_RE, line)
	times_without_spaces := strings.ReplaceAll(matches["times"], " ", "")
	numbers, err := utils.StringOfNumbersToInts(times_without_spaces)
	if err != nil {
		return -1, err
	}
	if len(numbers) != 1 {
		return -1, fmt.Errorf("invalid input")
	}
	return numbers[0], nil
}

func parseDistancesLineRemoveSpaces(line string) (distance int, err error) {
	matches := utils.GetNamedSubexpsCompiledRe(DISTANCES_LINE_RE, line)
	distances_without_spaces := strings.ReplaceAll(matches["distances"], " ", "")
	numbers, err := utils.StringOfNumbersToInts(distances_without_spaces)
	if err != nil {
		return -1, err
	}
	if len(numbers) != 1 {
		return -1, fmt.Errorf("invalid input")
	}
	return numbers[0], nil
}

func parseLinesRemoveSpaces(lines []string) (time int, distance int, err error) {
	if len(lines) != 2 {
		return -1, -1, fmt.Errorf("invalid input")
	}
	time, err = parseTimesLineRemoveSpaces(lines[0])
	if err != nil {
		return -1, -1, err
	}
	distance, err = parseDistancesLineRemoveSpaces(lines[1])
	if err != nil {
		return -1, -1, err
	}
	return time, distance, nil
}
