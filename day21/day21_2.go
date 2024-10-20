package day21

import (
	"aoc2023/utils"
	"fmt"
)

// const N2 = 26501365
// const N2 = 49 // to the edge of an empty 9x9
const N2 = 1

func main_2(lines []string, verbose bool) (n int, err error) {
	g, err := parseLines(lines)
	if err != nil {
		return 0, err
	}

	if g.rows != g.cols {
		return 0, fmt.Errorf("garden is not square")
	}

	if g.rows%2 != 1 {
		return 0, fmt.Errorf("garden size is not odd")
	}

	g.Clear()

	gs := g.rows
	hgs := (gs - 1) / 2
	fmt.Println("Half garden size:", hgs)

	g.Tile(3)
	// Dumb method. Just tile a lot and then step that many times

	for i := 0; i < N2; i++ {
		g.Step()
	}

	if verbose {
		fmt.Println("After", N2, "steps:")
		fmt.Println(g)
	}

	fmt.Println("After", hgs+gs, "steps:")
	n_full := (N2 - gs - hgs) / gs
	fmt.Println(N2)
	fmt.Println(n_full * gs)
	fmt.Println(N2 - n_full*gs)

	return -1, nil
}

// Tile garden into a larger nxn garden
func (g *Garden) Tile(n int) {

	if g.tiles != 0 {
		panic("already tiled")
	}

	new_grid := make([][]bool, g.rows*n)
	for i := 0; i < g.rows*n; i++ {
		new_grid[i] = make([]bool, g.cols*n)
	}
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					new_grid[i+k*g.rows][j+l*g.cols] = g.grid[i][j]
				}
			}
		}
	}

	delta_x := g.rows * ((n - 1) / 2)
	delta_y := g.cols * ((n - 1) / 2)
	new_positions := make([]utils.Point2, 0)
	for _, p := range g.positions {
		new_positions = append(new_positions, utils.Point2{X: p.X + delta_x, Y: p.Y + delta_y})
	}

	g.grid = new_grid
	g.rows *= n
	g.cols *= n
	g.positions = new_positions
	g.tiles = n
}

func (g *Garden) Clear() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			g.grid[i][j] = false
		}
	}
}
