package day10

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
	pipe_map, startx, starty := parseLinesToPipeMap(lines)
	if pipe_map.IsEmpty() {
		return 0, fmt.Errorf("empty pipe map")
	}

	fmt.Println(pipe_map)
	fmt.Println("start at", startx, starty)

	return 0, nil
}

type pipe rune

const (
	VERTICAL   pipe = '|'
	HORIZONTAL pipe = '-'
	NORTH_EAST pipe = 'L'
	NORTH_WEST pipe = 'J'
	SOUTH_EAST pipe = '7'
	SOUTH_WEST pipe = 'F'
	GROUND     pipe = '.'
	START      pipe = 'S'
)

var pipes = []pipe{VERTICAL, HORIZONTAL, NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST, GROUND, START}

func isPipe(r rune) bool {
	for _, p := range pipes {
		if r == rune(p) {
			return true
		}
	}
	return false
}

type PipeMap struct {
	m      [][]pipe
	width  int
	height int
}

func (m PipeMap) IsEmpty() bool {
	return m.width == 0 || m.height == 0
}

func (m PipeMap) Get(x, y int) pipe {
	return m.m[y][x]
}

func (m PipeMap) String() string {
	s := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			s += string(m.m[y][x])
		}
		if y < m.height-1 {
			s += "\n"
		}
	}
	return s
}

func parseLinesToPipeMap(lines []string) (pipe_map PipeMap, startx, starty int) {
	height := len(lines)
	if height == 0 {
		return PipeMap{}, 0, 0
	}
	width := len(lines[0])
	if width == 0 {
		return PipeMap{}, 0, 0
	}

	// Allocate the map
	pipe_map = PipeMap{make([][]pipe, height), width, height}
	for y := 0; y < height; y++ {
		pipe_map.m[y] = make([]pipe, width)
	}

	// Parse the map
	for y, line := range lines {
		if len(line) != width {
			panic(fmt.Errorf("line %d has length %d, expected %d", y, len(line), width))
		}
		for x, r := range line {
			if !isPipe(r) {
				panic(fmt.Errorf("invalid character %c at (%d,%d)", r, x, y))
			}
			p := pipe(r)
			pipe_map.m[y][x] = p
			if p == START {
				startx = x
				starty = y
			}
		}
	}

	return pipe_map, startx, starty
}
