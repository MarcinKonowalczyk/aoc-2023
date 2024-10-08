package dijkstra

import (
	"fmt"
	"sync"
)

type Vertex int

type Edge struct {
	from   Vertex // The vertex at the start of the edge
	to     Vertex // The vertex at the end of the edge
	weight int
}

func (e *Edge) String() string {
	return fmt.Sprintf("E%d(%d)", e.to, e.weight)
}

type Graph struct {
	vertices map[Vertex]struct{} // Set of vertices
	Edges    map[Vertex][]*Edge  // Edges coming out of a vertex
	lock     sync.RWMutex
}

// Add a vertex by pointer to the graph
func (g *Graph) AddVertex(i Vertex) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.vertices == nil {
		g.vertices = make(map[Vertex]struct{})
	}
	g.vertices[i] = struct{}{}
}

// Return a list of vertices in the graph
func (g *Graph) Vertices() []Vertex {
	g.lock.RLock()
	defer g.lock.RUnlock()
	vertices := make([]Vertex, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

// AddEdge adds an edge to the graph
func (g *Graph) AddDirectedEdge(v1, v2 Vertex, weight int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[Vertex][]*Edge)
	}
	g.Edges[v1] = append(g.Edges[v1], &Edge{
		from:   v1,
		to:     v2,
		weight: weight,
	})
}

func (g *Graph) AddUndirectedEdge(v1, v2 Vertex, weight int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[Vertex][]*Edge)
	}
	g.Edges[v1] = append(g.Edges[v1], &Edge{
		from:   v1,
		to:     v2,
		weight: weight,
	})
	g.Edges[v2] = append(g.Edges[v2], &Edge{
		from:   v2,
		to:     v1,
		weight: weight,
	})
}
