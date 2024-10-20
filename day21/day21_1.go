package day21

import (
	"aoc2023/utils"
	"fmt"
)

func Main(part int, lines []string, verbose bool) (n int, err error) {
	if part == 1 {
		return main_1(lines, verbose)
	} else if part == 2 {
		return main_2(lines, verbose)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

const N1 = 64

func main_1(lines []string, verbose bool) (n int, err error) {
	g, err := parseLines(lines)
	if err != nil {
		return 0, err
	}

	if verbose {
		fmt.Println("Initial state:")
		fmt.Println(g)
	}

	for i := 0; i < N1; i++ {
		g.Step()
	}

	if verbose {
		fmt.Println("After", N1, "steps:")
		fmt.Println(g)
	}

	return len(g.positions), nil
}

func (g *Garden) Step() {
	new_positions := make([]utils.Point2, 0)
	for _, p := range g.positions {
		if p.X > 0 {
			if !g.grid[p.X-1][p.Y] {
				new_positions = append(new_positions, utils.Point2{X: p.X - 1, Y: p.Y})
			}
		}
		if p.X < g.rows-1 {
			if !g.grid[p.X+1][p.Y] {
				new_positions = append(new_positions, utils.Point2{X: p.X + 1, Y: p.Y})
			}
		}
		if p.Y > 0 {
			if !g.grid[p.X][p.Y-1] {
				new_positions = append(new_positions, utils.Point2{X: p.X, Y: p.Y - 1})
			}
		}
		if p.Y < g.cols-1 {
			if !g.grid[p.X][p.Y+1] {
				new_positions = append(new_positions, utils.Point2{X: p.X, Y: p.Y + 1})
			}
		}

	}

	// Deduplicate
	new_positions = utils.ArrayUnique(new_positions)
	// fmt.Println(new_positions)

	g.positions = new_positions
}
