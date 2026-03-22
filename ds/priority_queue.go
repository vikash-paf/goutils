package ds

import "container/heap"

// genericHeap implements heap.Interface for a generic slice.
type genericHeap[T any] struct {
	items []T
	less  func(a, b T) bool
}

func (h *genericHeap[T]) Len() int           { return len(h.items) }
func (h *genericHeap[T]) Less(i, j int) bool { return h.less(h.items[i], h.items[j]) }
func (h *genericHeap[T]) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *genericHeap[T]) Push(x any)         { h.items = append(h.items, x.(T)) }
func (h *genericHeap[T]) Pop() any {
	old := h.items
	n := len(old)
	item := old[n-1]
	h.items = old[0 : n-1]
	return item
}

// PriorityQueue is a generic priority queue using an underlying heap.
type PriorityQueue[T any] struct {
	h *genericHeap[T]
}

// NewPriorityQueue creates a new PriorityQueue. The 'less' function determines the priority.
// For a min-heap, 'less' should return true if a < b. For a max-heap, it should return true if a > b.
func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	h := &genericHeap[T]{
		items: make([]T, 0),
		less:  less,
	}
	heap.Init(h)
	return &PriorityQueue[T]{h: h}
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue[T]) Push(item T) {
	heap.Push(pq.h, item)
}

// Pop removes and returns the highest priority item.
// It returns the zero value and false if the queue is empty.
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if pq.h.Len() == 0 {
		var zero T
		return zero, false
	}
	return heap.Pop(pq.h).(T), true
}

// Peek returns the highest priority item without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.h.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.h.items[0], true
}

// Len returns the number of items in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return pq.h.Len()
}
