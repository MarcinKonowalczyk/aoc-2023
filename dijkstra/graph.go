package dijkstra

import (
	"sync"
)

type Graph struct {
	vertices map[*Vertex]struct{} // Set of vertices
	Edges    map[*Vertex][]*Edge  // Edges comming out of a vertex
	lock     sync.RWMutex
}

// AddNode adds a node to the graph
func (g *Graph) AddVertex(n *Vertex) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.vertices == nil {
		g.vertices = make(map[*Vertex]struct{})
	}
	g.vertices[n] = struct{}{}
}

// Return a list of vertices in the graph
func (g *Graph) Vertices() []*Vertex {
	g.lock.RLock()
	defer g.lock.RUnlock()
	vertices := make([]*Vertex, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Vertex, weight int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[*Vertex][]*Edge)
	}
	g.Edges[n1] = append(g.Edges[n1], &Edge{
		Vertex: n2,
		Weight: weight,
	})
	g.Edges[n2] = append(g.Edges[n2], &Edge{
		Vertex: n1,
		Weight: weight,
	})
}
