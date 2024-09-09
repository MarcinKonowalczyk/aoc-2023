package day03

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
)

func main_2(lines []string, verbose bool) (n int, err error) {
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

	messages := make(chan *gear)
	pg := padGrid(g)

	go yieldGears(pg, messages)

	sum := 0
	for {
		gear := <-messages

		if gear == nil {
			// we're done
			break
		}

		number1, err := chaseNumber(pg, gear.center.Add(gear.number1))
		if err != nil {
			return -1, err
		}
		number2, err := chaseNumber(pg, gear.center.Add(gear.number2))
		if err != nil {
			return -1, err
		}

		// fmt.Println("gear", gear)
		// fmt.Println("number1:", number1)
		// fmt.Println("number2:", number2)

		sum += number1 * number2

	}

	return sum, nil
}

type gear struct {
	// position of the gear
	center utils.Point2
	// position somewhere in the adjacent number 1
	number1 utils.Point2
	// position somewhere in the adjacent number 2
	number2 utils.Point2
}

func yieldGears(g grid, messages chan *gear) {

	N_rows, N_cols := gridToNRowsAndNCols(g)
	if N_rows == 0 || N_cols == 0 {
		panic("invalid grid")
	}

	for i := 0; i < N_rows; i++ {
		for j := 0; j < N_cols; j++ {
			if g[i][j] == '*' {
				n, err := getNeighborhood(g, i, j, j+1)
				if err != nil {
					panic(err)
				}
				numbers := numbersAround(n)
				if len(numbers) == 2 {
					// We found a gear. Report it with enough context to chase the numbers themselves.
					// fmt.Println("found gear at", i, j)
					// fmt.Println("numbers:", numbers)
					messages <- &gear{
						center:  utils.Point2{X: j, Y: i},
						number1: numbers[0],
						number2: numbers[1],
					}
				}
			}
		}
	}
	messages <- nil
}

// Return a list of points around the given point. Origin is at teh center of the
// grid.
func numbersAround(g grid) []utils.Point2 {
	N_rows, N_cols := gridToNRowsAndNCols(g)
	if N_rows != 3 || N_cols != 3 {
		panic("invalid grid")
	}
	numbers := make([]utils.Point2, 0)

	// First row
	points := pointsInARow(g[0])
	for _, p := range points {
		numbers = append(numbers, utils.Point2{X: p, Y: -1})
	}

	// Second row
	if isNumber(g[1][0]) {
		// Number to the left
		numbers = append(numbers, utils.Point2{X: -1, Y: 0})
	}
	if isNumber(g[1][2]) {
		// Number to the right
		numbers = append(numbers, utils.Point2{X: 1, Y: 0})
	}

	// Third row
	points = pointsInARow(g[2])
	for _, p := range points {
		numbers = append(numbers, utils.Point2{X: p, Y: 1})
	}

	return numbers
}

func pointsInARow(row []byte) []int {
	if len(row) != 3 {
		panic("invalid row")
	}
	numbers := make([]int, 0)
	r0, r1, r2 := isNumber(row[0]), isNumber(row[1]), isNumber(row[2])
	if r0 && !r1 && !r2 {
		// One number at the beginning
		numbers = append(numbers, -1)
	} else if !r0 && r1 && !r2 {
		// One number in the middle
		numbers = append(numbers, 0)
	} else if !r0 && !r1 && r2 {
		// One number at the end
		numbers = append(numbers, 1)
	} else if r0 && r1 && !r2 {
		// Two numbers at the beginning. Point at one of them.
		numbers = append(numbers, -1)
	} else if !r0 && r1 && r2 {
		// Two numbers at the end. Point at one of them.
		numbers = append(numbers, 1)
	} else if r0 && !r1 && r2 {
		// Two numbers in top corners. These are different numbers so point at
		// both.
		numbers = append(numbers, -1)
		numbers = append(numbers, 1)
	} else if r0 && r1 && r2 {
		// Three numbers in the first row. This is all one number so point
		// somewhere in the middle.
		numbers = append(numbers, 0)
	} else if !r0 && !r1 && !r2 {
		// No numbers in the first row.
	} else {
		panic("invalid row")
	}
	return numbers
}

func chaseNumber(g grid, somewhere utils.Point2) (int, error) {
	N_rows, N_cols := gridToNRowsAndNCols(g)
	if N_rows == 0 || N_cols == 0 {
		panic("invalid grid")
	}

	i, j := somewhere.X, somewhere.X

	// Chase the number to the left
	for i >= 0 {
		if isNumber(g[somewhere.Y][i]) {
			i--
		} else {
			// We went one step too far. Back out.
			i++
			break
		}
	}

	// Chase the number to the right
	for j < N_cols {
		if isNumber(g[somewhere.Y][j]) {
			j++
		} else {
			// NOTE: no need to back out here. That's just how the indexing turns out.
			break
		}
	}

	return strconv.Atoi(string(g[somewhere.Y][i:j]))
}
