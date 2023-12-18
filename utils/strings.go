package utils

import (
	"strconv"
	"strings"
)

func StringOfNumbersToInts(s string) ([]int, error) {
	s = strings.Trim(s, " ")
	parts := strings.Split(s, " ")
	parts = ArrayFilter(parts, func(s string) bool {
		return s != ""
	})
	var result []int
	if len(parts) > 0 {
		result = make([]int, len(parts))
		for i, part := range parts {
			r, err := strconv.Atoi(part)
			if err != nil {
				return result, err
			}
			result[i] = r
		}
	}
	return result, nil
}
