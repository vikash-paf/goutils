// Package ds provides common data structures, such as Stacks, Queues,
// Priority Queues, and Ring Buffers.
//
// Usage:
//
//	s := ds.Stack[int]{}
//	s.Push(1)
//	val, ok := s.Pop()
package ds

// Stack is a generic LIFO (Last-In-First-Out) data structure.
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack.
// It returns the zero value and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx]
	return item, true
}

// Peek returns the top element of the stack without removing it.
// It returns the zero value and false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.items = nil
}

// Queue is a generic FIFO (First-In-First-Out) data structure.
// It dynamically shifts elements when memory is wasted to prevent continuous slice growth.
type Queue[T any] struct {
	items []T
	head  int
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue.
// It returns the zero value and false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head >= len(q.items) {
		var zero T
		return zero, false
	}

	item := q.items[q.head]
	q.head++

	// Prevent unbounded growth of the underlying slice array
	// Shift elements left if more than half the slice is "dead" space AND it's reasonably large.
	if q.head > len(q.items)/2 && len(q.items) > 1024 {
		copy(q.items, q.items[q.head:])
		q.items = q.items[:len(q.items)-q.head]
		q.head = 0
	}

	return item, true
}

// Peek returns the front element of the queue without removing it.
// It returns the zero value and false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	if q.head >= len(q.items) {
		var zero T
		return zero, false
	}
	return q.items[q.head], true
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(q.items) - q.head
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.items = nil
	q.head = 0
}
