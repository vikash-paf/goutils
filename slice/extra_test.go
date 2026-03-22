package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	res := Find(items, func(v int) bool { return v == 3 })
	if res == nil || *res != 3 {
		t.Errorf("Find() = %v, want pointer to 3", res)
	}

	res2 := Find(items, func(v int) bool { return v == 10 })
	if res2 != nil {
		t.Errorf("Find() = %v, want nil", res2)
	}
}

func TestFindIndex(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	if idx := FindIndex(items, func(v int) bool { return v == 3 }); idx != 2 {
		t.Errorf("FindIndex() = %v, want 2", idx)
	}
	if idx := FindIndex(items, func(v int) bool { return v == 10 }); idx != -1 {
		t.Errorf("FindIndex() = %v, want -1", idx)
	}
}

func TestSomeAndEvery(t *testing.T) {
	items := []int{2, 4, 6, 8}

	if !Some(items, func(v int) bool { return v == 4 }) {
		t.Error("Some() should be true")
	}
	if Some(items, func(v int) bool { return v == 5 }) {
		t.Error("Some() should be false")
	}

	if !Every(items, func(v int) bool { return v%2 == 0 }) {
		t.Error("Every() should be true")
	}
	if Every(items, func(v int) bool { return v > 2 }) {
		t.Error("Every() should be false")
	}
}

func TestReverseAndShuffle(t *testing.T) {
	items := []int{1, 2, 3}
	Reverse(items)
	if !reflect.DeepEqual(items, []int{3, 2, 1}) {
		t.Errorf("Reverse() failed, got %v", items)
	}

	Shuffle(items)
	if len(items) != 3 {
		t.Errorf("Shuffle() changed slice length")
	}
}

func TestPartition(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	passed, failed := Partition(items, func(v int) bool { return v%2 == 0 })
	if !reflect.DeepEqual(passed, []int{2, 4}) {
		t.Errorf("Partition passed = %v", passed)
	}
	if !reflect.DeepEqual(failed, []int{1, 3, 5}) {
		t.Errorf("Partition failed = %v", failed)
	}
}

func TestDiffState(t *testing.T) {
	old := []int{1, 2, 3}
	new := []int{2, 3, 4}
	added, removed := DiffState(old, new)

	if !reflect.DeepEqual(added, []int{4}) {
		t.Errorf("DiffState added = %v, want [4]", added)
	}
	if !reflect.DeepEqual(removed, []int{1}) {
		t.Errorf("DiffState removed = %v, want [1]", removed)
	}
}

func TestCountBy(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	count := CountBy(items, func(v int) bool { return v > 2 })
	if count != 3 {
		t.Errorf("CountBy() = %v, want 3", count)
	}
}

func ExampleDiffState() {
	oldRoles := []string{"admin", "editor"}
	newRoles := []string{"editor", "viewer"}

	added, removed := DiffState(oldRoles, newRoles)
	fmt.Printf("Added: %v, Removed: %v", added, removed)
	// Output: Added: [viewer], Removed: [admin]
}
