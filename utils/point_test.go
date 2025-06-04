package utils

import (
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestPointsIn2D(t *testing.T) {
	tests := []struct {
		width, height int
		expected      []Point2
	}{
		{
			width:  2,
			height: 2,
			expected: []Point2{
				{0, 0}, {1, 0},
				{0, 1}, {1, 1},
			},
		},
		{
			width:  3,
			height: 2,
			expected: []Point2{
				{0, 0}, {1, 0}, {2, 0},
				{0, 1}, {1, 1}, {2, 1},
			},
		},
		{
			width:  1,
			height: 1,
			expected: []Point2{
				{0, 0},
			},
		},
		{
			width:    0,
			height:   0,
			expected: []Point2{},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := PointsIn2D(tt.width, tt.height)
			assert.EqualArrays(t, result, tt.expected)
		})
	}
}
