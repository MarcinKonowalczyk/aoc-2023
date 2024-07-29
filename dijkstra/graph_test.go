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

	g.AddEdge(vA, vB, 1)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vB, 1}},
		vB: {{vA, 1}},
	})

	g.AddEdge(vA, vC, 2)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vB, 1}, {vC, 2}},
		vB: {{vA, 1}},
		vC: {{vA, 2}},
	})

	g.AddEdge(vB, vC, 3)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vB, 1}, {vC, 2}},
		vB: {{vA, 1}, {vC, 3}},
		vC: {{vA, 2}, {vB, 3}},
	})

	g.AddEdge(vC, vD, 4)
	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vB, 1}, {vC, 2}},
		vB: {{vA, 1}, {vC, 3}},
		vC: {{vA, 2}, {vB, 3}, {vD, 4}},
		vD: {{vC, 4}},
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
	g.AddEdge(vA, vB, 1)
	g.AddEdge(vA, vB, 2)

	assertEqualEdgeMaps(t, g.Edges, map[Vertex][]*Edge{
		vA: {{vB, 1}, {vB, 2}},
		vB: {{vA, 1}, {vA, 2}},
	})
}
