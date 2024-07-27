package dijkstra

import (
	"sync"
)

// Vertex is the interface that describes the methods that a Vertex must implement
// to be used in the Dijkstra algorithm
type Vertex struct {
	Distance int
}

type NodeQueue struct {
	Items []Vertex
	Lock  sync.Mutex
}

// Enqueue adds an Node to the end of the queue
func (s *NodeQueue) Enqueue(t Vertex) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	N := len(s.Items)
	if N == 0 || t.Distance < s.Items[0].Distance {
		// List is empty or t is smaller than the first element
		// Insert at the beginning
		s.Items = append([]Vertex{t}, s.Items...)
	} else if t.Distance > s.Items[N-1].Distance {
		// t is larger than the last element
		s.Items = append(s.Items, t)
	} else {
		// Insert in distance order by bisecting the list
		lo, hi := 0, N
		for lo < hi {
			mid := lo + (hi-lo)/2
			if t.Distance < s.Items[mid].Distance {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		// Make room for the new element
		s.Items = append(s.Items[:lo+1], s.Items[lo:]...)
		// Insert the new element
		s.Items[lo] = t
	}
}
