package day13

import (
	"aoc2023/utils"
	"fmt"
)

func main_2(lines []string) (n int, err error) {
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

		var vfold int = -1
		if has_vfolds {
			if len(vfolds) != 1 {
				return -1, fmt.Errorf("multiple vertical folds found for map %d", i)
			}
			vfold = vfolds[0]
		}

		var hfold int = -1
		if has_hfolds {
			if len(hfolds) != 1 {
				return -1, fmt.Errorf("multiple horizontal folds found for map %d", i)
			}
			hfold = hfolds[0]
		}

		var c chan Map = make(chan Map)
		var s chan bool = make(chan bool)
		var I int = -1
		go yieldAllAlterations(m, c, s)

		for m2 := range c {
			if m2 == nil {
				// We've run out of maps to try
				return -1, fmt.Errorf("No smudge found for map index %d", i)
			}

			vfolds2, hfolds2, err := tryAllFolds(m2)

			if err != nil {
				return -1, err
			}

			if has_vfolds {
				// Original map had a vertical fold. Remove it from the list of folds in the new map if it's there
				I = utils.ArrayIndexOf(vfolds2, vfold)
				if I >= 0 {
					vfolds2 = utils.ArrayRemoveIndex(vfolds2, I)
				}
			} else if has_hfolds {
				// Original map had a horizontal fold. Remove it from the list of folds in the new map if it's there
				I = utils.ArrayIndexOf(hfolds2, hfold)
				if I >= 0 {
					hfolds2 = utils.ArrayRemoveIndex(hfolds2, I)
				}
			}

			if len(vfolds2) > 0 {
				// We have some new vertical folds!
				vfold2 := vfolds2[0]
				summary += 1 + vfold2
				break
			}

			if len(hfolds2) > 0 {
				// We have some new horizontal folds!
				hfold2 := hfolds2[0]
				summary += 100 * (hfold2 + 1)
				break
			}

		}
		s <- true
	}

	return summary, nil
}

func (m Map) copy() Map {
	m2 := make(Map, m.nRows())
	for i := 0; i < m.nRows(); i++ {
		m2[i] = make([]bool, m.nCols())
		copy(m2[i], m[i])
	}
	return m2
}

// yield all possible single bit flips of the map
func yieldAllAlterations(m Map, c chan Map, s chan bool) {
	defer close(c)

	n_rows := m.nRows()
	n_cols := m.nCols()
	for i := 0; i < n_rows; i++ {
		for j := 0; j < n_cols; j++ {
			select {
			case <-s:
				return
			case c <- func() Map {
				m2 := m.copy()
				m2[i][j] = !m2[i][j]
				return m2
			}():
			}
		}
	}
	c <- nil // signal end
}
