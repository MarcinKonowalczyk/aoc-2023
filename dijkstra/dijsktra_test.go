package dijkstra

import (
	"aoc2023/utils"
	"testing"
)

func TestNodeQueue_Enqueue(t *testing.T) {

	q := NodeQueue{}

	q.Enqueue(Vertex{Distance: 1})
	q.Enqueue(Vertex{Distance: 7})
	q.Enqueue(Vertex{Distance: 3})
	q.Enqueue(Vertex{Distance: 2})
	q.Enqueue(Vertex{Distance: 5})
	q.Enqueue(Vertex{Distance: 10})
	q.Enqueue(Vertex{Distance: 4})
	q.Enqueue(Vertex{Distance: 6})
	q.Enqueue(Vertex{Distance: 9})
	q.Enqueue(Vertex{Distance: 8})

	// Check that the queue is sorted
	utils.AssertEqualArrays(t, q.Items, []Vertex{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}})
}
