package day17

import (
	"aoc2023/dijkstra"
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
	g, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	cell_mapping := createCellMapping(g)

	// fmt.Println("Cell mapping:", cell_mapping)
	gr := createGraph(cell_mapping, g)

	// fmt.Println("Graph:", gr)
	path, distance := dijkstra.ShortestPath(gr, cell_mapping[0][0][0], cell_mapping[g.rows-1][g.cols-1][0])

	index_path := pathToIndexPath(path, g.rows, g.cols)

	// fmt.Println("Path:", index_path)

	costs := costPath(index_path, g)

	sum_costs := utils.ArraySum(costs)

	// fmt.Println("Costs:", costs)
	fmt.Println("sum_costs:", sum_costs)
	if sum_costs != uint8(distance) {
		panic(fmt.Sprintf("sum_costs (%d) != distance (%d)", sum_costs, distance))
	}

	fmt.Println("Distance:", distance)

	// printGridWithPath(g, index_path)

	return 0, nil
}

// actual 859
// 861 is too high

type grid struct {
	cost [][]uint8
	rows int
	cols int
}

func parseLines(lines []string) (grid, error) {
	rows := len(lines)
	cols := len(lines[0])
	cost := make([][]uint8, rows)
	for i := range cost {
		cost[i] = make([]uint8, cols)
	}
	for i, line := range lines {
		for j, c := range line {
			p := c - '0'
			if p < 0 || p > 9 {
				return grid{}, fmt.Errorf("invalid character %c at %d,%d", c, i, j)
			}
			cost[i][j] = uint8(p)
		}
	}
	return grid{cost, rows, cols}, nil
}

func ijk_to_h(i, j, k, cols int) int {
	return i*(2*cols) + j + k*cols
}

func h_to_ijk(h, cols int) (int, int, int) {
	i := h / (2 * cols)
	j := h % cols
	k := (h / cols) % 2
	return i, j, k
}

// Hash each i,j,k to a unique number
func createCellMapping(g grid) [][][]dijkstra.Vertex {
	mapping := make([][][]dijkstra.Vertex, g.rows)
	all_vertices := make(map[int]struct{}) // to ensure uniqueness
	var h1, h2 int
	for i := 0; i < g.rows; i++ {
		mapping[i] = make([][]dijkstra.Vertex, g.cols)
		for j := 0; j < g.cols; j++ {
			mapping[i][j] = make([]dijkstra.Vertex, 2)

			h1 = ijk_to_h(i, j, 0, g.cols)
			ci, cj, ck := h_to_ijk(h1, g.cols)
			if ci != i || cj != j || ck != 0 {
				panic(fmt.Sprintf("h_to_ijk(%d) != (%d,%d,0)", h1, i, j))
			}

			if _, ok := all_vertices[h1]; ok {
				panic(fmt.Sprintf("duplicate vertex %d", h1))
			}
			all_vertices[h1] = struct{}{}

			h2 = ijk_to_h(i, j, 1, g.cols)
			ci, cj, ck = h_to_ijk(h2, g.cols)
			if ci != i || cj != j || ck != 1 {
				panic(fmt.Sprintf("h_to_ijk(%d) != (%d,%d,1)", h2, i, j))
			}

			if _, ok := all_vertices[h2]; ok {
				panic(fmt.Sprintf("duplicate vertex %d", h2))
			}
			all_vertices[h2] = struct{}{}

			mapping[i][j][0] = dijkstra.Vertex(h1)
			mapping[i][j][1] = dijkstra.Vertex(h2)
		}
	}

	all_vertices_arr := make([]int, 0, len(all_vertices))
	for h := range all_vertices {
		all_vertices_arr = append(all_vertices_arr, h)
	}

	// Sanity check
	min, _, err := utils.ArrayMin(all_vertices_arr)
	if err != nil {
		panic(err)
	}
	max, _, err := utils.ArrayMax(all_vertices_arr)
	if err != nil {
		panic(err)
	}

	if len(all_vertices) != g.rows*g.cols*2 {
		panic(fmt.Sprintf("expected %d vertices, got %d", g.rows*g.cols*2, len(all_vertices)))
	}

	if min != 0 {
		panic(fmt.Sprintf("expected min vertex 0, got %d", min))
	}

	if max != g.rows*g.cols*2-1 {
		panic(fmt.Sprintf("expected max vertex %d, got %d", g.rows*g.cols*2-1, max))
	}

	return mapping
}

