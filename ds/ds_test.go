package ds

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var s Stack[int]

	s.Push(1)
	s.Push(2)

	if l := s.Len(); l != 2 {
		t.Errorf("Len() = %d, want 2", l)
	}

	val, ok := s.Peek()
	if !ok || val != 2 {
		t.Errorf("Peek() = %v, %v, want 2, true", val, ok)
	}

	val, ok = s.Pop()
	if !ok || val != 2 {
		t.Errorf("Pop() = %v, want 2", val)
	}

	val, ok = s.Pop()
	if !ok || val != 1 {
		t.Errorf("Pop() = %v, want 1", val)
	}

	_, ok = s.Pop()
	if ok {
		t.Error("Pop() on empty stack should return false")
	}

	s.Push(10)
	s.Clear()
	if s.Len() != 0 {
		t.Error("Clear() failed")
	}
}

func TestQueue(t *testing.T) {
	var q Queue[int]

	q.Enqueue(1)
	q.Enqueue(2)

	if l := q.Len(); l != 2 {
		t.Errorf("Len() = %d, want 2", l)
	}

	val, ok := q.Peek()
	if !ok || val != 1 {
		t.Errorf("Peek() = %v, want 1", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue() = %v, want 1", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Dequeue() = %v, want 2", val)
	}

	_, ok = q.Dequeue()
	if ok {
		t.Error("Dequeue() on empty queue should return false")
	}

	// Test Memory Optimization Shift
	q.Clear()
	for i := 0; i < 2000; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 1500; i++ {
		q.Dequeue()
	}

	if l := len(q.items); l > 1000 {
		t.Errorf("Queue underlying slice didn't shift effectively, size %v", l)
	}
}

func ExampleQueue() {
	var q Queue[string]
	q.Enqueue("Job 1")
	q.Enqueue("Job 2")

	for q.Len() > 0 {
		if job, ok := q.Dequeue(); ok {
			fmt.Println("Processing:", job)
		}
	}
	// Output:
	// Processing: Job 1
	// Processing: Job 2
}
