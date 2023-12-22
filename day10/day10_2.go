package day10

import (
	"fmt"
)

func main_2(lines []string) (n int, err error) {
	pipe_map := parseLinesToPipeMap(lines)
	if pipe_map.IsEmpty() {
		return 0, fmt.Errorf("empty pipe map")
	}

	// fmt.Println(pipe_map)
	// fmt.Println("start at", pipe_map.x, pipe_map.y)

	directions := walkFullCircle(&pipe_map)

	// fmt.Println("directions:", directions)

	furthest := len(directions) / 2

	return furthest, nil
}
