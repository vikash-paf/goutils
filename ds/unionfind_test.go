package ds

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[string]()

	uf.Add("A")
	uf.Add("B")
	uf.Add("C")

	if uf.Count() != 3 {
		t.Errorf("expected 3 components, got %d", uf.Count())
	}

	if uf.Connected("A", "B") {
		t.Error("A and B should not be connected initially")
	}

	merged := uf.Union("A", "B")
	if !merged {
		t.Error("expected Union to return true when merging distinct sets")
	}

	if uf.Count() != 2 {
		t.Errorf("expected 2 components after one merge, got %d", uf.Count())
	}

	if !uf.Connected("A", "B") {
		t.Error("A and B should be connected")
	}

	merged = uf.Union("B", "A") // Already connected
	if merged {
		t.Error("expected Union to return false when merging already connected items")
	}

	uf.Union("B", "C")
	if uf.Count() != 1 {
		t.Errorf("expected 1 component, got %d", uf.Count())
	}

	if !uf.Connected("A", "C") {
		t.Error("A and C should be connected transitively")
	}

	// Finding non-existent item
	if _, ok := uf.Find("D"); ok {
		t.Error("expected Find to return false for missing item")
	}
}

func ExampleUnionFind() {
	uf := NewUnionFind[int]()
	uf.Add(1)
	uf.Add(2)
	uf.Add(3)

	uf.Union(1, 2)

	if uf.Connected(1, 2) {
		fmt.Println("1 and 2 belong to the same component")
	}

	if !uf.Connected(1, 3) {
		fmt.Println("1 and 3 are disconnected")
	}

	fmt.Println("Total independent networks:", uf.Count())
	// Output:
	// 1 and 2 belong to the same component
	// 1 and 3 are disconnected
	// Total independent networks: 2
}
