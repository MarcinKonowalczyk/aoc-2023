package day21

import "fmt"

type Field rune

const (
	START  Field = 'S'
	GARDEN Field = '.'
	ROCK   Field = '#'
)

type Garden struct {
	grid [][]Field
	rows int
	cols int
}

func (g Garden) String() string {
	s := "Garden:\n"
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			s += string(g.grid[i][j])
		}
		s += "\n"
	}
	return s
}

func parseLines(lines []string) (Garden, error) {
	rows := len(lines)
	if rows == 0 {
		return Garden{}, fmt.Errorf("no lines")
	}

	cols := len(lines[0])
	grid := make([][]Field, rows)
	for i, line := range lines {
		grid[i] = make([]Field, cols)
		if len(line) != cols {
			return Garden{}, fmt.Errorf("line %d has %d columns, expected %d", i, len(line), cols)
		}
		for j, r := range line {
			switch r {
			case 'S':
				grid[i][j] = START
			case '.':
				grid[i][j] = GARDEN
			case '#':
				grid[i][j] = ROCK
			default:
				return Garden{}, fmt.Errorf("invalid rune '%c'", r)
			}
		}
	}

	return Garden{grid: grid, rows: len(lines), cols: cols}, nil

}
