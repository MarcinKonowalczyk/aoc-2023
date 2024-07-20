package day16

import (
	"aoc2023/utils"
	"fmt"
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
	grid, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	// Add the starting beam
	grid.beam_ends = append(grid.beam_ends, beam_end{0, 0, RIGHT})

	for {
		carry_on, err := grid.stepBeams()
		if err != nil {
			return -1, err
		}
		if !carry_on {
			break
		}
	}

	energized := countEnergized(grid)
	return energized, nil
}

type tile rune

const (
	// tiles for the grid
	EMPTY            tile = '.'
	VERTIAL          tile = '|'
	HORIZONTAL       tile = '-'
	MIRROR_SLASH     tile = '/'
	MIRROR_BACKSLASH tile = '\\'

	// single directions
	UP    tile = '^'
	DOWN  tile = 'v'
	LEFT  tile = '<'
	RIGHT tile = '>'
	// L shaped beams
	UP_RIGHT   tile = 'L'
	DOWN_RIGHT tile = 'F'
	DOWN_LEFT  tile = '7'
	UP_LEFT    tile = 'J'
	// Line beams
	UP_DOWN    tile = ':'
	LEFT_RIGHT tile = '='
	// T shaped beams
	UP_RIGHT_DOWN   tile = 'E'
	RIGHT_DOWN_LEFT tile = 'T'
	DOWN_LEFT_UP    tile = '3'
	LEFT_UP_RIGHT   tile = 'W'
	// Cross
	UP_RIGHT_DOWN_LEFT tile = '+'
)

func (t tile) Name() string {
	switch t {
	// tiles for the grid
	case EMPTY:
		return "EMPTY"
	case VERTIAL:
		return "VERTIAL"
	case HORIZONTAL:
		return "HORIZONTAL"
	case MIRROR_SLASH:
		return "MIRROR_SLASH"
	case MIRROR_BACKSLASH:
		return "MIRROR_BACKSLASH"
	// single directions
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	// L shaped beams
	case UP_RIGHT:
		return "UP_RIGHT"
	case DOWN_RIGHT:
		return "DOWN_RIGHT"
	case DOWN_LEFT:
		return "DOWN_LEFT"
	case UP_LEFT:
		return "UP_LEFT"
	// Line beams
	case UP_DOWN:
		return "UP_DOWN"
	case LEFT_RIGHT:
		return "LEFT_RIGHT"
	// T shaped beams
	case UP_RIGHT_DOWN:
		return "UP_RIGHT_DOWN"
	case RIGHT_DOWN_LEFT:
		return "RIGHT_DOWN_LEFT"
	case DOWN_LEFT_UP:
		return "DOWN_LEFT_UP"
	case LEFT_UP_RIGHT:
		return "LEFT_UP_RIGHT"
	// Cross
	case UP_RIGHT_DOWN_LEFT:
		return "UP_RIGHT_DOWN_LEFT"
	default:
		return "UNKNOWN"
	}
}

type beam_end struct {
	x int
	y int
	t tile
}

type grid struct {
	tiles     [][]tile
	trails    [][]tile
	rows      int
	cols      int
	beam_ends []beam_end
}

func (g grid) String() string {
	s := ""
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			t := g.tiles[i][j]
			r := g.trails[i][j]
			if t == EMPTY {
				s += string(r)
			} else {
				s += string(t)
			}
		}
		if i < g.rows-1 {
			s += "\n"
		}
	}
	return s
}

func (g grid) DebugString() string {
	s := ""
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			t := g.tiles[i][j]
			r := g.trails[i][j]
			s += string(t)
			s += string(r)
		}
		if i < g.rows-1 {
			s += "\n"
		}
	}
	return s
}

func parseLines(lines []string) (grid, error) {
	n_rows := len(lines)
	if n_rows == 0 {
		return grid{}, fmt.Errorf("empty grid")
	}
	n_cols := len(lines[0])
	tiles := make([][]tile, n_rows)
	for i, line := range lines {
		tiles[i] = make([]tile, 0)
		for _, r := range line {
			tiles[i] = append(tiles[i], tile(r))
		}
		if len(tiles[i]) != n_cols {
			return grid{}, fmt.Errorf("inconsistent number of columns")
		}
	}
	trails := make([][]tile, n_rows)
	for i := 0; i < n_rows; i++ {
		trails[i] = make([]tile, n_cols)
		for j := 0; j < n_cols; j++ {
			trails[i][j] = EMPTY
		}
	}
	return grid{tiles, trails, n_rows, n_cols, []beam_end{}}, nil
}

