package day18

import (
	"aoc2023/utils"
	"fmt"
)

func Main(part int, lines []string, verbose bool) (n int, err error) {
	if part == 1 {
		return main_1(lines, verbose)
	} else if part == 2 {
		return main_2(lines, verbose)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string, verbose bool) (n int, err error) {
	parsed_lines, err := utils.ArrayMapWithError(lines, parseLine)
	if err != nil {
		return 0, err
	}

	tg := turtle_grid{}
	tg.addPoint(utils.Point2{X: 0, Y: 0})

	for _, l := range parsed_lines {
		// fmt.Println(l)
		tg.move(l.dir, l.distance)
		// fmt.Println(tg)
	}

	if !tg.isLoop() {
		return 0, fmt.Errorf("path does not loop")
	}

	// fmt.Println(tg)

	p, err := findInteriorPoint(tg)
	if err != nil {
		return 0, err
	}

	fmt.Println("Interior point:", p)

	tg2 := tg.Copy()

	floodFill(tg2.grid, p.Add(utils.Point2{X: -tg2.min.X, Y: -tg2.min.Y}))

	fmt.Println("Filled grid:")
	fmt.Println(tg2)

	count := 0
	for y := tg2.min.Y; y <= tg2.max.Y; y++ {
		for x := tg2.min.X; x <= tg2.max.X; x++ {
			p := utils.Point2{X: x, Y: y}.Sub(tg2.min)
			if _, ok := tg2.grid.points[p]; ok {
				count++
			}
		}
	}

	return count, nil
}

type sparse_grid struct {
	points map[utils.Point2]struct{}
	rows   int
	cols   int
}

func (sg sparse_grid) Copy() sparse_grid {
	new_sg := sparse_grid{
		points: make(map[utils.Point2]struct{}),
		rows:   sg.rows,
		cols:   sg.cols,
	}
	for k, v := range sg.points {
		new_sg.points[k] = v
	}
	return new_sg
}

func (sg *sparse_grid) addPoint(p utils.Point2) {
	if p.X < 0 || p.Y < 0 {
		panic(fmt.Sprintf("invalid point %v", p))
	}
	if p.X >= sg.cols || p.Y >= sg.rows {
		panic(fmt.Sprintf("invalid point %v", p))
	}

	sg.points[p] = struct{}{}
}

type turtle_grid struct {
	grid sparse_grid
	min  utils.Point2
	max  utils.Point2
	path []utils.Point2
}

func (tg *turtle_grid) Copy() turtle_grid {
	new_tg := turtle_grid{
		grid: tg.grid.Copy(),
	}
	new_tg.min = tg.min
	new_tg.max = tg.max
	new_tg.path = make([]utils.Point2, len(tg.path))
	copy(new_tg.path, tg.path)
	return new_tg
}

func (tg *turtle_grid) addPoint(p utils.Point2) error {
	if tg.path == nil {
		tg.path = make([]utils.Point2, 0)
		tg.min = utils.Point2{X: 0, Y: 0}
		tg.max = utils.Point2{X: 0, Y: 0}
		tg.grid = sparse_grid{
			points: make(map[utils.Point2]struct{}),
			rows:   0,
			cols:   0,
		}
	}

	// Expand grid if necessary

	if p.X >= tg.max.X {
		// Need to expand to the right
		pad_x := p.X - tg.max.X + 1
		tg.grid.cols += pad_x
		tg.max.X = p.X
	}
	if p.X < tg.min.X {
		// Need to expand to the left
		fmt.Println("Expanding to the left")
		fmt.Println(tg)
		pad_x := tg.min.X - p.X
		tg.grid.cols += pad_x
		new_points := make(map[utils.Point2]struct{})
		for k := range tg.grid.points {
			new_points[k.AddX(pad_x)] = struct{}{}
		}
		tg.grid.points = new_points
		tg.min.X = p.X
		fmt.Println("After expanding to the left")
		fmt.Println(tg)
	}
	if p.Y >= tg.max.Y {
		// Need to expand down
		pad_y := p.Y - tg.max.Y + 1
		tg.grid.rows += pad_y
		tg.max.Y = p.Y
	}
	if p.Y < tg.min.Y {
		// Need to expand up
		pad_y := tg.min.Y - p.Y
		tg.grid.rows += pad_y
		new_points := make(map[utils.Point2]struct{})
		for k := range tg.grid.points {
			new_points[k.AddY(pad_y)] = struct{}{}
		}
		tg.grid.points = new_points
		tg.min.Y = p.Y
	}

	if len(tg.path) == 0 {
		// First point
		tg.path = append(tg.path, p)
		tg.grid.addPoint(p)
		tg.min = p
		tg.max = p
	} else {
		prev_p := tg.path[len(tg.path)-1]
		if prev_p.X == p.X { // Vertical line
			abs_diff, sign := utils.AbsDiffAndSignBinary(p.Y, prev_p.Y)
			for d := 1; d <= abs_diff; d++ {
				p := prev_p.AddY(d * sign).Sub(tg.min)
				tg.grid.addPoint(p)
			}
		} else if prev_p.Y == p.Y { // Horizontal line
			abs_diff, sign := utils.AbsDiffAndSignBinary(p.X, prev_p.X)
			for d := 1; d <= abs_diff; d++ {
				p := prev_p.AddX(d * sign).Sub(tg.min)
				tg.grid.addPoint(p)
			}
		} else {
			return fmt.Errorf("invalid path")
		}
		// Actually append the point
		tg.path = append(tg.path, p)
	}

	return nil
}

func move(p utils.Point2, dir direction, distance uint) utils.Point2 {
	var new_p utils.Point2
	switch dir {
	case RIGHT:
		new_p = p.AddX(int(distance))
	case DOWN:
		new_p = p.AddY(int(distance))
	case LEFT:
		new_p = p.AddX(-int(distance))
	case UP:
		new_p = p.AddY(-int(distance))
	}

	return new_p
}

func (tg *turtle_grid) move(dir direction, distance uint) {
	if len(tg.path) == 0 {
		return
	}

	prev_p := tg.path[len(tg.path)-1]
	new_p := move(prev_p, dir, distance)

	tg.addPoint(new_p)
}

func (tg turtle_grid) String() string {
	s := ""
	for y := tg.min.Y; y <= tg.max.Y; y++ {
		for x := tg.min.X; x <= tg.max.X; x++ {
			p := utils.Point2{X: x, Y: y}.Sub(tg.min)
			if _, ok := tg.grid.points[p]; ok {
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

func floodFill(grid sparse_grid, p utils.Point2) {
	// Flood fill the grid starting from point p
	queue := []utils.Point2{p}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		// Check if point is valid
		if p.Y < 0 || p.Y >= grid.rows || p.X < 0 || p.X >= grid.cols {
			continue
		}

		if _, ok := grid.points[p]; ok {
			// Already visited
			continue
		}

		// Mark as visited
		// grid[p.Y][p.X] = true
		grid.addPoint(p)

		// Add neighbors to queue
		queue = append(queue, p.AddX(1))
		queue = append(queue, p.AddX(-1))
		queue = append(queue, p.AddY(1))
		queue = append(queue, p.AddY(-1))
	}
}
