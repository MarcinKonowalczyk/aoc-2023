package day17

import (
	"aoc2023/dijkstra"
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

	gr := createGraph(cell_mapping, g)

	path, distance := dijkstra.ShortestPath(&gr, cell_mapping[0][0], cell_mapping[g.rows-1][g.cols-1])

	fmt.Println("Shortest path:", path)
	fmt.Println("Distance:", distance)

	return 0, nil
}

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

// Hash each i,j to a unique number
func createCellMapping(g grid) [][]dijkstra.Vertex {
	mapping := make([][]dijkstra.Vertex, g.rows)
	all_vertices := make(map[dijkstra.Vertex]struct{}) // to ensure uniqueness
	for i := 0; i < g.rows; i++ {
		mapping[i] = make([]dijkstra.Vertex, g.cols)
		for j := 0; j < g.cols; j++ {
			new_vertex := dijkstra.Vertex(i*g.cols + j)
			if _, ok := all_vertices[new_vertex]; ok {
				panic("duplicate vertex")
			}
			mapping[i][j] = new_vertex
		}
	}
	return mapping
}

func createGraph(cell_mapping [][]dijkstra.Vertex, g grid) dijkstra.Graph {
	gr := dijkstra.Graph{}
	// Add each cell as a vertex
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			gr.AddVertex(cell_mapping[i][j])
		}
	}
	// Add edges. Each cell can move up, down, left, right by up to 3 cells (1, 2, or 3)
	// The cost of the move is the sum of the values of the cells moved through
	for i := 0; i < g.rows; i++ { // for each row
		for j := 0; j < g.cols; j++ { // for each column
			for k := 1; k <= 3; k++ { // for each possible move distance
				if i+k < g.rows {
					cost := 0
					for l := 1; l <= k; l++ {
						cost += int(g.cost[i+l][j])
					}
					gr.AddEdge(cell_mapping[i][j], cell_mapping[i+k][j], cost)
				}
				if i-k >= 0 {
					cost := 0
					for l := 1; l <= k; l++ {
						cost += int(g.cost[i-l][j])
					}
					gr.AddEdge(cell_mapping[i][j], cell_mapping[i-k][j], cost)
				}
				if j+k < g.cols {
					cost := 0
					for l := 1; l <= k; l++ {
						cost += int(g.cost[i][j+l])
					}
					gr.AddEdge(cell_mapping[i][j], cell_mapping[i][j+k], cost)
				}
				if j-k >= 0 {
					cost := 0
					for l := 1; l <= k; l++ {
						cost += int(g.cost[i][j-l])
					}
					gr.AddEdge(cell_mapping[i][j], cell_mapping[i][j-k], cost)
				}
			}
		}
	}

	return gr
}
