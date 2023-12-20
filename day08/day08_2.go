package day08

import (
	"aoc2023/utils"
	"fmt"
	"sort"
)

func main_2(lines []string) (n int, err error) {
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

	current_nodes := getStartingNodes(nodes_map)
	fmt.Printf("starting_nodes: %v\n", current_nodes)

	// NOTE: After len(nodes_map) steps, we will have visited all the nodes at least once.
	max_steps := len(nodes_map) * 2
	fmt.Printf("max_steps: %d\n", max_steps)

	i := 0
	directions_chan := make(chan EnumeratedDirection)
	go directionCycler(directions, directions_chan)
	for {
		direction := <-directions_chan
		if direction.right {
			for i, node := range current_nodes {
				current_nodes[i] = nodes_map[node].right
			}
		} else {
			for i, node := range current_nodes {
				current_nodes[i] = nodes_map[node].left
			}
		}
		i++
		if areAllEndNodes(current_nodes) {
			break
		}
		if i > max_steps {
			break
		}
	}

	// Failed to find a solution within max_steps. Try again with a more complicated cycle approach.
	if !areAllEndNodes(current_nodes) {
		current_nodes := getStartingNodes(nodes_map)
		cycles := []Cycle{}
		for _, node := range current_nodes {
			cycles = append(cycles, findCycle(nodes_map, directions, node))
		}

		// Align the cycles so that they all start at the same point
		alignCycles(cycles)

		// Run checkCycle again to make sure we've got it right
		for i, cycle := range cycles {
			err := checkCycle(nodes_map, directions, current_nodes[i], cycle)
			if err != nil {
				panic(err)
			}
		}

		shortest_cycle_length := walkToEndNodeOfShortestCycle(cycles)
		j := 0
		for {
			fmt.Printf("j: %d\n", j)
			fmt.Printf("cycles: %v\n", cycles)
			fmt.Printf("shortest_cycle_length: %d\n", shortest_cycle_length)
			// Calculate the steps until the next alignment for each cycle
			next_alignments := []int{}
			for _, cycle := range cycles {
				fmt.Printf("cycle: %v\n", cycle)
				next_alignments = append(next_alignments, calcNextEndStateAlignment(cycle, shortest_cycle_length))
			}

			fmt.Printf("next_alignments: %v\n", next_alignments)
			if j > 1 {
				panic("stop")
			}

			// Find the next alignment that occurs the soonest but is not 0
			largest, _, err := utils.MaxArrayFunc(next_alignments, func(a, b int) bool {
				if a == 0 {
					return false
				}
				if b == 0 {
					return true
				}
				return a > b
			})

			if err != nil {
				panic(err)
			}

			// Walk all the cycles forward by the largest number of steps
			walkCycles(cycles, largest)

			// Check if we're done
			if areAllCyclesAtEndNode(cycles) {
				break
			}

			// To find the next alignment, we need to find the cycle with the shortest cycle_steps
			// Find cycle_steps of the cycles that are at the end node
			cycle_steps := []int{}
			for _, cycle := range cycles {
				if cycle.end_node_steps == 0 {
					cycle_steps = append(cycle_steps, cycle.cycle_steps)
				}
			}

			// Append the largest number of steps to the cycle_steps
			cycle_steps = append(cycle_steps, largest)

			fmt.Printf("cycle_steps: %v\n", cycle_steps)

			// Find the lowest common multiple of the cycle_steps
			shortest_cycle_length = utils.ArrayReduce(cycle_steps, 1, utils.LCM)

			// New shortest cycle length is the lowest common multiple of the current shortest cycle length and the largest number of steps
			// shortest_cycle_length = utils.LCM(shortest_cycle_length, largest)
			j++
		}

		// time_start := time.Now()
		// j := 0
		// for {
		// 	if areAllCyclesAtEndNode(cycles) {
		// 		break
		// 	}
		// 	walkCycles(cycles, shortest_cycle_length)
		// 	j++
		// 	if j%10_000_000 == 0 {
		// 		fmt.Printf("j: %d, time: %v\n", j, time.Since(time_start))
		// 		fmt.Printf("cycles: %v\n", cycles)
		// 	}
		// }

		// fmt.Printf("j: %d\n", j)
		// fmt.Printf("cycles: %v\n", cycles)

		i = -1
	}

	return i, nil
}

// 5681417444136477648 too high

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

func areAllEndNodes(current_nodes []string) bool {
	for _, node := range current_nodes {
		if !isEndNode(node) {
			return false
		}
	}
	return true
}

type NodeAndDirection struct {
	node      string
	direction EnumeratedDirection
}

type Cycle struct {
	warmup_steps   int
	cycle_steps    int
	end_node_steps int
}

