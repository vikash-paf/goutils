package algo

import (
	"fmt"
	"sort"
	"testing"
)

func TestTopK(t *testing.T) {
	items := []int{1, 5, 2, 9, 8, 3, 7, 4, 6}

	// Find the 3 largest elements
	top3 := TopK(items, 3, func(a, b int) bool { return a < b })

	// The array might not be strictly sorted, so sort before comparison
	sort.Ints(top3)

	expected := []int{7, 8, 9}
	if len(top3) != len(expected) {
		t.Fatalf("Expected len %d, got %d", len(expected), len(top3))
	}
	for i, v := range expected {
		if top3[i] != v {
			t.Errorf("At index %d: expected %d, got %d", i, v, top3[i])
		}
	}
}

func TestTopK_SmallK(t *testing.T) {
	items := []string{"apple", "banana", "kiwi"}
	
	// Find the 1 longest string
	longest := TopK(items, 1, func(a, b string) bool { return len(a) < len(b) })
	
	if len(longest) != 1 || longest[0] != "banana" {
		t.Errorf("Expected [banana], got %v", longest)
	}
}

func ExampleTopK() {
	items := []int{10, 4, 25, 8, 3, 1, 15}
	
	// Get the 2 smallest items
	// Note: We use a > b because TopK naturally keeps the K "largest" items
	// based on the less function. Reversing the condition finds the smallest.
	smallest2 := TopK(items, 2, func(a, b int) bool { return a > b })
	
	sort.Ints(smallest2) // Just to make output deterministic
	fmt.Println(smallest2)
	// Output: [1 3]
}
