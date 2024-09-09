package day09

import (
	"aoc2023/utils"
	"fmt"
)

func Main(part int, lines []string, verbose bool) (n int, err error) {
	if part == 1 {
		return main_1(lines, verbose)
	} else if part == 2 {
		return main_2(lines, verbose)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string, verbose bool) (n int, err error) {
	extrapolated_values := make([]int, len(lines))
	for line_index, line := range lines {
		values, err := utils.StringOfNumbersToInts(line)
		if err != nil {
			return 0, err
		}
		diffs, err := valuesToDiffs(values)
		if err != nil {
			return 0, err
		}
		extrapolateDiffsForward(diffs)
		extrapolated_value := diffs[0][len(diffs[0])-1]
		extrapolated_values[line_index] = extrapolated_value
	}

	sum := utils.ArraySum(extrapolated_values)
	return sum, nil
}

func valuesToDiffs(values []int) (diffs map[int][]int, err error) {
	diffs = make(map[int][]int)
	diffs[0] = values

	if utils.ArrayAll(values, func(n int) bool { return n == 0 }) {
		return diffs, nil
	}

	for j := 1; true; j++ {
		diff, err := utils.ArrayDiff(values, 1)
		if err != nil {
			return nil, err
		}
		diffs[j] = diff
		values = diff
		if utils.ArrayAll(diff, func(n int) bool { return n == 0 }) {
			break
		}
	}

	return diffs, nil
}

// Extrapolate the diffs one step further starting from the last diff which is all 0s
func extrapolateDiffsForward(diffs map[int][]int) {
	N_diffs := len(diffs)
	if N_diffs == 0 {
		// No diffs
		return
	}
	diffs[N_diffs-1] = append(diffs[N_diffs-1], 0)
	if N_diffs == 1 {
		// Only one diff (the original values)
		return
	}
	for i := N_diffs - 2; i >= 0; i-- {
		row := diffs[i]
		next_row := diffs[i+1]
		new_value := row[len(row)-1] + next_row[len(next_row)-1]
		diffs[i] = append(row, new_value)
	}
}
