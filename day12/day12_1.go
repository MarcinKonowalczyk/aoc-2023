package day12

import (
	"aoc2023/utils"
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
		parsed_line, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		parsed_lines[line_index] = parsed_line
	}
	for _, parsed_line := range parsed_lines {
		fmt.Println(parsed_line)
		reduced_line := reduceBlocks(parsed_line)
		fmt.Println(reduced_line)
	}
	return 0, nil
}

type Spring int

const (
	OPERATIONAL Spring = iota
	DAMAGED
	UNKNOWN
)

func (s Spring) String() string {
	switch s {
	case OPERATIONAL:
		return "."
	case DAMAGED:
		return "#"
	case UNKNOWN:
		return "?"
	default:
		panic("invalid spring (1)")
	}
}

func fromString(s string) (Spring, error) {
	switch s {
	case ".":
		return OPERATIONAL, nil
	case "#":
		return DAMAGED, nil
	case "?":
		return UNKNOWN, nil
	default:
		return UNKNOWN, fmt.Errorf("invalid spring (2)")
	}
}

type Line struct {
	springs []Spring
	groups  []int
}

func (l Line) String() string {
	s := ""
	for _, spring := range l.springs {
		s += spring.String()
	}
	s += " "
	for i, group := range l.groups {
		s += fmt.Sprintf("%d", group)
		if i < len(l.groups)-1 {
			s += ","
		}
	}
	return s
}

func parseLine(line string) (Line, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return Line{}, fmt.Errorf("invalid line: %s", line)
	}
	springs := make([]Spring, len(parts[0]))
	for i, c := range parts[0] {
		spring, err := fromString(string(c))
		if err != nil {
			return Line{}, fmt.Errorf("invalid spring: %s", string(c))
		}
		springs[i] = spring
	}
	parts_2 := strings.Split(parts[1], ",")
	groups := make([]int, len(parts_2))
	for i, part := range parts_2 {
		int_part, err := strconv.Atoi(part)
		if err != nil {
			return Line{}, fmt.Errorf("invalid group: %s", part)
		}
		groups[i] = int_part
	}
	return Line{springs, groups}, nil
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

// Remove
func reduceBlocks(l Line) Line {
	blocks := splitSprings(l.springs, OPERATIONAL)
	knowns := make([]bool, len(blocks))
	lengths := make([]int, len(blocks))

	groups := l.groups

	for i, block := range blocks {
		lengths[i] = len(block)
		knowns[i] = !utils.ArrayContains(block, UNKNOWN)
	}

	if len(knowns) != len(blocks) {
		panic("invalid length")
	}

	// Going from left to right, reduce fully known blocks
	n_to_reduce := 0
	for i := range groups {
		known := knowns[i]
		if !known {
			// We can't carry on redusing in this direction or we might get off sync
			break
		}
		if groups[i] != lengths[i] {
			panic("invalid group")
		}
		n_to_reduce++
	}

	if n_to_reduce > 0 {
		blocks = blocks[n_to_reduce:]
		knowns = knowns[n_to_reduce:]
		lengths = lengths[n_to_reduce:]
		groups = groups[n_to_reduce:]
	}

	// Going from right to left, reduce fully known blocks
	n_to_reduce = 0
	for i := range groups {
		known := knowns[len(knowns)-1-i]
		if !known {
			// We can't carry on redusing in this direction or we might get off sync
			break
		}
		if groups[len(groups)-1-i] != lengths[len(lengths)-1-i] {
			panic("invalid group")
		}
		// We can reduce this block
		n_to_reduce++
	}

	if n_to_reduce > 0 {
		blocks = blocks[:len(blocks)-n_to_reduce]
		knowns = knowns[:len(knowns)-n_to_reduce]
		lengths = lengths[:len(lengths)-n_to_reduce]
		groups = groups[:len(groups)-n_to_reduce]
	}

	// fmt.Println("blocks", blocks)
	// fmt.Println("knowns", knowns)
	// fmt.Println("lengths", lengths)
	// fmt.Println("groups", groups)

	// Join the blocks back together to form the new springs
	new_springs := make([]Spring, 0)
	for i, block := range blocks {
		new_springs = append(new_springs, block...)
		if i < len(blocks)-1 {
			new_springs = append(new_springs, OPERATIONAL)
		}
	}

	l.springs = new_springs
	l.groups = groups

	return l
}
