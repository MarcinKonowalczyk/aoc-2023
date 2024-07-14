package day13

import (
	"fmt"
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
	maps, err := parseLines(lines)
	if err != nil {
		return -1, err
	}
	fmt.Println("Found maps:", len(maps))
	for _, m := range maps {
		fmt.Println(m)
	}
	return 0, nil
}

type Map [][]bool

func (m Map) nRows() int {
	return len(m)
}

func (m Map) nCols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Map) String() string {
	s := "Map:\n"
	for _, row := range m {
		for _, cell := range row {
			if cell {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func parseLines(lines []string) ([]Map, error) {
	var maps []Map = make([]Map, 0)
	var m Map = make(Map, 0)
	for _, line := range lines {
		if line == "" {
			if m.nRows() > 0 {
				maps = append(maps, m)
				m = make(Map, 0)
			}
		} else {
			row := make([]bool, 0)
			for _, c := range line {
				if c == '.' {
					row = append(row, false)
				} else if c == '#' {
					row = append(row, true)
				} else {
					return nil, fmt.Errorf("invalid character: %c", c)
				}
			}
			m = append(m, row)
		}
	}
	if m.nRows() > 0 {
		maps = append(maps, m)
	}
	return maps, nil
}
