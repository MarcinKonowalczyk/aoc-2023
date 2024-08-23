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
		// fmt.Println(l)
		tg.move(l.dir, l.distance)
	}

	if !tg.isLoop() {
		return 0, fmt.Errorf("path does not loop")
	}

	p, err := findInteriorPoint(tg)
	if err != nil {
		return 0, err
	}

	// fmt.Println("Interior point:", p)

	tg2 := tg.Copy()

	floodFill(tg2.grid, p.Add(utils.Point2{X: -tg2.min.X, Y: -tg2.min.Y}))

	// fmt.Println("Filled grid:")
	// fmt.Println(tg2)

	count := 0
	for y := tg2.min.Y; y <= tg2.max.Y; y++ {
		for x := tg2.min.X; x <= tg2.max.X; x++ {
			xx := x - tg2.min.X
			yy := y - tg2.min.Y
			if tg2.grid[yy][xx] {
				count++
			}
		}
	}

	return count, nil
}

type turtle_grid struct {
	grid [][]bool
	min  utils.Point2
	max  utils.Point2
	path []utils.Point2
}

func (tg *turtle_grid) Copy() turtle_grid {
	new_tg := turtle_grid{
		grid: make([][]bool, len(tg.grid)),
	}
	for i := range tg.grid {
		new_tg.grid[i] = make([]bool, len(tg.grid[i]))
		copy(new_tg.grid[i], tg.grid[i])
	}
	new_tg.min = tg.min
	new_tg.max = tg.max
	new_tg.path = make([]utils.Point2, len(tg.path))
	copy(new_tg.path, tg.path)
	return new_tg
}

func (tg *turtle_grid) addPoint(p utils.Point2) error {
	if tg.grid == nil {
		tg.grid = make([][]bool, 0)
		tg.min = utils.Point2{X: 0, Y: 0}
		tg.max = utils.Point2{X: 0, Y: 0}
	}
	if tg.path == nil {
		tg.path = make([]utils.Point2, 0)
	}

	// Expand grid if necessary

	if p.X >= tg.max.X {
		// Need to expand to the right
		pad_x := p.X - tg.max.X + 1
		for i := range tg.grid {
			tg.grid[i] = append(tg.grid[i], make([]bool, pad_x)...)
		}
		tg.max.X = p.X
	}
	if p.X < tg.min.X {
		// Need to expand to the left
		pad_x := tg.min.X - p.X
		for i := range tg.grid {
			tg.grid[i] = append(make([]bool, pad_x), tg.grid[i]...)
		}
		tg.min.X = p.X
	}
	if p.Y >= tg.max.Y {
		// Need to expand down
		pad_y := p.Y - tg.max.Y + 1
		for i := 0; i < pad_y; i++ {
			tg.grid = append(tg.grid, make([]bool, tg.max.X-tg.min.X+1))
		}
		tg.max.Y = p.Y
	}
	if p.Y < tg.min.Y {
		// Need to expand up
		pad_y := tg.min.Y - p.Y
		for i := 0; i < pad_y; i++ {
			new_row := make([]bool, tg.max.X-tg.min.X+1)
			for j := range new_row {
				new_row[j] = false
			}
			tg.grid = append([][]bool{new_row}, tg.grid...)
		}
		tg.min.Y = p.Y
	}

	if len(tg.path) == 0 {
		// First point
		tg.path = append(tg.path, p)
		tg.grid[p.Y][p.X] = true
		tg.min = p
		tg.max = p
	} else {
		prev_p := tg.path[len(tg.path)-1]
		if prev_p.X == p.X { // Vertical line
			abs_diff, sign := utils.AbsDiffAndSignBinary(p.Y, prev_p.Y)
			for d := 1; d <= abs_diff; d++ {
				p := prev_p.AddY(d * sign).Sub(tg.min)
				tg.grid[p.Y][p.X] = true
			}
		} else if prev_p.Y == p.Y { // Horizontal line
			abs_diff, sign := utils.AbsDiffAndSignBinary(p.X, prev_p.X)
			for d := 1; d <= abs_diff; d++ {
				p := prev_p.AddX(d * sign).Sub(tg.min)
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
		new_p = prev_p.AddX(int(distance))
	case DOWN:
		new_p = prev_p.AddY(int(distance))
	case LEFT:
		new_p = prev_p.AddX(-int(distance))
	case UP:
		new_p = prev_p.AddY(-int(distance))
	}

	tg.addPoint(new_p)
}

func (tg turtle_grid) String() string {
	s := ""
	for y := tg.min.Y; y <= tg.max.Y; y++ {
		for x := tg.min.X; x <= tg.max.X; x++ {
			xx := x - tg.min.X
			yy := y - tg.min.Y
			if tg.grid[yy][xx] {
				s += "#"
			} else {
				s += "."
			}
		}
		if y < tg.max.Y {
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

func findInteriorPoint(tg turtle_grid) (utils.Point2, error) {
	// Find any point in the grid which is interior to the loop
	if !tg.isLoop() {
		return utils.Point2{}, fmt.Errorf("path does not loop")
	}

	// Find top-left corner of the loop
	xy := utils.ArrayMap(tg.path, func(p utils.Point2) int {
		return p.X + p.Y
	})
	_, I, err := utils.ArrayMin(xy)
	if err != nil {
		return utils.Point2{}, err
	}

	p := tg.path[I]

	// Go into the interior
	p.X += 1
	p.Y += 1

	return p, nil
}

func floodFill(grid [][]bool, p utils.Point2) {
	// Flood fill the grid starting from point p
	queue := []utils.Point2{p}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		// Check if point is valid
		if p.Y < 0 || p.Y >= len(grid) || p.X < 0 || p.X >= len(grid[0]) {
			continue
		}

		if grid[p.Y][p.X] {
			// Already visited
			continue
		}

		// Mark as visited
		grid[p.Y][p.X] = true

		// Add neighbors to queue
		queue = append(queue, utils.Point2{X: p.X + 1, Y: p.Y})
		queue = append(queue, utils.Point2{X: p.X - 1, Y: p.Y})
		queue = append(queue, utils.Point2{X: p.X, Y: p.Y + 1})
		queue = append(queue, utils.Point2{X: p.X, Y: p.Y - 1})
	}
}