func findCycle(nodes map[string]Node, directions []bool, current_node string) Cycle {

	original_node := current_node
	directions_chan := make(chan EnumeratedDirection)
	go directionCycler(directions, directions_chan)
	defer close(directions_chan)

	visited_nodes := make(map[NodeAndDirection]bool)
	end_nodes_visited := 0

	var current_state NodeAndDirection
	var direction EnumeratedDirection

	for {
		direction = <-directions_chan

		current_state = NodeAndDirection{
			node:      current_node,
			direction: direction,
		}

		// Count the number of end nodes we've visited
		if isEndNode(current_node) {
			end_nodes_visited++
		}

		// Check if we've visited this node before
		if visited_nodes[current_state] {
			break
		}
		visited_nodes[current_state] = true

		// Move to the next node
		current_node = nextNode(nodes, current_node, direction)
	}

	// NOTE: This needn't be true, but it is for the input and it simplifies the problem.
	if end_nodes_visited != 1 {
		panic(fmt.Sprintf("expected to visit 1 end node, but visited %d", end_nodes_visited))
	}

	// We're now on the first node and direction of the cycle. We need to find the length of the cycle and the index of the end node.
	warmup_steps := len(visited_nodes) + 1
	// fmt.Println("warmup_steps:", warmup_steps)

	// Keep running through the cycle until we reach the end node again
	cycle_steps := 0
	// Copy of the current state to compare against
	stop_state := current_state
	// fmt.Println("stop_state:", stop_state)
	// fmt.Println("Running through the cycle again to find the cycle length")
	for {
		if cycle_steps > 0 {
			// For cycle_steps == 0 we already have the first direction form the loop above
			direction = <-directions_chan
		}

		current_state = NodeAndDirection{
			node:      current_node,
			direction: direction,
		}

		if cycle_steps == 0 {
			if current_state != stop_state {
				fmt.Println("current_state:", current_state)
				fmt.Println("stop_state:", stop_state)
				panic("internal sanity check failed (1)")
			}
		} else {
			if current_state == stop_state {
				// We've gone all the way around the cycle and back to the start + 1
				break
			}
		}

		// Actually move to the next node
		current_node = nextNode(nodes, current_node, direction)
		cycle_steps++
	}

	// We need to add 1 to the cycle count to account for the final step
	cycle_steps++

	// Check that if we take another cycle_steps steps, we end up at the same node
	// fmt.Println("cycle_steps:", cycle_steps)
	// fmt.Println("checking that we end up at the same node after another cycle_steps steps")
	for i := 0; i < cycle_steps; i++ {
		if i > 0 {
			// For i == 0 we already have the first direction form the loop above
			direction = <-directions_chan
		}

		current_state = NodeAndDirection{
			node:      current_node,
			direction: direction,
		}

		if i == 0 {
			if current_state != stop_state {
				panic("internal sanity check failed (2)")
			}
		}

		if i == cycle_steps-1 {
			if current_state != stop_state {
				panic("failed to find the correct cycle length")
			}
			break
		}

		// Actually move to the next node
		current_node = nextNode(nodes, current_node, direction)
	}

	// Check that if we take another cycle_steps steps, we end up at the same node. This time also get the index of the end node with respect to the cycle start
	end_node_steps := -1
	for i := 0; i < cycle_steps; i++ {
		if i > 0 {
			// For i == 0 we already have the first direction form the loop above
			direction = <-directions_chan
		}

		current_state = NodeAndDirection{
			node:      current_node,
			direction: direction,
		}

		if i == 0 {
			if current_state != stop_state {
				panic("internal sanity check failed (3)")
			}
		}

		if i == cycle_steps-1 {
			if current_state != stop_state {
				panic("failed to find the correct cycle length")
			}
			break
		}

		if isEndNode(current_node) {
			end_node_steps = i - 2
		}

		current_node = nextNode(nodes, current_node, direction)
	}

	if current_state != stop_state {
		panic("internal sanity check failed (4)")
	}

	end_node_steps--

	// // Check that the end_node_steps is correct
	// fmt.Println("end_node_steps:", end_node_steps)
	// for i := 0; i < end_node_steps; i++ {
	// 	if i > 0 {
	// 		// For i == 0 we already have the first direction form the loop above
	// 		direction = <-directions_chan
	// 	}

	// 	current_node = nextNode(nodes, current_node, direction)
	// }

	// if !isEndNode(current_node) {
	// 	panic("failed to find the correct end node index")
	// }

	cycle := Cycle{
		warmup_steps:   warmup_steps,
		cycle_steps:    cycle_steps,
		end_node_steps: end_node_steps,
	}

	err := checkCycle(nodes, directions, original_node, cycle)
	if err != nil {
		panic(err)
	}

	return cycle
}

