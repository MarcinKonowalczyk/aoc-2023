package day11

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
	u := parseLinesToUniverse(lines)

	distances := findDistances(u, 2)

	sum_distances := utils.MapReduceValues(distances, 0, func(a, b int) int {
		return a + b
	})

	return sum_distances, nil
}

type Universe struct {
	galaxies   []utils.Point2
	empty_rows []int
	empty_cols []int
	width      int
	height     int
}

func (u Universe) HasGalaxyAt(p utils.Point2) bool {
	for _, g := range u.galaxies {
		if g == p {
			return true
		}
	}
	return false
}

func (u Universe) String() string {
	s := ""
	for _, p := range utils.PointsIn2D(u.width, u.height) {
		if p.X == 0 && p.Y > 0 {
			// Add a newline at the end of each row, but not at the beginning.
			s += "\n"
		}
		if u.HasGalaxyAt(p) {
			s += "#"
		} else {
			s += "."
		}
	}
	return s
}

func parseLinesToUniverse(lines []string) (u Universe) {
	u.height = len(lines)
	if u.height == 0 {
		return u
	}
	u.width = len(lines[0])
	if u.width == 0 {
		return u
	}

	for y, line := range lines {
		for x, c := range line {
			p := utils.Point2{X: x, Y: y}
			if c == '#' {
				u.galaxies = append(u.galaxies, p)
			}
		}
	}

	// Find empty rows and columns.
	galaxies_in_row := make([]bool, u.height)
	galaxies_in_col := make([]bool, u.width)
	for _, g := range u.galaxies {
		galaxies_in_row[g.Y] = true
		galaxies_in_col[g.X] = true
	}

	for y := 0; y < u.height; y++ {
		if !galaxies_in_row[y] {
			u.empty_rows = append(u.empty_rows, y)
		}
	}

	for x := 0; x < u.width; x++ {
		if !galaxies_in_col[x] {
			u.empty_cols = append(u.empty_cols, x)
		}
	}

	return u
}

func findManhattanDistance(g1, g2 utils.Point2, empty_rows, empty_cols []int, expansion_factor int) int {
	dx := find1DDistance(g1.X, g2.X, empty_cols, expansion_factor)
	dy := find1DDistance(g1.Y, g2.Y, empty_rows, expansion_factor)
	return dx + dy
}

func find1DDistance(v1, v2 int, empties []int, expansion_factor int) int {
	// Trivial case.
	if v1 == v2 {
		return 0
	}

	// Make sure v1 is the smaller one.
	if v1 > v2 {
		v1, v2 = v2, v1
	}

	// Normal distance.
	dv := v2 - v1

	if dv <= 0 {
		// Internal sanity check.
		panic("dv <= 0")
	}

	// Find which empty between v1 and v2.
	empties = utils.ArrayFilter(empties, func(v int) bool {
		return v > v1 && v < v2
	})

	// For each empty column we crossed, we've actually crossed three columns,
	// so need to add 2*len(empty_cols) to the distance.
	dv = dv + (expansion_factor-1)*len(empties)

	return dv
}

type GalaxyPair struct {
	g1, g2 utils.Point2
}

func NewGalaxyPair(g1, g2 utils.Point2) GalaxyPair {
	// Sort galaxies by x coordinate. If they are the same, sort them by y. If
	// they are still the same, the order doesn't matter.
	if g1.X > g2.X {
		g1, g2 = g2, g1
	} else if g1.X == g2.X {
		if g1.Y > g2.Y {
			g1, g2 = g2, g1
		}
	}
	return GalaxyPair{g1, g2}
}

// Find pairwise distances between all the galaxies.
func findDistances(u Universe, expansion_factor int) map[GalaxyPair]int {
	distances := make(map[GalaxyPair]int)
	for _, g1 := range u.galaxies {
		for _, g2 := range u.galaxies {

			// Don't calculate distance between a galaxy and itself.
			if g1 == g2 {
				continue
			}
			pair := NewGalaxyPair(g1, g2)

			distance := findManhattanDistance(g1, g2, u.empty_rows, u.empty_cols, expansion_factor)

			if _, ok := distances[pair]; ok {
				// We've already calculated this distance. Make sure it's the same.
				if distances[pair] != distance {
					panic("distance mismatch")
				}
			} else {
				distances[pair] = distance
			}
		}
	}
	return distances
}
