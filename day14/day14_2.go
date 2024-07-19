package day14

import (
	"aoc2023/utils"
	"fmt"
)

const N = 1000000000

// const N = 100

func main_2(lines []string) (n int, err error) {

	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	log.Fatal("could not create CPU profile: ", err)
	// }
	// defer f.Close() // error handling omitted for example
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	log.Fatal("could not start CPU profile: ", err)
	// }
	// defer pprof.StopCPUProfile()

	grid := parseLines(lines)
	// spinCycle(&grid)
	// fmt.Println("After 1 spin cycle:")
	// fmt.Println(grid)
	// spinCycle(&grid)
	// fmt.Println("After 2 spin cycles:")
	// fmt.Println(grid)
	// spinCycle(&grid)
	// fmt.Println("After 3 spin cycles:")
	// fmt.Println(grid)

	// Do a bunch of cycles and record weigbt after each one, to find cycles
	weights := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		spinCycle(grid)
		weights[i] = calcNorthWeight(grid)
	}

	cycle_start, cycle_period := utils.DetectCycles(weights)
	if cycle_start == -1 {
		return -1, fmt.Errorf("no cycle found")
	}

	// fmt.Printf("Cycle detected: start=%d, period=%d\n", cycle_start, cycle_period)
	extrapolated_weight := utils.ExtrapolateCycle(weights, N-1, cycle_start, cycle_period)
	return extrapolated_weight, nil
}

func spinCycle(g Grid) {
	// slideNorth(g)      // actual north
	// rotateClockwise(g) // north is now east, so slideNorth is slideWest
	// slideNorth(g)      // west
	// rotateClockwise(g) // ...
	// slideNorth(g)      // south
	// rotateClockwise(g)
	// slideNorth(g)      // east
	// rotateClockwise(g) // back to the original orientation
	slideNorth(g) // north
	slideWest(g)  // west
	slideSouth(g) // south
	slideEast(g)  // east
}

func slideWest(g Grid) {
	for row := 0; row < g.rows; row++ {
		for col := 1; col < g.cols; col++ {
			if (*g.rocks)[row][col] != ROUND {
				continue
			}
			col2 := col
			rock_left := (*g.rocks)[row][col2-1]
			for rock_left == EMPTY {
				col2--
				if col2 == 0 {
					break
				}
				rock_left = (*g.rocks)[row][col2-1]
			}
			(*g.rocks)[row][col], (*g.rocks)[row][col2] = EMPTY, ROUND
		}
	}
}

func slideSouth(g Grid) {
	for row := g.rows - 2; row >= 0; row-- {
		for col := 0; col < g.cols; col++ {
			if (*g.rocks)[row][col] != ROUND {
				continue
			}
			row2 := row
			rock_below := (*g.rocks)[row2+1][col]
			for rock_below == EMPTY {
				row2++
				if row2 == g.rows-1 {
					break
				}
				rock_below = (*g.rocks)[row2+1][col]
			}
			(*g.rocks)[row][col], (*g.rocks)[row2][col] = EMPTY, ROUND
		}
	}
}

func slideEast(g Grid) {
	for row := 0; row < g.rows; row++ {
		for col := g.cols - 2; col >= 0; col-- {
			if (*g.rocks)[row][col] != ROUND {
				continue
			}
			col2 := col
			rock_right := (*g.rocks)[row][col2+1]
			for rock_right == EMPTY {
				col2++
				if col2 == g.cols-1 {
					break
				}
				rock_right = (*g.rocks)[row][col2+1]
			}
			(*g.rocks)[row][col], (*g.rocks)[row][col2] = EMPTY, ROUND
		}
	}
}
