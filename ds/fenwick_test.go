package ds

import (
	"fmt"
	"math"
	"testing"
)

func TestFenwickTree(t *testing.T) {
	// Underlying conceptual array: [0, 0, 0, 0, 0]
	ft := NewFenwickTree[int](5)

	ft.Add(0, 1)
	ft.Add(1, 2)
	ft.Add(2, 3)
	ft.Add(3, 4)
	ft.Add(4, 5)
	// Now: [1, 2, 3, 4, 5]

	if sum := ft.PrefixSum(2); sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}

	if sum := ft.RangeSum(1, 3); sum != 9 {
		t.Errorf("expected 9, got %d", sum)
	}

	if sum := ft.RangeSum(3, 1); sum != 0 {
		t.Errorf("expected 0, got %d", sum)
	}

	// Update index 2 by adding 2 (so the value at index 2 goes from 3 to 5)
	ft.Add(2, 2)
	// Now: [1, 2, 5, 4, 5]

	if sum := ft.PrefixSum(2); sum != 8 {
		t.Errorf("expected 8, got %d", sum)
	}
}

func TestFenwickTree_Float(t *testing.T) {
	ft := NewFenwickTree[float64](3)
	ft.Add(0, 1.5)
	ft.Add(1, 2.5)

	if sum := ft.RangeSum(0, 1); math.Abs(sum-4.0) > 0.0001 {
		t.Errorf("expected 4.0, got %f", sum)
	}
}

func ExampleFenwickTree() {
	// A tree to track values over 5 elements
	ft := NewFenwickTree[int](5)

	// Add occurrences/values at specific 0-based indices
	ft.Add(0, 10)
	ft.Add(2, 5)
	ft.Add(4, 15)

	// Calculate prefix sum from index 0 to 3
	fmt.Println("Prefix sum up to index 3:", ft.PrefixSum(3))

	// Calculate sum within a specific range
	fmt.Println("Range sum from 1 to 4:", ft.RangeSum(1, 4))

	// Output:
	// Prefix sum up to index 3: 15
	// Range sum from 1 to 4: 20
}
