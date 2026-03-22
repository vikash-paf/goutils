package algo

import (
	"container/heap"
)

// genericHeap implements heap.Interface for a generic slice.
// We use this as a min-heap to keep track of the K "largest" elements.
type genericHeap[T any] struct {
	items []T
	less  func(a, b T) bool
}

func (h genericHeap[T]) Len() int           { return len(h.items) }
func (h genericHeap[T]) Less(i, j int) bool { return h.less(h.items[i], h.items[j]) }
func (h genericHeap[T]) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *genericHeap[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

func (h *genericHeap[T]) Pop() any {
	old := h.items
	n := len(old)
	x := old[n-1]
	h.items = old[0 : n-1]
	return x
}

// TopK returns the k elements from items that are considered "largest" according to the less function.
// To find the k largest elements, the less function should return true when a < b.
// To find the k smallest elements, the less function should return true when a > b.
// Note: The returned slice is not necessarily sorted.
// Time Complexity: O(N log K), Space Complexity: O(K).
func TopK[T any](items []T, k int, less func(a, b T) bool) []T {
	if k <= 0 {
		return nil
	}
	if k >= len(items) {
		result := make([]T, len(items))
		copy(result, items)
		return result
	}

	h := &genericHeap[T]{
		items: make([]T, 0, k),
		less:  less,
	}

	for _, item := range items {
		if h.Len() < k {
			heap.Push(h, item)
		} else if h.less(h.items[0], item) {
			// If the current item is "larger" than the smallest in our top-k heap,
			// we remove the smallest and push the current item.
			heap.Pop(h)
			heap.Push(h, item)
		}
	}

	return h.items
}
