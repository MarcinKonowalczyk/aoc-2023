package day18

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
	parsed_lines, err := utils.ArrayMapWithError(lines, parseLine)
	if err != nil {
		return 0, err
	}

	tg := turtle_grid{}
	tg.addPoint(utils.Point2{X: 0, Y: 0})

	for _, l := range parsed_lines {
		fmt.Println(l)
		tg.move(l.dir, l.distance)
	}

	if !tg.isLoop() {
		return 0, fmt.Errorf("path does not loop")
	}

	return 0, nil
}

type turtle_grid struct {
	grid  [][]bool
	min_x int
	min_y int
	max_x int
	max_y int
	path  []utils.Point2
}

func (tg *turtle_grid) addPoint(p utils.Point2) error {
	if tg.grid == nil {
		tg.grid = make([][]bool, 0)
		tg.min_x = 0
		tg.min_y = 0
		tg.max_x = 0
		tg.max_y = 0
	}
	if tg.path == nil {
		tg.path = make([]utils.Point2, 0)
	}

	// Expand grid if necessary

	if p.X >= tg.max_x {
		// Need to expand to the right
		pad_x := p.X - tg.max_x + 1
		for i := range tg.grid {
			tg.grid[i] = append(tg.grid[i], make([]bool, pad_x)...)
		}
		tg.max_x = p.X
	}
	if p.X < tg.min_x {
		// Need to expand to the left
		pad_x := tg.min_x - p.X
		for i := range tg.grid {
			tg.grid[i] = append(make([]bool, pad_x), tg.grid[i]...)
		}
		tg.min_x = p.X
	}
	if p.Y >= tg.max_y {
		// Need to expand down
		pad_y := p.Y - tg.max_y + 1
		for i := 0; i < pad_y; i++ {
			tg.grid = append(tg.grid, make([]bool, tg.max_x-tg.min_x+1))
		}
		tg.max_y = p.Y
	}
	if p.Y < tg.min_y {
		// Need to expand up
		pad_y := tg.min_y - p.Y
		for i := 0; i < pad_y; i++ {
			new_row := make([]bool, tg.max_x-tg.min_x+1)
			for j := range new_row {
				new_row[j] = false
			}
			tg.grid = append([][]bool{new_row}, tg.grid...)
			// tg.grid = append(make([][]bool, 1), tg.grid...)
		}
		tg.min_y = p.Y
	}

	if len(tg.path) == 0 {
		// First point
		tg.path = append(tg.path, p)
		tg.grid[p.Y][p.X] = true
		tg.min_x = p.X
		tg.min_y = p.Y
		tg.max_x = p.X
		tg.max_y = p.Y
	} else {
		prev_p := tg.path[len(tg.path)-1]
		if prev_p.X == p.X { // Vertical line
			abs_diff, sign := utils.AbsDiffAndSignBinary(p.Y, prev_p.Y)
			for d := 1; d <= abs_diff; d++ {
				p := utils.Point2{X: p.X, Y: prev_p.Y + d*sign}
				p.X -= tg.min_x
				p.Y -= tg.min_y
				tg.grid[p.Y][p.X] = true
			}
		} else if prev_p.Y == p.Y { // Horizontal line
			abs_diff, sign := utils.AbsDiffAndSignBinary(p.X, prev_p.X)
			for d := 1; d <= abs_diff; d++ {
				p := utils.Point2{X: prev_p.X + d*sign, Y: p.Y}
				p.X -= tg.min_x
				p.Y -= tg.min_y
				tg.grid[p.Y][p.X] = true
			}
		} else {
			return fmt.Errorf("invalid path")
		}
		// Actually append the point
		tg.path = append(tg.path, p)
	}

	return nil
}

func (tg *turtle_grid) move(dir direction, distance uint) {
	if len(tg.path) == 0 {
		return
	}

	prev_p := tg.path[len(tg.path)-1]
	var new_p utils.Point2
	switch dir {
	case RIGHT:
		new_p = utils.Point2{X: prev_p.X + int(distance), Y: prev_p.Y}
	case DOWN:
		new_p = utils.Point2{X: prev_p.X, Y: prev_p.Y + int(distance)}
	case LEFT:
		new_p = utils.Point2{X: prev_p.X - int(distance), Y: prev_p.Y}
	case UP:
		new_p = utils.Point2{X: prev_p.X, Y: prev_p.Y - int(distance)}
	}

	tg.addPoint(new_p)
}

func (tg turtle_grid) String() string {
	s := ""
	for y := tg.min_y; y <= tg.max_y; y++ {
		for x := tg.min_x; x <= tg.max_x; x++ {
			xx := x - tg.min_x
			yy := y - tg.min_y
			if tg.grid[yy][xx] {
				s += "#"
			} else {
				s += "."
			}
		}
		if y < tg.max_y {
			s += "\n"
		}
	}
	return s
}

// Check if the path loops back on itself
func (tg *turtle_grid) isLoop() bool {
	if len(tg.path) < 4 {
		return false
	}

	// Check if the last point is the same as the first point
	if tg.path[0].X != tg.path[len(tg.path)-1].X || tg.path[0].Y != tg.path[len(tg.path)-1].Y {
		return false
	}

	return true
}
