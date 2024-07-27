package dijkstra

import (
	"aoc2023/utils"
	"testing"
)

func TestDjkstra_SimpleGraph1(t *testing.T) {
	g := Graph{}
	vA := NewVertex()
	vB := NewVertex()
	vC := NewVertex()
	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddEdge(vA, vB, 1)
	g.AddEdge(vB, vC, 2)

	path, distande := ShortestPath(&g, vA, vC)
	utils.AssertEqual(t, distande, 3)
	utils.AssertEqualArrays(t, path, []*Vertex{vA, vB, vC})
}

func TestDjkstra_SimpleGraph2(t *testing.T) {

}
