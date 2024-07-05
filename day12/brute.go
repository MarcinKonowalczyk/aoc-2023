package day12

import (
	"aoc2023/utils"
)

func blocksToSprings(blocks [][]Spring) []Spring {
	springs := make([]Spring, 0)
	for i, block := range blocks {
		springs = append(springs, block...)
		if i < len(blocks)-1 {
			springs = append(springs, OPERATIONAL)
		}
	}
	return springs
}
func recursiveBruteForce(l Line, depth int) (n_valid int) {

	springs := blocksToSprings(l.blocks)
	if !utils.ArrayContains(springs, UNKNOWN) {
		// We have no unknown springs. We're done
		return 1
	}

	// Put in a guess for the first unknown
	for guess := OPERATIONAL; guess <= DAMAGED; guess++ {
		test_line := l.copy()
		test_line_springs := blocksToSprings(test_line.blocks)
		for i, spring := range test_line_springs {
			if spring == UNKNOWN {
				test_line_springs[i] = guess
				break
			}
		}
		test_line.blocks = splitSprings(test_line_springs, OPERATIONAL)
		is_valid := checkLine(test_line)
		// fmt.Println(l.blocks, "->", test_line.blocks, "is valid:", is_valid)
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

// Check whether a line is valid
func checkLine(l Line) bool {
	springs := make([]Spring, 0)
	for _, block := range l.blocks {
		springs = append(springs, block...)
		springs = append(springs, OPERATIONAL)
	}
	if utils.ArrayContains(springs, UNKNOWN) {
		// We have unknown springs
		return true
	} else {
		// We have no unknown springs. We can check if the line is valid
		// fmt.Println(" Checking line", springsString(springs))
		// fmt.Println(" Blocks", l.blocks)
		// fmt.Println(" Groups", l.groups)
		if len(l.blocks) != len(l.groups) {
			return false
		}
		for i, block := range l.blocks {
			if len(block) != l.groups[i] {
				return false
			}
		}
		return true
	}
}