func createGraph(cell_mapping [][][]dijkstra.Vertex, g grid) *dijkstra.Graph {
	gr := dijkstra.Graph{}
	// Add each cell as a vertex
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			for k := 0; k < 2; k++ {
				h := cell_mapping[i][j][k]
				gr.AddVertex(h)
			}
		}
	}

	// Add edges
	// Each cell from plane 0 can move left or right and to plane 1
	// Each cell from plane 1 can move up or down and to plane 0
	// The move right can be by one, two or three cells
	// The cost of the move is the sum of the values of the cells moved through

	var h1, h2 dijkstra.Vertex

	// Plane 0 to plane 1 (left or right)
	for i := 0; i < g.rows; i++ { // for each row
		for j := 0; j < g.cols; j++ { // for each column
			// fmt.Println("Adding 0->1 edges for", i, j)
			for d := 1; d <= 3; d++ { // for each possible move distance
				// Right move
				if j+d < g.cols {
					cost := 0
					for l := 1; l <= d; l++ {
						cost += int(g.cost[i][j+l])
					}
					h1 := cell_mapping[i][j][0]
					h2 := cell_mapping[i][j+d][1]
					// fmt.Printf("Adding edge %d -> %d with cost %d\n", h1, h2, cost)
					gr.AddDirectedEdge(h1, h2, cost)
				}
				// Left move
				if j-d >= 0 {
					cost := 0
					for l := 1; l <= d; l++ {
						cost += int(g.cost[i][j-l])
					}
					h1 = cell_mapping[i][j][0]
					h2 = cell_mapping[i][j-d][1]
					// fmt.Printf("Adding edge %d -> %d with cost %d\n", h1, h2, cost)
					gr.AddDirectedEdge(h1, h2, cost)
				}
			}
		}
	}

	// Plane 1 to plane 0 (up or down)
	for i := 0; i < g.rows; i++ { // for each row
		for j := 0; j < g.cols; j++ { // for each column
			// fmt.Println("Adding 1->0 edges for", i, j)
			for d := 1; d <= 3; d++ { // for each possible move distance
				// Down move
				if i+d < g.rows {
					cost := 0
					for l := 1; l <= d; l++ {
						cost += int(g.cost[i+l][j])
					}
					h1 = cell_mapping[i][j][1]
					h2 = cell_mapping[i+d][j][0]
					// fmt.Printf("Adding edge %d -> %d with cost %d\n", h1, h2, cost)
					gr.AddDirectedEdge(h1, h2, cost)
				}
				// Up move
				if i-d >= 0 {
					cost := 0
					for l := 1; l <= d; l++ {
						cost += int(g.cost[i-l][j])
					}
					h1 = cell_mapping[i][j][1]
					h2 = cell_mapping[i-d][j][0]
					// fmt.Printf("Adding edge %d -> %d with cost %d\n", h1, h2, cost)
					gr.AddDirectedEdge(h1, h2, cost)
				}
			}
		}
	}

	// We can also move between the two planes for free, but only in the top left and bottom right corners
	// This represents the choice of starting in plane 0 or plane 1
	// and the choice of ending in plane 0 or plane 1
	gr.AddUndirectedEdge(cell_mapping[0][0][0], cell_mapping[0][0][1], 0)
	gr.AddUndirectedEdge(cell_mapping[g.rows-1][g.cols-1][1], cell_mapping[g.rows-1][g.cols-1][0], 0)

	return &gr
}