func checkCycle(nodes_map map[string]Node, directions []bool, current_node string, cycle Cycle) error {
	warmup_steps := cycle.warmup_steps
	cycle_steps := cycle.cycle_steps
	end_node_steps := cycle.end_node_steps

	if end_node_steps > cycle_steps {
		return fmt.Errorf("end_node_steps > cycle_steps")
	}

	// Final check. Starting from the beginning, take warmup_steps steps, then cycle_steps steps, then end_node_steps steps.
	// We should do the warm up, then take a full cycle, then end up at the end node.

	directions_chan := make(chan EnumeratedDirection)
	go directionCycler(directions, directions_chan)
	defer close(directions_chan)

	var current_state NodeAndDirection
	var stop_state NodeAndDirection

	for i := 0; i < warmup_steps; i++ {
		direction := <-directions_chan
		current_state = NodeAndDirection{
			node:      current_node,
			direction: direction,
		}
		current_node = nextNode(nodes_map, current_node, direction)
	}

	// This is the state we should end up at after the warmup
	stop_state = current_state

	// Take a full cycle steps
	for i := 0; i < cycle_steps; i++ {
		current_node = nextNode(nodes_map, current_node, <-directions_chan)
	}

	if current_state != stop_state {
		return fmt.Errorf("the cycle is not correct (1)")
	}

	// Take another full cycle steps
	for i := 0; i < cycle_steps; i++ {
		current_node = nextNode(nodes_map, current_node, <-directions_chan)
	}

	if current_state != stop_state {
		return fmt.Errorf("the cycle is not correct (2)")
	}

	// Take the steps to get to the end node
	for i := 0; i < end_node_steps; i++ {
		current_node = nextNode(nodes_map, current_node, <-directions_chan)
	}

	if !isEndNode(current_node) {
		return fmt.Errorf("the end node is not correct (1)")
	}

	// Take the remaining steps to get to the beginning of the cycle
	for i := 0; i < cycle_steps-end_node_steps; i++ {
		current_node = nextNode(nodes_map, current_node, <-directions_chan)
	}

	if current_state != stop_state {
		return fmt.Errorf("the cycle is not correct (3)")
	}

	return nil
}

func alignCycles(cycles []Cycle) {
	// Find the cycle with the longest warmup_steps
	max_warmup_steps := 0
	for _, cycle := range cycles {
		if cycle.warmup_steps > max_warmup_steps {
			max_warmup_steps = cycle.warmup_steps
		}
	}

	// Align all the cycles to the same starting point
	for i, cycle := range cycles {
		cycles[i].Walk(max_warmup_steps - cycle.warmup_steps)
	}
}

// Walk the cycle forward or backward by delta steps. Wrap around the end of the cycle.
func (cycle *Cycle) Walk(delta int) {
	cycle.warmup_steps += delta
	cycle.end_node_steps -= delta

	// Wrap the end node steps
	for cycle.end_node_steps < 0 {
		cycle.end_node_steps += cycle.cycle_steps
	}

	for cycle.end_node_steps >= cycle.cycle_steps {
		cycle.end_node_steps -= cycle.cycle_steps
	}
}

func walkCycles(cycles []Cycle, delta int) {
	for i := range cycles {
		cycles[i].Walk(delta)
	}
}

// Walk all the cycles forward such that the shortest cycle ends up at the end node.
// Return the length of the shortest cycle.
func walkToEndNodeOfShortestCycle(cycles []Cycle) int {
	// Find the cycle with the shortest cycle_steps
	min_cycle, _, err := utils.MinArrayFunc(cycles, func(c1, c2 Cycle) bool {
		return c1.cycle_steps < c2.cycle_steps
	})

	if err != nil {
		panic(err)
	}

	walkCycles(cycles, min_cycle.end_node_steps)

	shortest_cycle_cycle_steps := min_cycle.cycle_steps

	return shortest_cycle_cycle_steps
}

func areAllCyclesAtEndNode(cycles []Cycle) bool {
	for _, cycle := range cycles {
		if cycle.end_node_steps != 0 {
			return false
		}
	}
	return true
}

func calcNextEndStateAlignment(cycle Cycle, step int) int {
	// The end node is at the start of the cycle so we're aligned now
	if cycle.end_node_steps == 0 {
		return 0
	}

	// The end node is not the start so walk until it is
	cycle_copy := cycle
	i := 0
	for {
		cycle_copy.Walk(step)
		i++
		if cycle_copy.end_node_steps == 0 {
			break
		}
	}
	return i * step
}

func calcNextEndStateAlignmentSafe(cycle Cycle, step int) int {
	// The end node is at the start of the cycle so we're aligned now
	if cycle.end_node_steps == 0 {
		return 0
	}

	// The end node is not the start so walk until it is
	visited_end_node_steps := make(map[int]bool)
	visited_end_node_steps[cycle.end_node_steps] = true
	cycle_copy := cycle
	i := 0
	for {
		cycle_copy.Walk(step)
		i++
		if cycle_copy.end_node_steps == 0 {
			break
		}
		if visited_end_node_steps[cycle_copy.end_node_steps] {
			fmt.Printf("cycle: %v\n", cycle_copy)
			fmt.Printf("step: %d\n", step)
			fmt.Printf("len(visited_end_node_steps): %d\n", len(visited_end_node_steps))
			panic("failed to find the next alignment. The cycle is not correct")
		}
	}
	return i * step
}
