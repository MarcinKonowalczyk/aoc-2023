package day18

import (
	"aoc2023/utils"
	"fmt"
)

func main_2(lines []string) (n int, err error) {
	parsed_lines, err := utils.ArrayMapWithError(lines, parseLine)
	if err != nil {
		return 0, err
	}

	// parsed_lines, err = utils.ArrayMapWithError(parsed_lines, extractCorrectParams)
	// if err != nil {
	// 	return 0, err
	// }

	points := make([]utils.Point2, 0)
	points = append(points, utils.Point2{X: 0, Y: 0})

	for _, l := range parsed_lines {
		// fmt.Println(l)
		// tg.move(l.dir, l.distance)
		// fmt.Println(tg)

		points = append(points, move(points[len(points)-1], l.dir, l.distance))
	}

	if points[0] != points[len(points)-1] {
		return 0, fmt.Errorf("path does not loop")
	}

	points = points[:len(points)-1] // We don't need the last point anymore

	fmt.Println(points)

	// We should never go over n points since we're always reducing at least two point
	MAX_ITER := len(points) / 2

	for i := 0; i < MAX_ITER; i++ {
		points = reduceOnce(points)
		if !checkPath(points) {
			panic("Path is not valid")
		}
		fmt.Println("---")
		fmt.Println(points)

		if len(points) <= 4 {
			fmt.Println("Reduced to a rectangle")
			break
		}
	}

	fmt.Println(len(points))
	fmt.Println(points)

	// fmt.Println(tg)

	// p, err := findInteriorPoint(tg)
	// if err != nil {
	// 	return 0, err
	// }

	// fmt.Println("Interior point:", p)

	// tg2 := tg.Copy()

	// fmt.Println("Running flood fill")
	// floodFill(tg2.grid, p.Add(utils.Point2{X: -tg2.min.X, Y: -tg2.min.Y}))

	// fmt.Println("Filled grid:")
	// fmt.Println(tg2)

	count := 0
	// for y := tg2.min.Y; y <= tg2.max.Y; y++ {
	// 	for x := tg2.min.X; x <= tg2.max.X; x++ {
	// 		p := utils.Point2{X: x, Y: y}.Sub(tg2.min)
	// 		if _, ok := tg2.grid.points[p]; ok {
	// 			count++
	// 		}
	// 	}
	// }

	return count, nil
}

func reduceOnce(ps []utils.Point2) []utils.Point2 {
	// Find a reducible edge
	reducible_I := findReducibleEdge(ps)
	fmt.Println("Reducible edge:", reducible_I)
	if reducible_I == -1 {
		panic("No reducible edge found")
	}
	return reduceAtIndex(ps, reducible_I)
}

// Find an arrangement of edges where we can apply a reduction
// This is a U-shaped path where and edge goes, for example, right, then down, then left.
// In such case teh down edge can be reduced
// Returns the index of the reducible edge, in this case the index pointing to the beginning of the down edge
func findReducibleEdge(ps []utils.Point2) int {
	for i := 0; i < len(ps); i++ {
		i0 := (i - 2 + len(ps)) % len(ps)
		i1 := (i - 1 + len(ps)) % len(ps)
		i2 := i
		i3 := (i + 1) % len(ps)
		i4 := (i + 2) % len(ps)
		i5 := (i + 3) % len(ps)

		p1 := ps[i1]
		p2 := ps[i2]
		p3 := ps[i3]
		p4 := ps[i4]

		d2 := getDirection(p1, p2)
		d3 := getDirection(p2, p3)
		d4 := getDirection(p3, p4)

		if !isUShaped(d2, d3, d4) {
			utils.Cprintf(utils.Yellow, "Edge at %d is not U-shaped\n", i)
			continue
		} else {
			utils.Cprintf(utils.Green, "Edge at %d is U-shaped\n", i)
		}

		// We need to check if any other point is inside the U shape rectangle
		has_point_inside := false
		for j := 0; j < len(ps); j++ {
			if j == i0 || j == i1 || j == i2 || j == i3 || j == i4 || j == i5 {
				// Skip the points of the U shape
				continue
			}
			if utils.PointInRectangle(ps[j], p1, p3) {
				fmt.Println("Point inside:", ps[j])
				has_point_inside = true
				break
			}
		}

		if has_point_inside {
			utils.Cprintf(utils.Yellow, "Edge at %d has a point inside\n", i)
			continue
		} else {
			utils.Cprintf(utils.Green, "Edge at %d has no point inside\n", i)
			return i
		}
	}
	return -1

}

