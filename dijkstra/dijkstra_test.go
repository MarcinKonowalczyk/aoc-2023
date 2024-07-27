package dijkstra

import (
	"aoc2023/utils"
	"testing"
)

func TestDjkstra_ABC_Line(t *testing.T) {
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

func TestDjkstra_ABC_Cycle(t *testing.T) {
	g := Graph{}
	vA := NewVertex()
	vB := NewVertex()
	vC := NewVertex()
	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddEdge(vA, vB, 1)
	g.AddEdge(vA, vC, 2)
	g.AddEdge(vB, vC, 3)

	path, distande := ShortestPath(&g, vA, vC)
	utils.AssertEqual(t, distande, 2)
	utils.AssertEqualArrays(t, path, []*Vertex{vA, vC})
}

func TestDjkstra_LargerCase(t *testing.T) {
	// This is the graph from the wikipedia page on Dijkstra's algorithm
	g := Graph{}
	vA := NewVertex()
	vB := NewVertex()
	vC := NewVertex()
	vD := NewVertex()
	vE := NewVertex()
	vF := NewVertex()

	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddVertex(vD)
	g.AddVertex(vE)
	g.AddVertex(vF)

	g.AddEdge(vA, vB, 7)
	g.AddEdge(vA, vC, 9)
	g.AddEdge(vA, vD, 14)
	g.AddEdge(vB, vC, 10)
	g.AddEdge(vB, vF, 15)
	g.AddEdge(vC, vD, 2)
	g.AddEdge(vC, vF, 11)
	g.AddEdge(vD, vE, 9)
	g.AddEdge(vE, vF, 6)

	path, distande := ShortestPath(&g, vA, vE)
	utils.AssertEqual(t, distande, 20)
	utils.AssertEqualArrays(t, path, []*Vertex{vA, vC, vD, vE})
}
