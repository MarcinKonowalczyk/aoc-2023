package day09

import (
	"aoc2023/utils"
)

func main_2(lines []string, verbose bool) (n int, err error) {
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
		extrapolateDiffsBackward(diffs)
		extrapolated_value := diffs[0][0]
		extrapolated_values[line_index] = extrapolated_value
	}

	sum := utils.ArraySum(extrapolated_values)
	return sum, nil
}

// Extrapolate the diffs one step further starting from the last diff which is all 0s
func extrapolateDiffsBackward(diffs map[int][]int) {
	N_diffs := len(diffs)
	if N_diffs == 0 {
		// No diffs
		return
	}
	// Prepend a 0 to the last diff
	diffs[N_diffs-1] = append([]int{0}, diffs[N_diffs-1]...)
	if N_diffs == 1 {
		// Only one diff (the original values)
		return
	}
	for i := N_diffs - 2; i >= 0; i-- {
		row := diffs[i]
		next_row := diffs[i+1]
		new_value := row[0] - next_row[0]
		diffs[i] = append([]int{new_value}, diffs[i]...)
	}
}
