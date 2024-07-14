package day14

import "fmt"

const N = 1000000000

func main_2(lines []string) (n int, err error) {
	grid := parseLines(lines)
	spinCycle(&grid)
	fmt.Println("After 1 spin cycle:")
	fmt.Println(grid)
	spinCycle(&grid)
	fmt.Println("After 2 spin cycles:")
	fmt.Println(grid)
	spinCycle(&grid)
	fmt.Println("After 3 spin cycles:")
	fmt.Println(grid)

	// for i := 0; i < N; i++ {
	// 	if i%(N/10000) == 0 {
	// 		done_percent := float64(i) / N
	// 		fmt.Printf("%.4f%%\n", done_percent)
	// 	}
	// 	spinCycle(&grid)
	// }
	weight := calcNorthWeight(&grid)
	return weight, nil
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func rotateClockwise(g *Grid) {
	n_rows := len(*g)
	n_cols := len((*g)[0])
	var temp Rock // for swapping
	for i := 0; i < n_rows/2; i++ {
		for j := 0; j < n_cols/2; j++ {
			ii := n_rows - i - 1 // i from the bottom
			jj := n_cols - j - 1 // j from the right
			temp = (*g)[i][j]
			(*g)[i][j] = (*g)[jj][i]
			(*g)[jj][i] = (*g)[ii][jj]
			(*g)[ii][jj] = (*g)[j][ii]
			(*g)[j][ii] = temp
		}
	}
}

func spinCycle(g *Grid) {
	slideNorth(g)      // actual north
	rotateClockwise(g) // north is now east, so slideNorth is slideWest
	slideNorth(g)      // west
	rotateClockwise(g) // ...
	slideNorth(g)      // south
	rotateClockwise(g)
	slideNorth(g)      // east
	rotateClockwise(g) // back to the original orientation
}