func getDirection(p0, p1 utils.Point2) direction {
	if p0.X == p1.X {
		if p0.Y < p1.Y {
			return DOWN
		} else {
			return UP
		}
	} else {
		if p0.X < p1.X {
			return RIGHT
		} else {
			return LEFT
		}
	}
}

func isUShaped(d1, d2, d3 direction) bool {
	if d1 == RIGHT {
		if (d2 == DOWN && d3 == LEFT) || (d2 == UP && d3 == LEFT) {
			return true
		}
	} else if d1 == DOWN {
		if (d2 == LEFT && d3 == UP) || (d2 == RIGHT && d3 == UP) {
			return true
		}
	} else if d1 == LEFT {
		if (d2 == UP && d3 == RIGHT) || (d2 == DOWN && d3 == RIGHT) {
			return true
		}
	} else if d1 == UP {
		if (d2 == RIGHT && d3 == DOWN) || (d2 == LEFT && d3 == DOWN) {
			return true
		}
	}
	return false
}

func getDirs(ps []utils.Point2, i int) (direction, direction, direction, direction, direction) {
	p0 := ps[(i-2+len(ps))%len(ps)]
	p1 := ps[(i-1+len(ps))%len(ps)]
	p2 := ps[i]
	p3 := ps[(i+1)%len(ps)]
	p4 := ps[(i+2)%len(ps)]
	p5 := ps[(i+3)%len(ps)]

	d1 := getDirection(p0, p1)
	d2 := getDirection(p1, p2)
	d3 := getDirection(p2, p3)
	d4 := getDirection(p3, p4)
	d5 := getDirection(p4, p5)

	return d1, d2, d3, d4, d5
}

func reduceAtIndex(ps []utils.Point2, i int) []utils.Point2 {
	d1, d2, d3, d4, d5 := getDirs(ps, i)
	p1 := ps[(i-1+len(ps))%len(ps)]
	p2 := ps[i]
	p3 := ps[(i+1)%len(ps)]
	p4 := ps[(i+2)%len(ps)]
	ad1, ad2 := p1.L1(p2), p3.L1(p4)
	// ad_delta := utils.IntMin(ad1, ad2)
	fmt.Println("Directions:", d1, d2, d3, d4, d5)
	if ad1 < ad2 {
		fmt.Println("ad1 < ad2")
	} else if ad1 > ad2 {
		fmt.Println("ad1 > ad2")
	} else {
		fmt.Println("ad1 == ad2")
	}

	switch d3 {
	case RIGHT, LEFT:
		switch d2 {
		case DOWN:
			if ad1 < ad2 {
				new_p3 := p3.AddY(-ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
			} else if ad1 > ad2 {
				new_p2 := p2.AddY(-ad2)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
			} else {
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
			}
		case UP:
			if ad1 < ad2 {
				new_p3 := p3.AddY(ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
			} else if ad1 > ad2 {
				new_p2 := p2.AddY(ad2)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
			} else {
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
			}
		default:
			panic("Invalid direction")
		}
	case DOWN, UP:
		switch d2 {
		case RIGHT:
			if ad1 < ad2 {
				new_p3 := p3.AddX(-ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
			} else if ad1 > ad2 {
				new_p2 := p2.AddX(-ad2)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
			} else {
				// Remove all 4 points. No need to keep the U shape
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
			}
		case LEFT:
			if ad1 < ad2 {
				new_p3 := p3.AddX(ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
			} else if ad1 > ad2 {
				new_p2 := p2.AddX(ad2)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
			} else {
				// Remove all 4 points. No need to keep the U shape
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
			}
		default:
			panic("Invalid direction")
		}
	}

	return ps
}

// Check that each segment of the path is either horizontal or vertical
func checkPath(ps []utils.Point2) bool {
	for i := 0; i < len(ps); i++ {
		p1 := ps[i]
		p2 := ps[(i+1)%len(ps)]
		if p1.X != p2.X && p1.Y != p2.Y {
			return false
		}
	}
	return true
}
