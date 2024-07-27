package dijkstra

import (
	"fmt"
)

var vertex_counter int     // Global counter for number of vertices initialized
var all_vertices []*Vertex // List of all vertices

type Vertex struct {
	// Distance int
	index int
}

func NewVertex() *Vertex {
	vertex_counter++
	v := Vertex{
		index: vertex_counter,
	}
	all_vertices = append(all_vertices, &v)
	return &v
}

func (n *Vertex) String() string {
	return fmt.Sprintf("V%d", n.index)
}

var edge_counter int  // Global counter for number of edges initialized
var all_edges []*Edge // List of all edges

type Edge struct {
	Vertex *Vertex // The vertex at the end of the edge
	Weight int
}

func (e *Edge) String() string {
	return fmt.Sprintf("E%p(%d)", e.Vertex, e.Weight)
}
