package utils

import (
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestMapMap(t *testing.T) {
	m := map[int]int{
		1: 2,
		3: 4,
		5: 6,
	}
	e := map[int]int{
		1: 4,
		3: 8,
		5: 12,
	}
	result := MapMap(m, func(n int) int { return n * 2 })
	assert.EqualMaps(t, result, e)
}
func TestMapLookup(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	v, err := MapLookup(m, []int{2})
	assert.NoError(t, err)
	assert.EqualArrays(t, v, []string{"two"})

	v, err = MapLookup(m, []int{2, 3})
	assert.NoError(t, err)
	assert.EqualArrays(t, v, []string{"two", "three"})

	v, err = MapLookup(m, []int{3, 2})
	assert.NoError(t, err)
	assert.EqualArrays(t, v, []string{"three", "two"})

	v, err = MapLookup(m, []int{2, 3, 4})
	assert.NotEqual(t, err, nil)
	assert.EqualArrays(t, v, nil)
}
