package ds

import "testing"

func TestPriorityQueue(t *testing.T) {
	// Min-heap
	pq := NewPriorityQueue[int](func(a, b int) bool { return a < b })

	pq.Push(3)
	pq.Push(1)
	pq.Push(4)
	pq.Push(2)

	if l := pq.Len(); l != 4 {
		t.Errorf("Len() = %d, want 4", l)
	}

	val, ok := pq.Peek()
	if !ok || val != 1 {
		t.Errorf("Peek() = %d, want 1", val)
	}

	expected := []int{1, 2, 3, 4}
	for _, exp := range expected {
		val, ok := pq.Pop()
		if !ok || val != exp {
			t.Errorf("Pop() = %d, want %d", val, exp)
		}
	}

	_, ok = pq.Pop()
	if ok {
		t.Error("Pop() on empty queue should return false")
	}
}