func (g *grid) stepBeams() (bool, error) {
	if len(g.beam_ends) == 0 {
		return false, nil
	}
	nb := make([]beam_end, 0)
	for _, b := range g.beam_ends {
		t := g.tiles[b.y][b.x]
		// fmt.Printf("%d,%d %s::%s\n", b.y, b.x, t.Name(), b.t.Name())
		switch t {
		case EMPTY:
			switch b.t {
			case UP:
				if b.y > 0 {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			case DOWN:
				if b.y < g.rows-1 {
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			case LEFT:
				if b.x > 0 {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
			case RIGHT:
				if b.x < g.cols-1 {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			}
		case VERTIAL:
			switch b.t {
			case UP:
				if b.y > 0 {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			case DOWN:
				if b.y < g.rows-1 {
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			case LEFT, RIGHT:
				if b.y > 0 {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
				if b.y < g.rows-1 {
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			}
		case HORIZONTAL:
			switch b.t {
			case UP, DOWN:
				if b.x > 0 {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
				if b.x < g.cols-1 {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			case LEFT:
				if b.x > 0 {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
			case RIGHT:
				if b.x < g.cols-1 {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			}
		case MIRROR_SLASH:
			switch b.t {
			case UP:
				if b.x < g.cols-1 {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			case DOWN:
				if b.x > 0 {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
			case LEFT:
				if b.y < g.rows-1 {
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			case RIGHT:
				if b.y > 0 {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			}
		case MIRROR_BACKSLASH:
			switch b.t {
			case UP:
				if b.x > 0 {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
			case DOWN:
				if b.x < g.cols-1 {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			case LEFT:
				if b.y > 0 {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			case RIGHT:
				if b.y < g.rows-1 {
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			}
		default:
			panic(fmt.Sprintf("tile %s should not be in this grid", t.Name()))
		}
	}

	// Prune new beam ends
	to_prune := []int{}
	for i := 0; i < len(nb); i++ {
		b := nb[i]
		r := g.trails[b.y][b.x]
		switch b.t {
		case UP:
			switch r {
			case UP, UP_LEFT, UP_RIGHT, UP_DOWN, UP_RIGHT_DOWN, LEFT_UP_RIGHT, DOWN_LEFT_UP, UP_RIGHT_DOWN_LEFT:
				to_prune = append(to_prune, i)
			}
		case DOWN:
			switch r {
			case DOWN, DOWN_LEFT, DOWN_RIGHT, UP_DOWN, UP_RIGHT_DOWN, RIGHT_DOWN_LEFT, UP_RIGHT_DOWN_LEFT:
				to_prune = append(to_prune, i)
			}
		case LEFT:
			switch r {
			case LEFT, UP_LEFT, DOWN_LEFT, LEFT_RIGHT, RIGHT_DOWN_LEFT, DOWN_LEFT_UP, LEFT_UP_RIGHT, UP_RIGHT_DOWN_LEFT:
				to_prune = append(to_prune, i)
			}
		case RIGHT:
			switch r {
			case RIGHT, UP_RIGHT, DOWN_RIGHT, LEFT_RIGHT, LEFT_UP_RIGHT, UP_RIGHT_DOWN, RIGHT_DOWN_LEFT, UP_RIGHT_DOWN_LEFT:
				to_prune = append(to_prune, i)
			}
		}
	}
	// TODO: Actually prune new beam ends
	var n_removed int
	if len(to_prune) > 0 {
		// nb_subset := make([]beam_end, 0)
		// for _, I := range to_prune {
		// 	nb_subset = append(nb_subset, nb[I])
		// }

		// fmt.Println("Would prune", to_prune, "aka", nb_subset)
		// panic("pruning not implemented")
		nb, n_removed = utils.ArrayRemoveIndices(nb, to_prune)
		if n_removed != len(to_prune) {
			panic("unexpected number of elements removed")
		}
	}

	// Write the old beam ends to the trails
	for _, b := range g.beam_ends {
		r := g.trails[b.y][b.x]
		switch r {
		case EMPTY:
			g.trails[b.y][b.x] = b.t
		case UP:
			switch b.t {
			case UP:
			case DOWN:
				g.trails[b.y][b.x] = UP_DOWN
			case LEFT:
				g.trails[b.y][b.x] = UP_LEFT
			case RIGHT:
				g.trails[b.y][b.x] = UP_RIGHT
			}
		case DOWN:
			switch b.t {
			case DOWN:
			case UP:
				g.trails[b.y][b.x] = UP_DOWN
			case LEFT:
				g.trails[b.y][b.x] = DOWN_LEFT
			case RIGHT:
				g.trails[b.y][b.x] = DOWN_RIGHT
			}
		case LEFT:
			switch b.t {
			case LEFT:
			case UP:
				g.trails[b.y][b.x] = UP_LEFT
			case DOWN:
				g.trails[b.y][b.x] = DOWN_LEFT
			case RIGHT:
				g.trails[b.y][b.x] = LEFT_RIGHT
			}
		case RIGHT:
			switch b.t {
			case RIGHT:
			case UP:
				g.trails[b.y][b.x] = UP_RIGHT
			case DOWN:
				g.trails[b.y][b.x] = DOWN_RIGHT
			case LEFT:
				g.trails[b.y][b.x] = LEFT_RIGHT
			}
		// L shaped beams
		case UP_RIGHT:
			switch b.t {
			case UP, RIGHT:
			case DOWN:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN
			case LEFT:
				g.trails[b.y][b.x] = LEFT_UP_RIGHT
			}
		case DOWN_RIGHT:
			switch b.t {
			case DOWN, RIGHT:
			case UP:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN
			case LEFT:
				g.trails[b.y][b.x] = RIGHT_DOWN_LEFT
			}
		case DOWN_LEFT:
			switch b.t {
			case DOWN, LEFT:
			case UP:
				g.trails[b.y][b.x] = DOWN_LEFT_UP
			case RIGHT:
				g.trails[b.y][b.x] = RIGHT_DOWN_LEFT
			}
		case UP_LEFT:
			switch b.t {
			case UP, LEFT:
			case DOWN:
				g.trails[b.y][b.x] = DOWN_LEFT_UP
			case RIGHT:
				g.trails[b.y][b.x] = LEFT_UP_RIGHT
			}
		// Line beams
		case UP_DOWN:
			switch b.t {
			case UP, DOWN:
			case LEFT:
				g.trails[b.y][b.x] = DOWN_LEFT_UP
			case RIGHT:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN
			}
		case LEFT_RIGHT:
			switch b.t {
			case LEFT, RIGHT:
			case UP:
				g.trails[b.y][b.x] = LEFT_UP_RIGHT
			case DOWN:
				g.trails[b.y][b.x] = RIGHT_DOWN_LEFT
			}
		// T shaped beams
		case UP_RIGHT_DOWN:
			switch b.t {
			case UP, RIGHT, DOWN:
			case LEFT:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			}
		case RIGHT_DOWN_LEFT:
			switch b.t {
			case RIGHT, DOWN, LEFT:
			case UP:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			}
		case DOWN_LEFT_UP:
			switch b.t {
			case DOWN, LEFT, UP:
			case RIGHT:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			}
		case LEFT_UP_RIGHT:
			switch b.t {
			case LEFT, UP, RIGHT:
			case DOWN:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			}
		// Cross
		case UP_RIGHT_DOWN_LEFT:
			// Do nothing. All the beams directions are already set
		default:
			panic(fmt.Sprintf("tile %s should not be in the trail map", r.Name()))
		}
	}

	// Set the beam ends
	g.beam_ends = nb

	if len(g.beam_ends) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func countEnergized(g grid) int {
	count := 0
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.trails[i][j] != EMPTY {
				count++
			}
		}
	}
	return count
}
