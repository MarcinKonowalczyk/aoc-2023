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

	boundary := boundaryArea(points)
	fmt.Println("Boundary area:", boundary)

	// We should never go over n points since we're always reducing at least two point
	MAX_ITER := len(points) / 2
	// MAX_ITER := 1

	total_area := 0
	var area int
	for i := 0; i < MAX_ITER; i++ {
		points, area = reduceOnce(points)
		fmt.Println("Reduced area:", area)
		total_area += area
		if !checkPath(points) {
			panic("Path is not valid")
		}

		if len(points) <= 4 {
			fmt.Println("Reduced to a rectangle")
			break
		}
	}

	if len(points) != 4 {
		return -1, fmt.Errorf("failed to reduce to a rectangle")
	}

	fmt.Println(len(points))
	fmt.Println(points)

	rect_area := rectangleArea(points[1], points[3])
	fmt.Println("Rectangle area:", rect_area)

	total_area += rect_area
	total_area += boundary

	fmt.Println("Total area:", total_area)

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

func reduceOnce(ps []utils.Point2) ([]utils.Point2, int) {
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

func reduceAtIndex(ps []utils.Point2, i int) ([]utils.Point2, int) {
	d1, d2, d3, d4, d5 := getDirs(ps, i)
	p0 := ps[(i-2+len(ps))%len(ps)]
	p1 := ps[(i-1+len(ps))%len(ps)]
	p2 := ps[i]
	p3 := ps[(i+1)%len(ps)]
	p4 := ps[(i+2)%len(ps)]
	p5 := ps[(i+3)%len(ps)]
	ad0 := p0.L1(p1)
	ad1 := p1.L1(p2)
	ad2 := p2.L1(p3)
	ad3 := p3.L1(p4)
	ad4 := p4.L1(p5)
	// ad_delta := utils.IntMin(ad1, ad3)
	fmt.Println("Points:", p0, p1, p2, p3, p4, p5)
	fmt.Println("Directions:", d1, d2, d3, d4, d5)
	fmt.Println("Distances:", ad0, ad1, ad2, ad3, ad4)
	if ad1 < ad3 {
		fmt.Println("ad1 < ad3")
	} else if ad1 > ad3 {
		fmt.Println("ad1 > ad3")
	} else {
		fmt.Println("ad1 == ad3")
	}

	var ra1, ra2 int

	switch d3 {
	case RIGHT, LEFT:
		switch d2 {
		case DOWN:
			if ad1 < ad3 {
				new_p3 := p3.AddY(-ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
				if p0.X < p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.X > p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			} else if ad1 > ad3 {
				new_p2 := p2.AddY(-ad3)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.X < p1.X {
					ra1 = (ad3 - 1) * (ad2 - 1)
					ra2 = (ad3 - 1) * (ad2 - 1)
				} else if p0.X > p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			} else {
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.X < p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.X > p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			}
		case UP:
			if ad1 < ad3 {
				new_p3 := p3.AddY(ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
				if p0.X < p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.X > p1.X {
					ra1 = (ad1 - 1) * (ad2 - 1)
					ra2 = (ad1-1)*(ad2-1) - (p0.X - p1.X)
				} else {
					panic("Invalid direction")
				}

			} else if ad1 > ad3 {
				new_p2 := p2.AddY(ad3)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.X < p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.X > p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			} else {
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.X < p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.X > p1.X {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			}
		default:
			panic("Invalid direction")
		}
	case DOWN, UP:
		switch d2 {
		case RIGHT:
			if ad1 < ad3 {
				new_p3 := p3.AddX(-ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
				if p0.Y < p1.Y {
					ra1 = (ad1 - 1) * (ad2 - 1)
					ra2 = (ad1 - 1) * (ad2 - 1)
				} else if p0.Y > p1.Y {
					ra1 = (ad1 - 1) * (ad2 - 1)
					ra2 = (ad1-1)*(ad2-1) - (p0.Y - p1.Y)
				} else {
					panic("Invalid direction")
				}
			} else if ad1 > ad3 {
				new_p2 := p2.AddX(-ad3)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.Y < p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.Y > p1.Y {
					ra1 = (ad3 - 1) * (ad2 - 1)
					ra2 = (ad3 - 1) * (ad2 - 1)
				} else {
					panic("Invalid direction")
				}
			} else {
				// Remove all 4 points. No need to keep the U shape
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.Y < p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.Y > p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			}
		case LEFT:
			if ad1 < ad3 {
				new_p3 := p3.AddX(ad1)
				ps[(i+1)%len(ps)] = new_p3
				ps, _ = utils.ArrayRemoveIndices(ps, []int{i, (i - 1 + len(ps)) % len(ps)})
				if p0.Y < p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.Y > p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			} else if ad1 > ad3 {
				new_p2 := p2.AddX(ad3)
				ps[i] = new_p2
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.Y < p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.Y > p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			} else {
				// Remove all 4 points. No need to keep the U shape
				ps, _ = utils.ArrayRemoveIndices(ps, []int{(i - 1 + len(ps)) % len(ps), i, (i + 1) % len(ps), (i + 2) % len(ps)})
				if p0.Y < p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else if p0.Y > p1.Y {
					utils.Panicf("Calculating reduced area for %v::%v not implemented", d2, d3)
				} else {
					panic("Invalid direction")
				}
			}
		default:
			panic("Invalid direction")
		}
	}

	reduced_area := ra1 + ra2
	// fmt.Println("ad0:", ad0)
	// fmt.Println("ad1:", ad1)
	// fmt.Println("ad2:", ad2)
	// fmt.Println("ad3:", ad3)
	// fmt.Println("ad4:", ad4)
	// fmt.Println("1:", (ad1-1)*(ad2-1))
	// fmt.Println("2:", (ad1-1)*(ad2-1) - ad
	// reduced_area := utils.Ternary(ad1 < ad3, (ad1-1)*(ad2-1), (ad3-1)*(ad2-1))

	return ps, reduced_area
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

func rectangleArea(p1, p2 utils.Point2) int {
	if p1.X == p2.X || p1.Y == p2.Y {
		panic("Invalid rectangle")
	}
	dx := utils.AbsDiff(p1.X, p2.X) - 1
	dy := utils.AbsDiff(p1.Y, p2.Y) - 1
	return dx * dy
}

func boundaryArea(ps []utils.Point2) int {
	area := 0
	for i := 0; i < len(ps); i++ {
		p1 := ps[i]
		p2 := ps[(i+1)%len(ps)]
		area += p1.L1(p2)
	}
	return area
}
