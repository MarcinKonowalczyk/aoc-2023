package dijkstra

import "sync"

type queueItem[T any] struct {
	value    T
	distance int
}

type PriorityQueue[T any] struct {
	items []queueItem[T]
	lock  sync.Mutex
}

func (pq *PriorityQueue[T]) Items() []T {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	items := make([]T, len(pq.items))
	for i, item := range pq.items {
		items[i] = item.value
	}
	return items
}

// Enqueue adds an Node to the end of the queue
func (pq *PriorityQueue[T]) Enqueue(t T, d int) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	N := len(pq.items)
	new_item := queueItem[T]{value: t, distance: d}
	if N == 0 || d < pq.items[0].distance {
		// List is empty or t is smaller than the first element
		// Insert at the beginning
		// pq.items = append([]*Vertex{t}, pq.Items...)
		pq.items = append([]queueItem[T]{new_item}, pq.items...)
	} else if d > pq.items[N-1].distance {
		// t is larger than the last element
		pq.items = append(pq.items, new_item)
	} else {
		// Insert in distance order by bisecting the list
		lo, hi := 0, N
		for lo < hi {
			mid := lo + (hi-lo)/2
			if d < pq.items[mid].distance {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		before := pq.items[:lo]
		after := make([]queueItem[T], N-lo)
		copy(after, pq.items[lo:])
		pq.items = append(append(before, new_item), after...)
	}
}

// Remove a Node from the start of the queue
func (pq *PriorityQueue[T]) Pop() (T, int) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	item := pq.items[0]
	pq.items = pq.items[1:len(pq.items)]
	return item.value, item.distance
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return len(pq.items) == 0
}

func (pq *PriorityQueue[T]) Size() int {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return len(pq.items)
}
