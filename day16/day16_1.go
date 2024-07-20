package day16

import (
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
	fmt.Println(grid)
	for {
		carry_on, err := grid.stepBeams()
		if err != nil {
			return -1, err
		}
		fmt.Println(grid)
		if !carry_on {
			break
		}
	}
	return 0, nil
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

func (t tile) isBeam() bool {
	return t == UP || t == DOWN || t == LEFT || t == RIGHT || t == UP_RIGHT || t == DOWN_RIGHT || t == DOWN_LEFT || t == UP_LEFT || t == UP_RIGHT_DOWN || t == RIGHT_DOWN_LEFT || t == DOWN_LEFT_UP || t == LEFT_UP_RIGHT || t == UP_RIGHT_DOWN_LEFT
}

// func (t tile) isMirror() bool {
// 	return t == MIRROR_SLASH || t == MIRROR_BACKSLASH
// }

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

// func (g grid) String() string {
// 	s := "GRID:\n"
// 	for i := 0; i < g.rows; i++ {
// 		s += " "
// 		for j := 0; j < g.cols; j++ {
// 			s += string(g.tiles[i][j])
// 		}
// 		if i < g.rows-1 {
// 			s += "\n"
// 		}
// 	}
// 	s += "\nTRAILS:\n"
// 	for i := 0; i < g.rows; i++ {
// 		s += " "
// 		for j := 0; j < g.cols; j++ {
// 			s += string(g.trails[i][j])
// 		}
// 		if i < g.rows-1 {
// 			s += "\n"
// 		}
// 	}
// 	s += "\n"
// 	return s
// }

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
	return grid{tiles, trails, n_rows, n_cols, []beam_end{{0, 0, RIGHT}}}, nil
}

