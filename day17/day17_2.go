package day17

import (
	"aoc2023/dijkstra"
	"aoc2023/utils"
	"fmt"
)

func main_2(lines []string) (n int, err error) {
	g, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	cell_mapping := createCellMapping(g)

	gr, err := createGraph(cell_mapping, g, 4, 10)
	if err != nil {
		return -1, err
	}

	path, distance := dijkstra.ShortestPath(gr, cell_mapping[0][0][0], cell_mapping[g.rows-1][g.cols-1][0])

	checkPath(path, g.rows, g.cols, 4, 10)

	index_path := pathToIndexPath(path, g.rows, g.cols)

	costs := costPath(index_path, g)

	sum_costs := utils.ArraySum(utils.ArrayMap(costs, func(c uint8) int { return int(c) }))

	fmt.Println("sum_costs:", sum_costs)
	if sum_costs != distance {
		panic(fmt.Sprintf("sum_costs (%d) != distance (%d)", sum_costs, distance))
	}

	fmt.Println("Distance:", distance)

	// printGridWithPath(g, index_path)

	return 0, nil
}
