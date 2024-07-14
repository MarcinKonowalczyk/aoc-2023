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
	slideNorth(grid)
	weight := calcNorthWeight(grid)
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

type Grid struct {
	rocks *[][]Rock
	rows  int
	cols  int
}

func (g Grid) String() string {
	s := ""
	for i, row := range *g.rocks {
		for _, rock := range row {
			s += rock.String()
		}
		if i < len(*g.rocks)-1 {
			s += "\n"
		}
	}
	return s
}

func parseLines(lines []string) Grid {
	rocks := make([][]Rock, len(lines))
	n_rows := len(lines)
	n_cols := len(lines[0])
	for i, line := range lines {
		rocks[i] = make([]Rock, len(line))
		for j, c := range line {
			switch c {
			case '.':
				rocks[i][j] = EMPTY
			case 'O':
				rocks[i][j] = ROUND
			case '#':
				rocks[i][j] = SQUARE
			default:
				panic("invalid character")
			}
		}
	}
	return Grid{&rocks, n_rows, n_cols}
}

// Slide all the round rocks in the grid to the north
func slideNorth(g Grid) {
	for row := 1; row < g.rows; row++ {
		for col := 0; col < g.cols; col++ {
			if (*g.rocks)[row][col] != ROUND {
				continue
			}
			row2 := row
			rock_above := (*g.rocks)[row2-1][col]
			for rock_above == EMPTY {
				row2--
				if row2 == 0 {
					break
				}
				rock_above = (*g.rocks)[row2-1][col]
			}
			(*g.rocks)[row][col], (*g.rocks)[row2][col] = EMPTY, ROUND
		}
	}
}

// ======================================================================

func calcNorthWeight(g Grid) int {
	weight := 0

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if (*g.rocks)[i][j] == ROUND {
				weight += (g.rows - i)
			}
		}
	}
	return weight
}
