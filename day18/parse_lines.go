package day18

import (
	"fmt"
	"strconv"
	"strings"
)

type direction byte

const (
	RIGHT direction = iota
	DOWN
	LEFT
	UP
)

func (d direction) String() string {
	switch d {
	case RIGHT:
		return "R"
	case DOWN:
		return "D"
	case LEFT:
		return "L"
	case UP:
		return "U"
	default:
		return "INVALID"
	}
}

type rgb struct {
	r uint8
	g uint8
	b uint8
}
type line struct {
	dir      direction
	distance uint
	hexrgb   string
}

func (l line) String() string {
	return fmt.Sprintf("%v %d %s", l.dir, l.distance, l.hexrgb)
}

func parseLine(s string) (l line, err error) {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return l, fmt.Errorf("invalid line")
	}
	dir := parts[0]
	if len(dir) != 1 {
		return l, fmt.Errorf("invalid direction")
	}
	switch dir {
	case "R":
		l.dir = RIGHT
	case "D":
		l.dir = DOWN
	case "L":
		l.dir = LEFT
	case "U":
		l.dir = UP
	default:
		return l, fmt.Errorf("invalid direction")
	}
	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		return l, fmt.Errorf("invalid distance")
	}
	if distance < 0 {
		return l, fmt.Errorf("invalid distance")
	}
	l.distance = uint(distance)
	hexrgb := parts[2]
	if hexrgb[0] == '(' {
		hexrgb = hexrgb[1:]
	}
	if hexrgb[len(hexrgb)-1] == ')' {
		hexrgb = hexrgb[:len(hexrgb)-1]
	}
	if hexrgb[0] != '#' {
		return l, fmt.Errorf("invalid hex color")
	}
	hexrgb = hexrgb[1:]
	if len(hexrgb) != 6 {
		return l, fmt.Errorf("invalid hex color")
	}

	l.hexrgb = hexrgb

	return l, nil
}

// Extract the correct parameters from the line and set them in the output line
func extractCorrectParams(l line) (line, error) {
	if len(l.hexrgb) != 6 {
		return l, fmt.Errorf("invalid hex color")
	}
	hex_distance := l.hexrgb[:5]
	distance, err := strconv.ParseUint(hex_distance, 16, 32)
	l.distance = uint(distance)
	if err != nil {
		return l, err
	}
	hex_dir := l.hexrgb[5]
	switch hex_dir {
	case '0':
		l.dir = RIGHT
	case '1':
		l.dir = DOWN
	case '2':
		l.dir = LEFT
	case '3':
		l.dir = UP
	default:
		return l, fmt.Errorf("invalid direction")
	}

	return l, nil

}
