package day21

import (
	"aoc2023/utils"
	"fmt"
)

type Field rune

const (
	GARDEN Field = '.'
	ROCK   Field = '#'
)

type Garden struct {
	grid      [][]Field
	rows      int
	cols      int
	positions []utils.Point2
}

func (g Garden) String() string {
	s := "Garden:\n"
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			for _, p := range g.positions {
				if i == p.X && j == p.Y {
					s += "O"
					goto next
				}
			}
			s += string(g.grid[i][j])
		next:
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
	start_pos := utils.Point2{-1, -1}
	for i, line := range lines {
		grid[i] = make([]Field, cols)
		if len(line) != cols {
			return Garden{}, fmt.Errorf("line %d has %d columns, expected %d", i, len(line), cols)
		}
		for j, r := range line {
			switch r {
			case 'S':
				start_pos = utils.Point2{i, j}
				grid[i][j] = GARDEN
			case '.':
				grid[i][j] = GARDEN
			case '#':
				grid[i][j] = ROCK
			default:
				return Garden{}, fmt.Errorf("invalid rune '%c'", r)
			}
		}
	}
	if start_pos.X == -1 {
		return Garden{}, fmt.Errorf("no starting position")
	}

	return Garden{grid: grid, rows: len(lines), cols: cols, positions: []utils.Point2{start_pos}}, nil

}
