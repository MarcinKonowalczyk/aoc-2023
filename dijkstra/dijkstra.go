package dijkstra

import (
	"fmt"
	"math"
)

// Find the shortest path between two vertices in a graph.
// Both the start and end vertices must be in the graph.
func ShortestPath(g *Graph, start Vertex, end Vertex) ([]Vertex, int) {
	// Make sure the start and end vertices are in the graph
	if _, ok := g.vertices[start]; !ok {
		panic(fmt.Sprintf("Start vertex %v not in graph", start))
	}
	if _, ok := g.vertices[end]; !ok {
		panic(fmt.Sprintf("End vertex %v not in graph", end))
	}

	// // Set the start node distance to 0
	// for v, _ := range g.vertices {
	// 	v.Distance = math.MaxInt64
	// }
	// start.Distance = 0

	// Distance map
	dist := make(map[Vertex]int)
	for v, _ := range g.vertices {
		dist[v] = math.MaxInt64
	}
	dist[start] = 0

	// Priority queue of nodes
	pq := PriorityQueue[Vertex]{}
	pq.Enqueue(start, 0)

	// Visited nodes
	visited := make(map[Vertex]bool)

	// Previous node
	prev := make(map[Vertex]Vertex)

	for !pq.IsEmpty() {
		v, _ := pq.Pop()
		if visited[v] {
			continue
		}
		visited[v] = true
		near, ok := g.Edges[v]
		if !ok {
			// Graph consistantcy check
			panic(fmt.Sprintf("Vertex %v not in graph", v))
		}

		for _, e := range near {
			if !visited[e.from] {
				d := dist[v] + e.weight
				if d < dist[e.from] {
					dist[e.from] = d
					prev[e.from] = v
					pq.Enqueue(e.from, d)
				}
				// visited[e.Vertex] = true
			}
		}
	}

	// Reconstruct the path
	path_val := prev[end]
	var path []Vertex
	path = append(path, end)
	for path_val != start {
		path = append(path, path_val)
		path_val = prev[path_val]
	}
	path = append(path, path_val)

	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, dist[end]
}