func pathToIndexPath(path []dijkstra.Vertex, rows, cols int) []utils.Point {
	// Check that all the moves are correct
	// Except for the top left and bottom right corners, all moves should be to the right or down
	// Plane 0 should only move to plane 1 and right
	// Plane 1 should only move to plane 0 and down
	index_path := make([]utils.Point, 0)

	for i := 0; i < len(path)-1; i++ {
		// fmt.Printf("Checking move %d (h%d -> h%d)\n", i, path[i], path[i+1])
		ii1, jj1, kk1 := h_to_ijk(int(path[i]), cols)
		ii2, jj2, kk2 := h_to_ijk(int(path[i+1]), cols)
		// fmt.Printf("Step %d: %d,%d (%d) -> %d,%d (%d)\n", i, jj1, ii1, kk1, jj2, ii2, kk2)
		if ii1 == 0 && jj1 == 0 && ii2 == 0 && jj2 == 0 {
			// Skip the top left corner
			continue
		}
		if ii1 == rows-1 && jj1 == cols-1 && ii2 == rows-1 && jj2 == cols-1 {
			// Skip the bottom right corner
			continue
		}
		if kk1 == 0 && kk2 == 1 {
			// Were on plane 0 and moving to plane 1
			diff := jj2 - jj1
			if ii1 != ii2 || diff == 0 || diff > 3 || diff < -3 {
				panic(fmt.Sprintf("Invalid 0->1 move from %d,%d to %d,%d but should be right", jj1, ii1, jj2, ii2))
			}
		} else if kk1 == 1 && kk2 == 0 {
			// Were on plane 1 and moving to plane 0
			diff := ii2 - ii1
			if jj1 != jj2 || diff == 0 || diff > 3 || diff < -3 {
				panic(fmt.Sprintf("Invalid 1->0 move from %d,%d to %d,%d but should be down", jj1, ii1, jj2, ii2))
			}
		} else {
			panic(fmt.Sprintf("Invalid move from plane %d to plane %d (should be 0->1 or 1->0)", kk1, kk2))
		}

		index_path = append(index_path, utils.Point{X: jj2, Y: ii2})
	}

	// Add the last point
	ii, jj, _ := h_to_ijk(int(path[len(path)-1]), cols)
	index_path = append(index_path, utils.Point{X: jj, Y: ii})

	// Deduplicate first point
	if len(index_path) > 1 && index_path[0].X == index_path[1].X && index_path[0].Y == index_path[1].Y {
		index_path = index_path[1:]
	}

	// Deduplicate last point
	if len(index_path) > 1 && index_path[len(index_path)-1].X == index_path[len(index_path)-2].X && index_path[len(index_path)-1].Y == index_path[len(index_path)-2].Y {
		index_path = index_path[:len(index_path)-1]
	}

	// Add {0,0} if the first point is not {0,0}
	if index_path[0].X != 0 || index_path[0].Y != 0 {
		index_path = append([]utils.Point{{X: 0, Y: 0}}, index_path...)
	}

	return index_path
}

// Get the list of costs along the path
func costPath(index_path []utils.Point, g grid) []uint8 {
	costs := make([]uint8, 0)
	// Walk from point to point and sum the costs
	for i := 0; i < len(index_path)-1; i++ {
		p1 := index_path[i]
		p2 := index_path[i+1]
		if p1.X == p2.X {
			// Move left or right
			delta := p2.Y - p1.Y
			abs_delta := delta
			sign_delta := 1
			if delta < 0 {
				abs_delta = -delta
				sign_delta = -1
			}
			for j := 1; j < abs_delta+1; j++ {
				costs = append(costs, g.cost[p1.Y+sign_delta*j][p1.X])
			}
		} else if p1.Y == p2.Y {
			// Move up or down
			delta := p2.X - p1.X
			abs_delta := delta
			sign_delta := 1
			if delta < 0 {
				abs_delta = -delta
				sign_delta = -1
			}
			for j := 1; j < abs_delta+1; j++ {
				costs = append(costs, g.cost[p1.Y][p1.X+sign_delta*j])
			}
		} else {
			panic(fmt.Sprintf("Invalid move from %v to %v", p1, p2))
		}
	}

	return costs
}

func printGridWithPath(g grid, index_path []utils.Point) {
	// Print the grid, with the path marked by X
	to_print := make([][]byte, g.rows)
	for i := 0; i < g.rows; i++ {
		to_print[i] = make([]byte, g.cols)
		for j := 0; j < g.cols; j++ {
			to_print[i][j] = byte(g.cost[i][j] + '0')
		}
	}

	for i := 0; i < len(index_path)-1; i++ {
		p1 := index_path[i]
		p2 := index_path[i+1]
		if p1.X == p2.X {
			// Move up or down
			delta := p2.Y - p1.Y
			abs_delta := delta
			sign_delta := 1
			s := byte('v')
			if delta < 0 {
				abs_delta = -delta
				sign_delta = -1
				s = byte('^')
			}
			for j := 1; j < abs_delta+1; j++ {
				to_print[p1.Y+sign_delta*j][p1.X] = s
			}
		} else if p1.Y == p2.Y {
			// Move left or right
			delta := p2.X - p1.X
			abs_delta := delta
			sign_delta := 1
			s := byte('>')
			if delta < 0 {
				abs_delta = -delta
				sign_delta = -1
				s = byte('<')
			}
			for j := 1; j < abs_delta+1; j++ {
				to_print[p1.Y][p1.X+sign_delta*j] = s
			}

		} else {
			panic(fmt.Sprintf("Invalid move from %v to %v", p1, p2))
		}
	}

	for i := 0; i < g.rows; i++ {
		fmt.Println(string(to_print[i]))
	}
}
