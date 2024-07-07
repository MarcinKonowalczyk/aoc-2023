package day12

import (
	"fmt"
	"strconv"
	"strings"
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
	parsed_lines := make([]Line, len(lines))
	for line_index, line := range lines {
		parsed_line, err := parseInputLine(line)
		if err != nil {
			return -1, err
		}
		parsed_lines[line_index] = parsed_line
	}

	sum_counts := 0
	// for _, parsed_line := range parsed_lines {
	// 	c := recursiveBruteForce(parsed_line, 0)
	// 	fmt.Println(parsed_line, "->", c)
	// 	sum_counts += c
	// }

	for _, parsed_line := range parsed_lines {
		c := recursiveStepFromLeft(parsed_line, 0)
		// fmt.Println(parsed_line, "->", c)
		sum_counts += c
	}

	// fmt.Println(parsed_lines[len(parsed_lines)-1])
	// c := recursiveStepFromLeft(parsed_lines[len(parsed_lines)-1], 0)
	// fmt.Println(c)

	// c2 := recursiveBruteForce(parsed_lines[len(parsed_lines)-1], 0)
	// fmt.Println(c2)

	return sum_counts, nil
}

func parseInputLine(line string) (Line, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return Line{}, fmt.Errorf("invalid line: %s", line)
	}
	springs_string := parts[0]
	springs, err := parseSprings(springs_string)
	if err != nil {
		return Line{}, err
	}
	groups_string := parts[1]
	groups, err := parseGroups(groups_string)
	if err != nil {
		return Line{}, err
	}
	blocks := splitSprings(springs, OPERATIONAL)
	l := Line{springs_string, blocks, groups}
	return l, nil
}

func parseSprings(springs_string string) ([]Spring, error) {
	springs := make([]Spring, len(springs_string))
	for i, c := range springs_string {
		spring, err := fromString(string(c))
		if err != nil {
			return nil, fmt.Errorf("invalid spring: %s", string(c))
		}
		springs[i] = spring
	}
	return springs, nil
}

func parseGroups(groups_string string) ([]int, error) {
	groups := strings.Split(groups_string, ",")
	groups_int := make([]int, len(groups))
	for i, group := range groups {
		group_int, err := strconv.Atoi(group)
		if err != nil {
			return nil, fmt.Errorf("invalid group: %s", group)
		}
		groups_int[i] = group_int
	}
	return groups_int, nil
}

// Split the springs block into groups of springs at a given spring
func splitSprings(springs []Spring, at Spring) [][]Spring {
	blocks := make([][]Spring, 0)
	block := make([]Spring, 0)
	for _, spring := range springs {
		if spring == at {
			if len(block) > 0 {
				// We've just finished a block
				blocks = append(blocks, block)
				block = make([]Spring, 0)
			}
		} else {
			block = append(block, spring)
		}
	}
	if len(block) > 0 {
		blocks = append(blocks, block)
	}
	return blocks
}

