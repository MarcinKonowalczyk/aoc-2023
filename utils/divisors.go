package utils

import (
	"math"
	"sort"
)

// Return unordered divisors of n.
func DivisorsUnordered(n int) map[int]struct{} {
	divisors := make(map[int]struct{})
	if n <= 0 {
		return divisors
	}
	N := int(math.Sqrt(float64(n)))
	for i := 1; i <= N; i++ {
		if n%i == 0 {
			divisors[i] = struct{}{}
			if i != n/i {
				divisors[n/i] = struct{}{}
			}
		}
	}
	return divisors
}

// Return the divisors of n in increasing order.
func Divisors(n int) []int {
	divisors_map := DivisorsUnordered(n)
	divisors := make([]int, 0, len(divisors_map))
	for k := range divisors_map {
		divisors = append(divisors, k)
	}
	sort.Ints(divisors)
	return divisors
}
