package day08

import (
	"aoc2023/utils"
	"fmt"
	"sort"
)

func main_2(lines []string, verbose bool) (n int, err error) {
	directions, err := parseDirections(lines[0])

	if err != nil {
		return 0, err
	}

	if lines[1] != "" {
		return 0, fmt.Errorf("invalid file. expected blank line after directions")
	}

	nodes_map, err := nodeLinesParser(lines[2:])
	if err != nil {
		return 0, err
	}

	starting_nodes := getStartingNodes(nodes_map)

	periods := make([]int, 0)
	for _, node := range starting_nodes {
		period, err := findPeriod(nodes_map, directions, node)
		if err != nil {
			return 0, err
		}
		periods = append(periods, period)
	}

	// We will need to walk the multiple of each cycle's period to get to the end node
	// Find the lowest common multiple of all the periods
	lcm_periods := utils.ArrayReduce(periods, 1, utils.LCM)

	return lcm_periods, nil
}

func getStartingNodes(nodes map[string]Node) []string {
	startingNodes := []string{}
	for k := range nodes {
		if k[len(k)-1] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}
	sort.Strings(startingNodes)
	return startingNodes
}

func isEndNode(node string) bool {
	return node[len(node)-1] == 'Z'
}

const N_offsets = 10

func findEndNodeOffsets(nodes map[string]Node, directions []bool, current_node string) []int {
	directions_chan := make(chan EnumeratedDirection)
	go directionCycler(directions, directions_chan)
	defer close(directions_chan)

	end_node_offsets := make([]int, 0)

	for j := 0; true; j++ {
		if isEndNode(current_node) {
			end_node_offsets = append(end_node_offsets, j)
			if len(end_node_offsets) >= N_offsets {
				break
			}
		}

		direction := <-directions_chan
		// Move to the next node
		current_node = nextNode(nodes, current_node, direction)
	}

	return end_node_offsets
}

// Find the period of the cycle starting at the given node
func findPeriod(nodes_map map[string]Node, directions []bool, starting_node string) (int, error) {
	end_node_offsets := findEndNodeOffsets(nodes_map, directions, starting_node)
	end_node_offsets_diff, err := utils.ArrayDiff(end_node_offsets, 1)
	if err != nil {
		return -1, err
	}

	first_end_node_offset := end_node_offsets[0]

	// Make sure all the end node offsets are the same. This is not necessarily true for the general case,
	// but it is true for the input and it makes the problem easier.
	if !utils.ArrayAll(end_node_offsets_diff, func(n int) bool { return n == first_end_node_offset }) {
		return -1, fmt.Errorf("end node offsets are not the same")
	}

	return first_end_node_offset, nil
}
