package day15

import (
	"fmt"
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
	if len(lines) != 1 {
		return -1, fmt.Errorf("invalid input")
	}
	steps := strings.Split(lines[0], ",")

	sum := 0
	for _, step := range steps {
		h := hash(step)
		// fmt.Printf("Step %d: %s, hash: %d\n", i, step, h)
		sum += int(h)
	}
	return sum, nil
}

func hash(s string) uint8 {
	c := 0
	for _, r := range s {
		c += int(r)
		c *= 17
		c %= 256
	}
	if c < 0 || c > 255 {
		// Sanity check
		panic("invalid hash")
	}
	return uint8(c)
}
