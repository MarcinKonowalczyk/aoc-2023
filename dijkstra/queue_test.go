package dijkstra

import (
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestPriorityQueue(t *testing.T) {

	q := PriorityQueue[Vertex]{}

	v1 := Vertex(1)
	v2 := Vertex(2)
	v3 := Vertex(3)
	v4 := Vertex(4)
	v5 := Vertex(5)
	v6 := Vertex(6)
	v7 := Vertex(7)
	v8 := Vertex(8)
	v9 := Vertex(9)
	v10 := Vertex(10)

	enqueue_order := []Vertex{v1, v7, v3, v2, v5, v10, v4, v6, v9, v8}
	enqueue_distances := []int{1, 7, 3, 2, 5, 10, 4, 6, 9, 8}
	expected_order := []Vertex{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}

	for i := 0; i < len(enqueue_order); i++ {
		q.Enqueue(enqueue_order[i], enqueue_distances[i])
	}
	// Check that the queue is sorted
	assert.EqualArrays(t, q.Items(), expected_order)
}
