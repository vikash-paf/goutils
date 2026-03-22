package iterx

import (
	"fmt"
	"testing"
)

func TestIterx(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	seq := FromSlice(nums)

	// Map
	doubled := Map(seq, func(n int) int { return n * 2 })
	doubledSlice := ToSlice(doubled)
	if len(doubledSlice) != 5 || doubledSlice[0] != 2 || doubledSlice[4] != 10 {
		t.Errorf("Map failed: %v", doubledSlice)
	}

	// Filter
	evens := Filter(seq, func(n int) bool { return n%2 == 0 })
	evensSlice := ToSlice(evens)
	if len(evensSlice) != 2 || evensSlice[0] != 2 || evensSlice[1] != 4 {
		t.Errorf("Filter failed: %v", evensSlice)
	}

	// Reduce
	sum := Reduce(seq, func(acc, n int) int { return acc + n }, 0)
	if sum != 15 {
		t.Errorf("Reduce failed: %d", sum)
	}

	// Take
	firstThree := Take(seq, 3)
	firstThreeSlice := ToSlice(firstThree)
	if len(firstThreeSlice) != 3 || firstThreeSlice[2] != 3 {
		t.Errorf("Take failed: %v", firstThreeSlice)
	}
}

func TestGroupBy(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	seq := FromSlice(nums)

	grouped := GroupBy(seq, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})

	if len(grouped["even"]) != 3 || len(grouped["odd"]) != 3 {
		t.Errorf("GroupBy failed: %v", grouped)
	}
}

func ExampleMap() {
	nums := []int{1, 2, 3}
	seq := FromSlice(nums)
	doubled := Map(seq, func(n int) int { return n * 2 })
	for n := range doubled {
		fmt.Println(n)
	}
	// Output:
	// 2
	// 4
	// 6
}
