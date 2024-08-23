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
	color    rgb
}

func (l line) String() string {
	return fmt.Sprintf("%v %d (%d,%d,%d)", l.dir, l.distance, l.color.r, l.color.g, l.color.b)
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
	rgbhex := parts[2]
	if rgbhex[0] == '(' {
		rgbhex = rgbhex[1:]
	}
	if rgbhex[len(rgbhex)-1] == ')' {
		rgbhex = rgbhex[:len(rgbhex)-1]
	}

	l.color, err = parseColor(rgbhex)
	if err != nil {
		return l, err
	}

	return l, nil
}

func parseColor(s string) (c rgb, err error) {
	if s[0] != '#' {
		return c, fmt.Errorf("invalid hex")
	}
	if len(s) != 7 {
		return c, fmt.Errorf("invalid hex")
	}
	s = s[1:]
	r := s[:2]
	g := s[2:4]
	b := s[4:]
	r_int, err := strconv.ParseInt(r, 16, 64)
	if err != nil {
		return c, err
	}
	if r_int < 0 || r_int > 255 {
		return c, fmt.Errorf("invalid r")
	}
	c.r = uint8(r_int)
	g_int, err := strconv.ParseInt(g, 16, 64)
	if g_int < 0 || g_int > 255 {
		return c, fmt.Errorf("invalid g")
	}
	if err != nil {
		return c, err
	}
	c.g = uint8(g_int)
	b_int, err := strconv.ParseInt(b, 16, 64)
	if b_int < 0 || b_int > 255 {
		return c, fmt.Errorf("invalid b")
	}
	if err != nil {
		return c, err
	}
	c.b = uint8(b_int)
	return c, nil
}
