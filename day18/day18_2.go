package day18

import (
	"aoc2023/utils"
	"fmt"
)

func main_2(lines []string, verbose bool) (n int, err error) {
	parsed_lines, err := utils.ArrayMapWithError(lines, parseLine)
	if err != nil {
		return 0, err
	}

	parsed_lines, err = utils.ArrayMapWithError(parsed_lines, extractCorrectParams)
	if err != nil {
		return 0, err
	}

	points := make([]utils.Point2, 0)
	points = append(points, utils.Point2{X: 0, Y: 0})

	for _, l := range parsed_lines {
		points = append(points, move(points[len(points)-1], l.dir, l.distance))
	}

	if points[0] != points[len(points)-1] {
		return 0, fmt.Errorf("path does not loop")
	}

	points = points[:len(points)-1] // We don't need the last point anymore

	// We should never go over n points since we're always reducing at least two point
	MAX_ITER := len(points) / 2
	// MAX_ITER := 1

	total_area := 0
	var area int
	for i := 0; i < MAX_ITER; i++ {
		points, area = reduceOnce(points)
		total_area += area
		if !checkPath(points) {
			panic("Path is not valid")
		}

		if len(points) <= 4 {
			// fmt.Println("Reduced to a rectangle")
			break
		}
	}

	if len(points) != 4 {
		return -1, fmt.Errorf("failed to reduce to a rectangle")
	}

	rect_area := rectangleArea(points[1], points[3])
	total_area += rect_area

	return total_area, nil
}

func reduceOnce(ps []utils.Point2) ([]utils.Point2, int) {
	// Find a reducible edge
	reducible_I := findReducibleEdge(ps)
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

		if !isClockwiseU(d2, d3, d4) {
			// utils.Cprintf(utils.Yellow, "Edge at %d is not U-shaped\n", i)
			continue
		} else {
			// utils.Cprintf(utils.Green, "Edge at %d is U-shaped\n", i)
		}

		// We need to check if any other point is inside the U shape rectangle
		has_point_inside := false
		for j := 0; j < len(ps); j++ {
			if j == i0 || j == i1 || j == i2 || j == i3 || j == i4 || j == i5 {
				// Skip the points of the U shape
				continue
			}
			if utils.PointInRectangle(ps[j], p1, p3) {
				// fmt.Println("Point inside:", ps[j])
				has_point_inside = true
				break
			}
		}

		if has_point_inside {
			// utils.Cprintf(utils.Yellow, "Edge at %d has a point inside\n", i)
			continue
		} else {
			// utils.Cprintf(utils.Green, "Edge at %d has no point inside\n", i)
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

func isClockwiseU(d1, d2, d3 direction) bool {
	return (d1 == RIGHT && d2 == DOWN && d3 == LEFT) ||
		(d1 == DOWN && d2 == LEFT && d3 == UP) ||
		(d1 == LEFT && d2 == UP && d3 == RIGHT) ||
		(d1 == UP && d2 == RIGHT && d3 == DOWN)
}

func reduceAtIndex(ps []utils.Point2, i int) ([]utils.Point2, int) {
	i0 := (i - 2 + len(ps)) % len(ps)
	i1 := (i - 1 + len(ps)) % len(ps)
	i2 := i
	i3 := (i + 1) % len(ps)
	i4 := (i + 2) % len(ps)
	i5 := (i + 3) % len(ps)
	p0 := ps[i0]
	p1 := ps[i1]
	p2 := ps[i2]
	p3 := ps[i3]
	p4 := ps[i4]
	p5 := ps[i5]
	d0 := getDirection(p0, p1)
	d1 := getDirection(p1, p2)
	d2 := getDirection(p2, p3)
	// d3 := getDirection(p3, p4)
	d4 := getDirection(p4, p5)
	ad0 := p0.L1(p1)
	ad1 := p1.L1(p2)
	ad2 := p2.L1(p3)
	ad3 := p3.L1(p4)
	ad4 := p4.L1(p5)
	// fmt.Println("Points:", p0, p1, p2, p3, p4, p5)
	// fmt.Println("Directions:", d0, d1, d2, d3, d4)
	// fmt.Println("Distances:", ad0, ad1, ad2, ad3, ad4)

	var p utils.Point2 = utils.Point2{X: 0, Y: 0}
	if (d1 == UP && d2 == RIGHT) || (d1 == DOWN && d2 == LEFT) {
		sign := utils.BoolToSign(d1 == UP)
		if ad1 < ad3 {
			p = utils.Point2{X: 0, Y: sign * ad1}
		} else if ad1 > ad3 {
			p = utils.Point2{X: 0, Y: sign * ad3}
		}
	} else if (d1 == RIGHT && d2 == DOWN) || (d1 == LEFT && d2 == UP) {
		sign := utils.BoolToSign(d1 == LEFT)
		if ad1 < ad3 {
			p = utils.Point2{X: sign * ad1, Y: 0}
		} else if ad1 > ad3 {
			p = utils.Point2{X: sign * ad3, Y: 0}
		}
	} else {
		panic("Invalid U shape")
	}

	var reduced_area int = 0

	if ad1 < ad3 {
		ps[i3] = p3.Add(p)
		ps, _ = utils.ArrayRemoveIndices(ps, i1, i2)
		reduced_area = ad1*(ad2+1) + ad0*utils.BoolToInt(d0 != d2)
	} else if ad1 > ad3 {
		ps[i2] = p2.Add(p)
		ps, _ = utils.ArrayRemoveIndices(ps, i3, i4)
		reduced_area = ad3*(ad2+1) + ad4*utils.BoolToInt(d4 != d2)
	} else {
		ps, _ = utils.ArrayRemoveIndices(ps, i1, i2, i3, i4)
		reduced_area = ad1*(ad2+1) + ad0*utils.BoolToInt(d0 != d2) + ad4*utils.BoolToInt(d4 != d2)
	}

	if reduced_area < 0 {
		panic("Negative area")
	}

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
	dx := utils.AbsDiff(p1.X, p2.X) + 1
	dy := utils.AbsDiff(p1.Y, p2.Y) + 1
	return dx * dy
}

// func flipPointsLeftRight(ps []utils.Point2) []utils.Point2 {
// 	for i := 0; i < len(ps); i++ {
// 		ps[i].X = -ps[i].X
// 	}
// 	return ps
// }

// func flipPointsUpDown(ps []utils.Point2) []utils.Point2 {
// 	for i := 0; i < len(ps); i++ {
// 		ps[i].Y = -ps[i].Y
// 	}
// 	return ps
// }

// func flipPointsXY(ps []utils.Point2) []utils.Point2 {
// 	for i := 0; i < len(ps); i++ {
// 		new_x := ps[i].Y
// 		new_y := ps[i].X
// 		ps[i].X = new_x
// 		ps[i].Y = new_y
// 	}
// 	return ps
// }

// func rotatePoints90(ps []utils.Point2) []utils.Point2 {
// 	ps = flipPointsXY(ps)
// 	ps = flipPointsLeftRight(ps)
// 	return ps
// }

// func cycleShiftPoints(ps []utils.Point2, n int) []utils.Point2 {
// 	if n == 0 {
// 		return ps
// 	}
// 	if n < 0 {
// 		n = len(ps) + n
// 	}
// 	return append(ps[n:], ps[:n]...)
// }
