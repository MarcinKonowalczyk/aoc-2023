package day03

import (
	"errors"
	"fmt"
	"strconv"
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
	N_rows := len(lines)
	if N_rows == 0 {
		return 0, nil
	}
	N_cols := len(lines[0])

	g := initGrid(N_cols, N_rows, ' ')

	for line_index, line := range lines {
		if len(line) != N_cols {
			return -1, fmt.Errorf("invalid line. Expected %d characted but got %d", N_cols, len(line))
		}
		g[line_index] = []byte(line)
	}

	messages := make(chan grid)

	go yieldNeighborhoods(g, messages)

	sum := 0
	for {
		neighborhood := <-messages

		if neighborhood == nil {
			// we're done
			break
		}

		// printGrid(&neighborhood)
		out, err := evaluateNeighborhood(neighborhood, '.')
		if err != nil {
			return -1, err
		}
		if out {
			n, err = numberFromNeighborhoodGrid(neighborhood)
			if err != nil {
				return -1, err
			}
			sum += n
		}

	}
	// fmt.Println("done")

	return sum, nil
}

type grid [][]byte

func initGrid(width int, height int, char byte) grid {
	new_gird := make(grid, height)
	for i := 0; i < height; i++ {
		new_gird[i] = make([]byte, width)
		for j := 0; j < width; j++ {
			new_gird[i][j] = char
		}
	}
	return new_gird
}

func printGrid(g grid) {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func yieldNeighborhoods(g grid, out chan grid) {
	pg := padGrid(g)
	for line_index, line := range pg {
		i := 0
		in_number := false
		for char_index, char := range line {
			if isNumber(char) {
				if !in_number {
					i = char_index
				}
				in_number = true
			} else {
				if in_number {
					neighborhood, err := getNeighborhood(pg, line_index, i, char_index)
					if err != nil {
						panic(err)
					}
					out <- neighborhood
				}
				in_number = false
			}
		}
		if in_number {
			neighborhood, err := getNeighborhood(pg, line_index, i, len(line))
			if err != nil {
				panic(err)
			}
			out <- neighborhood
		}
	}

	// we're done
	out <- nil
}

func gridToNRowsAndNCols(g grid) (int, int) {
	N_rows := len(g)
	if N_rows == 0 {
		return 0, 0
	}
	N_cols := len(g[0])
	return N_rows, N_cols
}

func padGrid(g grid) grid {
	N_rows, N_cols := gridToNRowsAndNCols(g)
	padded_grid := initGrid(N_cols+2, N_rows+2, '.')

	for i := 0; i < N_rows; i++ {
		for j := 0; j < N_cols; j++ {
			padded_grid[i+1][j+1] = g[i][j]
		}
	}

	return padded_grid
}

func getNeighborhood(g grid, row int, i1 int, i2 int) (grid, error) {
	N_rows, N_cols := gridToNRowsAndNCols(g)
	if N_rows == 0 {
		return nil, errors.New("grid is empty")
	}
	if N_cols == 0 {
		return nil, errors.New("grid is empty")
	}
	if row > N_rows-1 || row < 1 {
		return nil, fmt.Errorf("invalid row index %d", row)
	}
	if i1 > N_cols-1 || i1 < 1 {
		return nil, fmt.Errorf("invalid column index %d", i1)
	}
	if i2 > N_cols-1 || i2 < 1 {
		return nil, fmt.Errorf("invalid column index %d", i2)
	}
	if i1 > i2 {
		return nil, fmt.Errorf("invalid column index %d > %d", i1, i2)
	}

	neighborhood_width := i2 - i1 + 2
	neighborhood := initGrid(neighborhood_width, 3, ' ')
	for i := 0; i < neighborhood_width; i++ {
		for j := 0; j < 3; j++ {
			neighborhood[j][i] = g[row-1+j][i1-1+i]
		}
	}
	return neighborhood, nil
}

// Go through the edges of the grid. If we find a non-empty character, we have a
func evaluateNeighborhood(g grid, empty byte) (bool, error) {
	N_rows, N_cols := gridToNRowsAndNCols(g)
	if N_rows == 0 {
		return false, errors.New("grid is empty")
	}
	if N_cols == 0 {
		return false, errors.New("grid is empty")
	}

	// top and bottom edges
	for i := 0; i < N_cols; i++ {
		if g[0][i] != empty {
			return true, nil
		}
		if g[N_rows-1][i] != empty {
			return true, nil
		}
	}

	// left and right edges
	for i := 0; i < N_rows; i++ {
		if g[i][0] != empty {
			return true, nil
		}
		if g[i][N_cols-1] != empty {
			return true, nil
		}
	}

	return false, nil
}

func numberFromNeighborhoodGrid(g grid) (int, error) {
	N_rows, N_cols := gridToNRowsAndNCols(g)
	if N_rows != 3 {
		return 0, fmt.Errorf("invalid number of rows %d", N_rows)
	}
	if N_cols < 2 {
		return 0, fmt.Errorf("invalid number of columns %d", N_cols)
	}

	middle_row := string(g[1][1 : N_cols-1])
	return strconv.Atoi(middle_row)
}
