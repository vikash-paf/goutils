package ds

import (
	"reflect"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer[int](3)

	if rb.Capacity() != 3 {
		t.Errorf("Capacity() = %d, want 3", rb.Capacity())
	}

	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	if !rb.IsFull() {
		t.Error("Buffer should be full")
	}

	if !reflect.DeepEqual(rb.Values(), []int{1, 2, 3}) {
		t.Errorf("Values() = %v, want [1 2 3]", rb.Values())
	}

	// Overwrite oldest
	rb.Push(4)
	if !reflect.DeepEqual(rb.Values(), []int{2, 3, 4}) {
		t.Errorf("Values() = %v, want [2 3 4]", rb.Values())
	}

	rb.Push(5)
	if !reflect.DeepEqual(rb.Values(), []int{3, 4, 5}) {
		t.Errorf("Values() = %v, want [3 4 5]", rb.Values())
	}
}
