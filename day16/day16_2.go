package day16

func main_2(lines []string) (n int, err error) {
	grid, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	max_energized := 0

	// Try all possible starting beam ends on left and right sides
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < 2; j++ {
			grid.Reset()
			var be beam_end
			if j == 0 {
				be = beam_end{0, i, RIGHT}
			} else {
				be = beam_end{grid.cols - 1, i, LEFT}
			}
			grid.beam_ends = append(grid.beam_ends, be)
			err = grid.stepBeamsAll()
			if err != nil {
				return -1, err
			}
			energized := countEnergized(grid)
			if energized > max_energized {
				max_energized = energized
				// fmt.Println("New max energized:", max_energized)
				// fmt.Println(grid)
			}
		}
	}

	// Try all possible starting beam ends on top and bottom sides
	for i := 0; i < grid.cols; i++ {
		for j := 0; j < 2; j++ {
			grid.Reset()
			var be beam_end
			if j == 0 {
				be = beam_end{i, 0, DOWN}
			} else {
				be = beam_end{i, grid.rows - 1, UP}
			}
			grid.beam_ends = append(grid.beam_ends, be)
			err = grid.stepBeamsAll()
			if err != nil {
				return -1, err
			}
			energized := countEnergized(grid)
			if energized > max_energized {
				max_energized = energized
				// fmt.Println("New max energized:", max_energized)
				// fmt.Println(grid)
			}
		}
	}

	return max_energized, nil
}

// Drop all the beams and clear the trail map
func (g *grid) Reset() {
	g.beam_ends = []beam_end{}
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			g.trails[i][j] = EMPTY
		}
	}
}
