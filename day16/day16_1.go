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

	// tiles describing the direction of the beam
	UP                 tile = '^'
	DOWN               tile = 'v'
	LEFT               tile = '<'
	RIGHT              tile = '>'
	UP_RIGHT           tile = '1'
	DOWN_RIGHT         tile = '2'
	DOWN_LEFT          tile = '3'
	UP_LEFT            tile = '4'
	UP_RIGHT_DOWN      tile = '5'
	RIGHT_DOWN_LEFT    tile = '6'
	DOWN_LEFT_UP       tile = '7'
	LEFT_UP_RIGHT      tile = '8'
	UP_RIGHT_DOWN_LEFT tile = '9'
)

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
	rows      int
	cols      int
	beam_ends []beam_end
}

func (g grid) String() string {
	s := ""
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			s += string(g.tiles[i][j])
		}
		s += "\n"
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
		tiles[i] = make([]tile, n_cols)
		for j, r := range line {
			tiles[i][j] = tile(r)
		}
	}
	if tiles[0][0] != EMPTY {
		return grid{}, fmt.Errorf("0, 0 is not empty")
	}
	tiles[0][0] = RIGHT
	return grid{tiles, n_rows, n_cols, []beam_end{{0, 0, RIGHT}}}, nil
}

func (g *grid) stepBeams() (bool, error) {
	if len(g.beam_ends) == 0 {
		return false, nil
	}
	for _, beam_end := range g.beam_ends {
		t := g.tiles[beam_end.y][beam_end.x]
		if !t.isBeam() {
			panic("beam end is not a beam")
		}
		current_tile := g.tiles[beam_end.y][beam_end.x]
		fmt.Println("current_tile", current_tile)
		switch current_tile {
		case EMPTY:
			panic("EMPTY not implemented")
		case VERTIAL:
			panic("VERTICAL not implemented")
		case HORIZONTAL:
			panic("HORIZONTAL not implemented")
		case MIRROR_SLASH:
			panic("MIRROR_SLASH not implemented")
		case MIRROR_BACKSLASH:
			panic("MIRROR_BACKSLASH not implemented")
		case UP:
			panic("UP not implemented")
		case DOWN:
			panic("DOWN not implemented")
		case LEFT:
			panic("LEFT not implemented")
		case RIGHT:
			panic("RIGHT not implemented")
		case UP_RIGHT:
			panic("UP_RIGHT not implemented")
		case DOWN_RIGHT:
			panic("DOWN_RIGHT not implemented")
		case DOWN_LEFT:
			panic("DOWN_LEFT not implemented")
		case UP_LEFT:
			panic("UP_LEFT not implemented")
		case UP_RIGHT_DOWN:
			panic("UP_RIGHT_DOWN not implemented")
		case RIGHT_DOWN_LEFT:
			panic("RIGHT_DOWN_LEFT not implemented")
		case DOWN_LEFT_UP:
			panic("DOWN_LEFT_UP not implemented")
		case LEFT_UP_RIGHT:
			panic("LEFT_UP_RIGHT not implemented")
		case UP_RIGHT_DOWN_LEFT:
			panic("UP_RIGHT_DOWN_LEFT not implemented")
		default:
			panic("invalid tile")
		}

		// if beam_end.t == UP {
		// 	stepBeamUp(g, beam_end)
		// } else if beam_end.t == DOWN {
		// 	//
		// } else if beam_end.t == LEFT {
		// 	//
		// } else if beam_end.t == RIGHT {
		// 	//
		// } else {
		// 	panic("invalid beam end")
		// }
	}

	if len(g.beam_ends) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
