package day12

import (
	"fmt"
	"testing"
)

func TestStepFromLeft(t *testing.T) {
	// [?#?,?#,?] 5,1
	blocks := [][]Spring{
		{UNKNOWN, DAMAGED, UNKNOWN},
		{UNKNOWN, DAMAGED},
		{UNKNOWN},
	}
	groups := []int{5, 1}

	line := Line{"", blocks, groups}
	ll, g, end := stepFromLeft(line)

	// expected_blocks := [][]Spring{
	// 	{DAMAGED, DAMAGED, UNKNOWN},
	fmt.Println(ll)
	fmt.Println(g)
	fmt.Println(end)

}
