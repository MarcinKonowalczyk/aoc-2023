package dijkstra

import (
	"aoc2023/utils"
	"testing"
)

// Comparator for edge maps
func compareEdgeMaps(a, b map[Vertex][]*Edge) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			return false
		}
		if len(va) != len(vb) {
			return false
		}
		for i, ea := range va {
			eb := vb[i]
			if ea.weight != eb.weight {
				return false
			}
			if ea.from != eb.from {
				return false
			}
		}
	}
	return true
}

func assertEqualEdgeMaps(t *testing.T, a, b map[Vertex][]*Edge) {
	utils.AssertEqualWithComparator(t, a, b, compareEdgeMaps)
}

func TestDijkstra_BuildGraph(t *testing.T) {
	g := Graph{}
	vA := Vertex(1)
	vB := Vertex(2)
	vC := Vertex(3)
	vD := Vertex(4)

	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddVertex(vD)

	expected_vertices := []Vertex{vA, vB, vC, vD}
	utils.AssertEqualArraysUnordered(t, g.Vertices(), expected_vertices)

	g.AddUndirectedEdge(vA, vB, 1)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vA, vB, 1}},
		vB: {{vB, vA, 1}},
	})

	g.AddUndirectedEdge(vA, vC, 2)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vA, vB, 1}, {vA, vC, 2}},
		vB: {{vB, vA, 1}},
		vC: {{vC, vA, 2}},
	})

	g.AddUndirectedEdge(vB, vC, 3)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vA, vB, 1}, {vA, vC, 2}},
		vB: {{vB, vA, 1}, {vB, vC, 3}},
		vC: {{vC, vA, 2}, {vC, vB, 3}},
	})

	g.AddUndirectedEdge(vC, vD, 4)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vA, vB, 1}, {vA, vC, 2}},
		vB: {{vB, vA, 1}, {vB, vC, 3}},
		vC: {{vC, vA, 2}, {vC, vB, 3}, {vC, vD, 4}},
		vD: {{vD, vC, 4}},
	})
}

func TestDijkstra_AddExistingVertex(t *testing.T) {
	g := Graph{}
	vA := Vertex(1)

	g.AddVertex(vA)
	g.AddVertex(vA)

	expected_vertices := []Vertex{vA}
	utils.AssertEqualArrays(t, g.Vertices(), expected_vertices)
}

func TestDijkstra_AddExistingEdge(t *testing.T) {
	g := Graph{}
	vA := Vertex(1)
	vB := Vertex(2)

	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddUndirectedEdge(vA, vB, 1)
	g.AddUndirectedEdge(vA, vB, 2)

	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vA, vB, 1}, {vA, vB, 2}},
		vB: {{vB, vA, 1}, {vB, vA, 2}},
	})
}