func stepFromLeft(l Line) (ll []Line, g []Spring, end bool) {

	if len(l.blocks) == 0 {
		// No more blocks, so we should have no more groups
		if len(l.groups) == 0 {
			end = true
			return
		}
		return
	}

	if len(l.groups) == 0 {
		lc2 := l.copy()
		if lc2.blocks[0][0] == DAMAGED {
			// First block is DAMAGED, so we can't do anything
			return
		} else if lc2.blocks[0][0] == UNKNOWN {
			lc2.blocks[0] = lc2.blocks[0][1:]
			if len(lc2.blocks[0]) == 0 {
				lc2.blocks = lc2.blocks[1:]
			}
			ll = append(ll, lc2)
			g = append(g, OPERATIONAL)
			return
		}
	}

	if len(l.blocks[0]) == 0 {
		panic("len(l.blocks[[0]]) == 0")
	}

	lc1 := l.copy()

	// This is for both DAMAGED and UNKNOWN first block
	ok := true
	for {
		if lc1.groups[0] == 0 {
			panic("lc1.groups[0] == 0")
		}
		if lc1.groups[0] > 1 {
			// First group is at least 2
			lc1.groups[0] -= 1
			lc1.blocks[0] = lc1.blocks[0][1:]
			if len(lc1.blocks[0]) == 0 {
				// We've exhausted a block, but we still have a group to go
				// This is not ok
				ok = false
				break
			} else {
				// We still have both a group and a block
				// So we can continue
			}
		} else {
			// First element of first group is 1. This should be the end if this group
			lc1.groups = lc1.groups[1:]
			if len(lc1.blocks[0]) == 1 {
				// If the spring is DAMAGED, thats ok. If its UNKNOWN thats also ok
				lc1.blocks[0] = lc1.blocks[0][1:]
			} else {
				// This is not the last bit of this block
				next_spring := lc1.blocks[0][1]
				if next_spring == UNKNOWN {
					// Next spring be UNKNOWN. Thats ok.
					lc1.blocks[0] = lc1.blocks[0][2:]
				} else {
					// Next spring is DAMAGED. Thats not ok because we've jsut finished a group
					ok = false
					break
				}
			}

			if len(lc1.blocks[0]) == 0 {
				// We've finished with this block. Remove it
				lc1.blocks = lc1.blocks[1:]
			}

			if len(lc1.groups) == 0 {
				// We've finished all groups
				if len(lc1.blocks) == 0 {
					// We've finished all groups and all blocks
					// This is the end
					break
				} else {
					// We've finished all groups, but there are still blocks left
					// This is ok only if they're filled with UNKNOWN
					for _, block := range lc1.blocks {
						for _, spring := range block {
							if spring != UNKNOWN {
								ok = false
								goto _break
							}
						}
					}
				_break:
					break
				}
			}
			break
			// if len(lc1.groups) == 0 && len(lc1.blocks) > 0 {
			// 	// This was the last group, but there are still blocks left
			// 	ok = false
			// }
		}
		// if len(lc1.blocks[0]) == 0 {
		// 	// First block is empty, so we can remove it
		// 	lc1.blocks = lc1.blocks[1:]
		// 	if len(lc1.blocks) == 0 {
		// 		_continue = false
		// 	}
		// }
	}
	if ok {
		ll = append(ll, lc1)
		g = append(g, DAMAGED)
	}

	// If the first block is UNKNOWN, we can also try to make it OPERATIONAL
	if l.blocks[0][0] == UNKNOWN {
		lc2 := l.copy()
		if len(lc2.blocks[0]) > 1 {
			lc2.blocks[0] = l.blocks[0][1:]
		} else {
			lc2.blocks = lc2.blocks[1:]
		}
		ll = append(ll, lc2)
		g = append(g, OPERATIONAL)
	} else if l.blocks[0][0] != DAMAGED {
		panic(fmt.Sprintf("invalid block_0: %s", l.blocks[0][0]))
	}

	return
}

func recursiveStepFromLeft(l Line, depth int) (c int) {
	// ll, g, end := stepFromLeft(l)
	ll, _, end := stepFromLeft(l)
	// pad_space := strings.Repeat(" ", depth+1)
	// pad_underscore := strings.Repeat("_", depth+1)
	if end {
		// fmt.Println(pad_space + utils.Csprintf(utils.Green, "END"))
		return 1
	} else if len(ll) == 0 {
		// fmt.Println(pad_space + utils.Csprintf(utils.Yellow, "No"))
	}
	if end {
		return 1
	}
	// for i, line := range ll {
	// 	fmt.Println(pad_space + " " + g[i].String() + " " + line.String())
	// }
	for _, line := range ll {
		c += recursiveStepFromLeft(line, depth+1)
	}
	// for i, line := range ll {
	// 	fmt.Println(pad_underscore + " " + g[i].String() + " " + line.String())
	// 	c += recursiveStepFromLeft(line, depth+1)
	// }
	return
}
