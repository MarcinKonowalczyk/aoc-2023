package day12

import (
	"aoc2023/utils"
)

func recursiveBruteForce(l Line, depth int) (n_valid int) {
	if !utils.ArrayContains(l.springs, UNKNOWN) {
		// We dont have unknowns. We're done
		return 1
	}

	// Put in a guess for the first unknown
	for guess := OPERATIONAL; guess <= DAMAGED; guess++ {
		line_copy := make([]Spring, len(l.springs))
		copy(line_copy, l.springs)
		test_line := Line{line_copy, l.groups}
		for i, spring := range test_line.springs {
			if spring == UNKNOWN {
				test_line.springs[i] = guess
				break
			}
		}
		is_valid := checkLine(test_line)
		// fmt.Println(springsString(l.springs), "->", springsString(test_line.springs), "is valid:", is_valid)
		if is_valid {
			// We have a valid line
			n_valid += recursiveBruteForce(test_line, depth+1)
		} else {
			// We have an invalid line
			continue
		}
	}
	return
}
