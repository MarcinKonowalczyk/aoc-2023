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
		parsed_line, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		parsed_lines[line_index] = parsed_line
	}
	for _, parsed_line := range parsed_lines {
		fmt.Println(parsed_line)
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
	for _, group := range l.groups {
		s += fmt.Sprintf("%d", group)
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
