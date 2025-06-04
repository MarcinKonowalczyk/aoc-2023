package dijkstra

import (
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestDjkstra_ABC_Line(t *testing.T) {
	g := Graph{}
	vA := Vertex(1)
	vB := Vertex(2)
	vC := Vertex(3)
	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddUndirectedEdge(vA, vB, 1)
	g.AddUndirectedEdge(vB, vC, 2)

	path, distande := ShortestPath(&g, vA, vC)
	assert.Equal(t, distande, 3)
	assert.EqualArrays(t, path, []Vertex{vA, vB, vC})
}

func TestDjkstra_ABC_Cycle(t *testing.T) {
	g := Graph{}
	vA := Vertex(1)
	vB := Vertex(2)
	vC := Vertex(3)
	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddUndirectedEdge(vA, vB, 1)
	g.AddUndirectedEdge(vA, vC, 2)
	g.AddUndirectedEdge(vB, vC, 3)

	path, distande := ShortestPath(&g, vA, vC)
	assert.Equal(t, distande, 2)
	assert.EqualArrays(t, path, []Vertex{vA, vC})
}

func TestDjkstra_LargerCase(t *testing.T) {
	// This is the graph from the wikipedia page on Dijkstra's algorithm
	g := Graph{}
	vA := Vertex(1)
	vB := Vertex(2)
	vC := Vertex(3)
	vD := Vertex(4)
	vE := Vertex(5)
	vF := Vertex(6)

	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddVertex(vD)
	g.AddVertex(vE)
	g.AddVertex(vF)

	g.AddUndirectedEdge(vA, vB, 7)
	g.AddUndirectedEdge(vA, vC, 9)
	g.AddUndirectedEdge(vA, vD, 14)
	g.AddUndirectedEdge(vB, vC, 10)
	g.AddUndirectedEdge(vB, vF, 15)
	g.AddUndirectedEdge(vC, vD, 2)
	g.AddUndirectedEdge(vC, vF, 11)
	g.AddUndirectedEdge(vD, vE, 9)
	g.AddUndirectedEdge(vE, vF, 6)

	path, distande := ShortestPath(&g, vA, vE)
	assert.Equal(t, distande, 20)
	assert.EqualArrays(t, path, []Vertex{vA, vC, vD, vE})
}
