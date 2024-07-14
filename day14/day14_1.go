package day14

import "fmt"

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
	grid := parseLines(lines)
	slideNorth(&grid)
	weight := calcNorthWeight(&grid)
	return weight, nil
}

type Rock int

const (
	EMPTY Rock = iota
	ROUND
	SQUARE
)

func (r Rock) String() string {
	switch r {
	case EMPTY:
		return "."
	case ROUND:
		return "O"
	case SQUARE:
		return "#"
	}
	return "?"
}

type Grid [][]Rock

func (g Grid) String() string {
	s := ""
	for i, row := range g {
		for _, rock := range row {
			s += rock.String()
		}
		if i < len(g)-1 {
			s += "\n"
		}
	}
	return s
}
func parseLines(lines []string) Grid {
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]Rock, len(line))
		for j, c := range line {
			switch c {
			case '.':
				grid[i][j] = EMPTY
			case 'O':
				grid[i][j] = ROUND
			case '#':
				grid[i][j] = SQUARE
			default:
				panic("invalid character")
			}
		}
	}
	return grid
}

// Slide all the round rocks in the grid to the north
func slideNorth(g *Grid) {
	n_rows := len(*g)
	for i := 0; i < n_rows; i++ {
		// Slide the ith row to the north as much as possible
		slideRowNorth(g, i)
	}
}

func slideRowNorth(g *Grid, row int) {
	n_cols := len((*g)[0])
	for col := 0; col < n_cols; col++ {
		// Slide the ith column to the north as much as possible
		slideRowColNorth(g, row, col)
	}
}

func slideRowColNorth(g *Grid, row, col int) {
	if (*g)[row][col] != ROUND {
		// The rock is not round, so it cannot be slid
		return
	}
	if row == 0 {
		// The rock is already at the top of the grid so cannot be slid
		return
	}
	rock_above := (*g)[row-1][col]
	if rock_above == EMPTY {
		// The rock can be slid to the north
		(*g)[row-1][col] = ROUND
		(*g)[row][col] = EMPTY
		// Call the function recursively to keep sliding the rock to the north
		slideRowColNorth(g, row-1, col)
	}
}

// ======================================================================

func calcNorthWeight(g *Grid) int {
	weight := 0
	n_rows := len(*g)
	for i := 0; i < n_rows; i++ {
		for j := 0; j < len((*g)[0]); j++ {
			if (*g)[i][j] == ROUND {
				weight += (n_rows - i)
			}
		}
	}
	return weight
}
