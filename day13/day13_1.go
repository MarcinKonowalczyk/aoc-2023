package day13

import (
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
	maps, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	summary := 0

	for i, m := range maps {
		vfolds, hfolds, err := tryAllFolds(m)
		if err != nil {
			return -1, err
		}

		has_vfolds := len(vfolds) > 0
		has_hfolds := len(hfolds) > 0

		if !has_vfolds && !has_hfolds {
			return -1, fmt.Errorf("no fold found for map %d", i)
		} else if has_vfolds && has_hfolds {
			return -1, fmt.Errorf("both vertical and horizontal folds found for map %d", i)
		}

		if has_vfolds {
			if len(vfolds) != 1 {
				return -1, fmt.Errorf("multiple vertical folds found for map %d", i)
			}
			vfold := vfolds[0]
			// fmt.Println("Vertical fold found for map", i+1, "between columns", vfold+1, "and", vfold+2)
			summary += vfold + 1
		}

		if has_hfolds {
			if len(hfolds) != 1 {
				return -1, fmt.Errorf("multiple horizontal folds found for map %d", i)
			}
			hfold := hfolds[0]
			// fmt.Println("Horizontal fold found for map", i+1, "between rows", hfold+1, "and", hfold+2)
			summary += 100 * (hfold + 1)
		}

	}

	return summary, nil
}

type Map [][]bool

func (m Map) nRows() int {
	return len(m)
}

func (m Map) nCols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Map) String() string {
	s := "Map:\n"
	for i, row := range m {
		for _, cell := range row {
			if cell {
				s += "#"
			} else {
				s += "."
			}
		}
		if i < len(m)-1 {
			s += "\n"
		}
	}
	return s
}

func parseLines(lines []string) ([]Map, error) {
	var maps []Map = make([]Map, 0)
	var m Map = make(Map, 0)
	for _, line := range lines {
		if line == "" {
			if m.nRows() > 0 {
				maps = append(maps, m)
				m = make(Map, 0)
			}
		} else {
			row := make([]bool, 0)
			for _, c := range line {
				if c == '.' {
					row = append(row, false)
				} else if c == '#' {
					row = append(row, true)
				} else {
					return nil, fmt.Errorf("invalid character: %c", c)
				}
			}
			m = append(m, row)
		}
	}
	if m.nRows() > 0 {
		maps = append(maps, m)
	}
	return maps, nil
}

// Fold about a vertical axis between i and i+1 columns
func vFold(m Map, i int) (Map, error) {
	if i < 0 {
		return nil, fmt.Errorf("invalid column index: %d", i)
	}
	if i+1 >= m.nCols() {
		return nil, fmt.Errorf("invalid column index: %d", i+1)
	}

	dx1 := i + 1
	dx2 := m.nCols() - i - 1
	n_cols := min(dx1, dx2)
	out := make(Map, m.nRows())
	for r := 0; r < m.nRows(); r++ {
		row := make([]bool, n_cols)
		for c := 0; c < n_cols; c++ {
			right_c := i + 1 + c
			left_c := i - c
			left_val := m[r][left_c]
			right_val := m[r][right_c]
			if left_val == right_val {
				row[c] = true
			}
		}
		out[r] = row
	}
	return out, nil
}

func (m Map) allTrue() bool {
	for _, row := range m {
		for _, cell := range row {
			if !cell {
				return false
			}
		}
	}
	return true
}

func (m Map) transpose() Map {
	out := make(Map, m.nCols())
	for c := 0; c < m.nCols(); c++ {
		row := make([]bool, m.nRows())
		for r := 0; r < m.nRows(); r++ {
			row[r] = m[r][c]
		}
		out[c] = row
	}
	return out
}

// Fold about a horizontal axis between i and i+1 rows
// This is the same as folding about a vertical axis between i and i+1 columns
func hFold(m Map, i int) (Map, error) {
	transposed := m.transpose()
	folded, err := vFold(transposed, i)
	if err != nil {
		return nil, err
	}
	return folded.transpose(), nil
}

func tryVerticalFolds(m Map) ([]int, error) {
	folds := make([]int, 0)
	for j := 0; j < m.nCols()-1; j++ {
		folded, err := vFold(m, j)
		if err != nil {
			return []int{}, err
		}
		if folded.allTrue() {
			folds = append(folds, j)
		}
	}
	return folds, nil
}

func tryHorizontalFolds(m Map) ([]int, error) {
	folds := make([]int, 0)
	for j := 0; j < m.nRows()-1; j++ {
		folded, err := hFold(m, j)
		if err != nil {
			return []int{}, err
		}
		if folded.allTrue() {
			folds = append(folds, j)
		}
	}
	return folds, nil
}

func tryAllFolds(m Map) (vfold, hfold []int, err error) {
	v_folds, err := tryVerticalFolds(m)
	if err != nil {
		return []int{}, []int{}, err
	}
	h_folds, err := tryHorizontalFolds(m)
	if err != nil {
		return []int{}, []int{}, err
	}
	return v_folds, h_folds, nil
}
