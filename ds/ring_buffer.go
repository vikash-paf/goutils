package ds

// RingBuffer is a generic, fixed-size circular buffer.
type RingBuffer[T any] struct {
	items    []T
	capacity int
	head     int
	tail     int
	count    int
}

// NewRingBuffer creates a new RingBuffer with the specified capacity.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		panic("RingBuffer capacity must be > 0")
	}
	return &RingBuffer[T]{
		items:    make([]T, capacity),
		capacity: capacity,
	}
}

// Push adds an item to the buffer. If the buffer is full, it overwrites the oldest item.
func (rb *RingBuffer[T]) Push(item T) {
	rb.items[rb.tail] = item
	rb.tail = (rb.tail + 1) % rb.capacity

	if rb.count < rb.capacity {
		rb.count++
	} else {
		rb.head = (rb.head + 1) % rb.capacity
	}
}

// Values returns all items currently in the buffer, ordered from oldest to newest.
func (rb *RingBuffer[T]) Values() []T {
	result := make([]T, 0, rb.count)
	for i := 0; i < rb.count; i++ {
		idx := (rb.head + i) % rb.capacity
		result = append(result, rb.items[idx])
	}
	return result
}

// IsFull returns true if the buffer has reached its capacity.
func (rb *RingBuffer[T]) IsFull() bool {
	return rb.count == rb.capacity
}

// Len returns the number of items currently in the buffer.
func (rb *RingBuffer[T]) Len() int {
	return rb.count
}

// Capacity returns the maximum capacity of the buffer.
func (rb *RingBuffer[T]) Capacity() int {
	return rb.capacity
}
