package day13

import "fmt"

func main_2(lines []string) (n int, err error) {
	maps, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	summary := 0

	for i, m := range maps {
		vfold, hfold, err := tryAllFolds(m)
		if err != nil {
			return -1, err
		}

		if vfold == -1 && hfold == -1 {
			return -1, fmt.Errorf("no fold found for map %d", i)
		} else if vfold >= 0 && hfold >= 0 {
			return -1, fmt.Errorf("both vertical and horizontal folds found for map %d", i)
		}

		if vfold >= 0 {
			// fmt.Println("Vertical fold found for map", i+1, "between columns", vfold+1, "and", vfold+2)
			summary += vfold + 1
		}

		if hfold >= 0 {
			// fmt.Println("Horizontal fold found for map", i+1, "between rows", hfold+1, "and", hfold+2)
			summary += 100 * (hfold + 1)
		}

	}

	return summary, nil
}
