package day12

import (
	"aoc2023/utils"
	"testing"
)

func TestLineHashDifferentSprings(t *testing.T) {
	blocks := [][]Spring{
		MustSpringsFromString("##?##?##?"),
	}
	groups := []uint8{2}
	l1 := Line{blocks, groups}
	h1 := l1.Hash()

	blocks = [][]Spring{
		MustSpringsFromString("##?##?##?#"),
	}
	groups = []uint8{2}
	l2 := Line{blocks, groups}
	h2 := l2.Hash()

	utils.AssertNotEqual(t, h1, h2)
}

func TestLineHashDifferentGroups(t *testing.T) {
	blocks := [][]Spring{
		MustSpringsFromString("##?##?##?"),
	}
	groups := []uint8{2}
	l1 := Line{blocks, groups}
	h1 := l1.Hash()

	blocks = [][]Spring{
		MustSpringsFromString("##?##?##?"),
	}
	groups = []uint8{3}
	l2 := Line{blocks, groups}
	h2 := l2.Hash()

	utils.AssertNotEqual(t, h1, h2)
}

func TestLineHashSpecificCase01(t *testing.T) {
	// [?,###] 1,3 -> 9447435613075004898
	// [###] 1,3 -> 9447435613075004898
	//
	h1 := Line{
		blocks: [][]Spring{
			MustSpringsFromString("?"),
		},
	}.Hash()

	h2 := Line{
		blocks: [][]Spring{},
	}.Hash()

	utils.AssertNotEqual(t, h1, h2)
}

func TestLineHashSpecificCase02(t *testing.T) {
	// [?,?]
	// [??]
	//
	h1 := Line{
		blocks: [][]Spring{
			MustSpringsFromString("?"),
			MustSpringsFromString("?"),
		},
	}.Hash()

	h2 := Line{
		blocks: [][]Spring{
			MustSpringsFromString("??"),
		},
	}.Hash()

	utils.AssertNotEqual(t, h1, h2)
}

func TestLineHashSpecificCase03(t *testing.T) {
	// [#,?] 3 -> 6104263669925314871
	// [##?] 3 -> 6104263669925314871

	h1 := Line{
		blocks: [][]Spring{
			MustSpringsFromString("#"),
			MustSpringsFromString("?"),
		},
	}.Hash()

	h2 := Line{
		blocks: [][]Spring{
			MustSpringsFromString("##?"),
		},
	}.Hash()

	utils.AssertNotEqual(t, h1, h2)
}
