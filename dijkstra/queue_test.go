package dijkstra

import (
	"aoc2023/utils"
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	q := PriorityQueue[*Vertex]{}

	v1 := &Vertex{}
	v2 := &Vertex{}
	v3 := &Vertex{}
	v4 := &Vertex{}
	v5 := &Vertex{}
	v6 := &Vertex{}
	v7 := &Vertex{}
	v8 := &Vertex{}
	v9 := &Vertex{}
	v10 := &Vertex{}

	enqueue_order := []*Vertex{v1, v7, v3, v2, v5, v10, v4, v6, v9, v8}
	enqueue_distances := []int{1, 7, 3, 2, 5, 10, 4, 6, 9, 8}
	expected_order := []*Vertex{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}

	for i := 0; i < len(enqueue_order); i++ {
		q.Enqueue(enqueue_order[i], enqueue_distances[i])
	}
	// Check that the queue is sorted
	utils.AssertEqualArrays(t, q.Items(), expected_order)
}
