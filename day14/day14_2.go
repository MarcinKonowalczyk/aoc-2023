package day14

import (
	"aoc2023/cycledetector"
	"fmt"
)

const N = 1000000000

func main_2(lines []string, verbose bool) (n int, err error) {
	grid := parseLines(lines)

	// Do a bunch of cycles and record weigbt after each one, to find cycles
	weights := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		spinCycle(grid)
		weights[i] = calcNorthWeight(grid)
	}

	cycle_start, cycle_period := cycledetector.DetectCycle(weights)
	if cycle_start == -1 {
		return -1, fmt.Errorf("no cycle found")
	}

	extrapolated_weight := cycledetector.ExtrapolateCycle(weights, N-1, cycle_start, cycle_period)
	return extrapolated_weight, nil
}

func rotateClockwise(g Grid) {
	var temp Rock // for swapping
	for i := 0; i < g.rows/2; i++ {
		for j := 0; j < g.cols/2; j++ {
			ii := g.rows - i - 1 // i from the bottom
			jj := g.cols - j - 1 // j from the right
			temp = (*g.rocks)[i][j]
			(*g.rocks)[i][j] = (*g.rocks)[jj][i]
			(*g.rocks)[jj][i] = (*g.rocks)[ii][jj]
			(*g.rocks)[ii][jj] = (*g.rocks)[j][ii]
			(*g.rocks)[j][ii] = temp
		}
	}
}

func spinCycle(g Grid) {
	slideNorth(g)      // actual north
	rotateClockwise(g) // north is now east, so slideNorth is slideWest
	slideNorth(g)      // west
	rotateClockwise(g) // ...
	slideNorth(g)      // south
	rotateClockwise(g)
	slideNorth(g)      // east
	rotateClockwise(g) // back to the original orientation
}
