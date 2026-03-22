package parallel

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestMap(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	results := Map(inputs, func(v int) int {
		return v * 2
	})

	expected := []int{2, 4, 6, 8, 10}
	for i, v := range results {
		if v != expected[i] {
			t.Errorf("At index %d: got %d, want %d", i, v, expected[i])
		}
	}
}

func TestForEach(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	var count atomic.Int32
	ForEach(inputs, func(v int) {
		count.Add(1)
	})

	if count.Load() != 5 {
		t.Errorf("Expected count 5, got %d", count.Load())
	}
}

func TestMapBatched(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	results := MapBatched(inputs, 2, func(v int) int {
		return v * 2
	})

	expected := []int{2, 4, 6, 8, 10}
	for i, v := range results {
		if v != expected[i] {
			t.Errorf("At index %d: got %d, want %d", i, v, expected[i])
		}
	}
}

func TestForEachBatched(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	var count atomic.Int32
	ForEachBatched(inputs, 2, func(v int) {
		count.Add(1)
	})

	if count.Load() != 5 {
		t.Errorf("Expected count 5, got %d", count.Load())
	}
}

func ExampleMap() {
	inputs := []string{"a", "bb", "ccc"}
	lengths := Map(inputs, func(v string) int {
		return len(v)
	})
	fmt.Println(lengths)
	// Output: [1 2 3]
}