func (g *grid) stepBeams() (bool, error) {
	if len(g.beam_ends) == 0 {
		return false, nil
	}
	nb := make([]beam_end, 0)
	for _, b := range g.beam_ends {
		t := g.tiles[b.y][b.x]
		fmt.Printf("%d,%d %s::%s\n", b.y, b.x, t.Name(), b.t.Name())
		switch t {
		case EMPTY:
			switch b.t {
			case UP:
				if b.y == 0 {
					// We cant go up anymore because we are at the top edge of the grid
				} else {
					// We keep going up
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			case DOWN:
				if b.y == g.rows-1 {
					// We cant go down anymore because we are at the bottom edge of the grid
				} else {
					// We keep going down
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			case LEFT:
				if b.x == 0 {
					// We cant go left anymore because we are at the left edge of the grid
				} else {
					// We keep going left
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
			case RIGHT:
				if b.x == g.cols-1 {
					// We cant go right anymore because we are at the right edge of the grid
				} else {
					// We keep going right
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			default:
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
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
			default:
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case HORIZONTAL:
			switch b.t {
			case UP:
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			case DOWN:
				if b.x == 0 {
					// We can't go left from here
				} else {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
				if b.x == g.cols-1 {
					// We can't go right from here
				} else {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			case LEFT:
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			case RIGHT:
				// Carry on as if this is empty space
				if b.x == g.cols-1 {
					// We cant go right anymore because we are at the right edge of the grid
				} else {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			default:
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case MIRROR_SLASH:
			if b.t == UP {
				if b.x == g.cols-1 {
					// We cant go right anymore because we are at the right edge of the grid
				} else {
					nb = append(nb, beam_end{b.x + 1, b.y, RIGHT})
				}
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				if b.y == 0 {
					// We can't go up from here
				} else {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case MIRROR_BACKSLASH:
			if b.t == UP {
				if b.x == 0 {
					// We cant go left anymore because we are at the left edge of the grid
				} else {
					nb = append(nb, beam_end{b.x - 1, b.y, LEFT})
				}
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				if b.y == 0 {
					// We can't go up from here
				} else {
					nb = append(nb, beam_end{b.x, b.y - 1, UP})
				}
			} else if b.t == RIGHT {
				if b.y == g.rows-1 {
					// We can't go down from here
				} else {
					nb = append(nb, beam_end{b.x, b.y + 1, DOWN})
				}
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case UP:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case DOWN:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case LEFT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case RIGHT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case UP_RIGHT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case DOWN_RIGHT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case DOWN_LEFT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case UP_LEFT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case UP_DOWN:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case LEFT_RIGHT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case UP_RIGHT_DOWN:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case RIGHT_DOWN_LEFT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case DOWN_LEFT_UP:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case LEFT_UP_RIGHT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		case UP_RIGHT_DOWN_LEFT:
			if b.t == UP {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == DOWN {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == LEFT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else if b.t == RIGHT {
				panic(fmt.Sprintf("%s::%s not implemented", t.Name(), b.t.Name()))
			} else {
				panic(fmt.Sprintf("%s: unknown direction", t.Name()))
			}
		default:
			panic("invalid tile")
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
		default:
			panic("invalid beam end")
		}
	}
	// TODO: Actually prune new beam ends
	if len(to_prune) > 0 {
		fmt.Println("Would prune", to_prune)
	}

	// Write the old beam ends to the trails
	for _, b := range g.beam_ends {
		r := g.trails[b.y][b.x]
		switch r {
		case EMPTY:
			g.trails[b.y][b.x] = b.t
		case VERTIAL, HORIZONTAL, MIRROR_SLASH, MIRROR_BACKSLASH:
			panic(fmt.Sprintf("tile %s shoudl not appear in the trail map", r.Name()))
		case UP:
			switch b.t {
			case UP:
			case DOWN:
				g.trails[b.y][b.x] = UP_DOWN
			case LEFT:
				g.trails[b.y][b.x] = UP_LEFT
			case RIGHT:
				g.trails[b.y][b.x] = UP_RIGHT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
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
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
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
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
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
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		// L shaped beams
		case UP_RIGHT:
			switch b.t {
			case UP, RIGHT:
			case DOWN:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN
			case LEFT:
				g.trails[b.y][b.x] = LEFT_UP_RIGHT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case DOWN_RIGHT:
			switch b.t {
			case DOWN, RIGHT:
			case UP:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN
			case LEFT:
				g.trails[b.y][b.x] = RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case DOWN_LEFT:
			switch b.t {
			case DOWN, LEFT:
			case UP:
				g.trails[b.y][b.x] = DOWN_LEFT_UP
			case RIGHT:
				g.trails[b.y][b.x] = RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case UP_LEFT:
			switch b.t {
			case UP, LEFT:
			case DOWN:
				g.trails[b.y][b.x] = DOWN_LEFT_UP
			case RIGHT:
				g.trails[b.y][b.x] = LEFT_UP_RIGHT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		// Line beams
		case UP_DOWN:
			switch b.t {
			case UP, DOWN:
			case LEFT:
				g.trails[b.y][b.x] = DOWN_LEFT_UP
			case RIGHT:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case LEFT_RIGHT:
			switch b.t {
			case LEFT, RIGHT:
			case UP:
				g.trails[b.y][b.x] = LEFT_UP_RIGHT
			case DOWN:
				g.trails[b.y][b.x] = RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		// T shaped beams
		case UP_RIGHT_DOWN:
			switch b.t {
			case UP, RIGHT, DOWN:
			case LEFT:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case RIGHT_DOWN_LEFT:
			switch b.t {
			case RIGHT, DOWN, LEFT:
			case UP:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case DOWN_LEFT_UP:
			switch b.t {
			case DOWN, LEFT, UP:
			case RIGHT:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		case LEFT_UP_RIGHT:
			switch b.t {
			case LEFT, UP, RIGHT:
			case DOWN:
				g.trails[b.y][b.x] = UP_RIGHT_DOWN_LEFT
			default:
				panic(fmt.Sprintf("%s: unknown direction", r.Name()))
			}
		// Cross
		case UP_RIGHT_DOWN_LEFT:
			// Do nothing. All the beams directions are already set
		default:
			panic("invalid tile")
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
