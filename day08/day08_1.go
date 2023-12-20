package day08

import (
	"aoc2023/utils"
	"fmt"
	"regexp"
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

const START_NODE string = "AAA"
const END_NODE string = "ZZZ"

const MAX_STEPS int = 10_000_000

// const MAX_STEPS int = 10_000_000_000

func main_1(lines []string) (n int, err error) {
	directions, err := parseDirections(lines[0])

	if err != nil {
		return 0, err
	}

	if lines[1] != "" {
		return 0, fmt.Errorf("invalid file. expected blank line after directions")
	}

	nodes, err := nodeLinesParser(lines[2:])
	if err != nil {
		return 0, err
	}

	node := START_NODE
	i := 0
	directions_chan := make(chan EnumeratedDirection)
	go directionCycler(directions, directions_chan)
	for {
		node = nextNode(nodes, node, <-directions_chan)
		i++
		if node == END_NODE {
			break
		}
		if i > MAX_STEPS {
			return 0, fmt.Errorf("max steps exceeded")
		}
	}

	return i, nil
}

func parseDirections(line string) ([]bool, error) {
	split := strings.Split(line, "")
	return utils.ArrayMapWithError(split, func(s string) (bool, error) {
		if s == "R" {
			return true, nil
		} else if s == "L" {
			return false, nil
		} else {
			return false, fmt.Errorf("invalid direction: %s", s)
		}
	})
}

type EnumeratedDirection struct {
	right bool
	index int
}

// Cycle through the directions infinitely
func directionCycler(directions []bool, out chan EnumeratedDirection) {
	defer func() { recover() }() // Recover from panic when out is closed
	for {
		for index, right := range directions {
			out <- EnumeratedDirection{
				right: right,
				index: index,
			}
		}
	}
}

var NODE_LINE_RE = regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)

type Node struct {
	left, right string
}

func nodeLinesParser(lines []string) (map[string]Node, error) {
	nodes_map := make(map[string]Node)
	for _, line := range lines {
		matches := NODE_LINE_RE.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		nodes_map[matches[1]] = Node{matches[2], matches[3]}
	}
	return nodes_map, nil
}

func nextNode(nodes map[string]Node, node string, direction EnumeratedDirection) string {
	if direction.right {
		node = nodes[node].right
	} else {
		node = nodes[node].left
	}
	return node
}
