package set

import (
	"fmt"
	"sort"
	"testing"
)

func TestSetBasic(t *testing.T) {
	s := New(1, 2, 3)
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}

	s.Add(4)
	if !s.Contains(4) {
		t.Error("Set should contain 4 after Add")
	}

	s.Remove(2)
	if s.Contains(2) {
		t.Error("Set should not contain 2 after Remove")
	}

	vals := s.Values()
	sort.Ints(vals)
	if len(vals) != 3 || vals[0] != 1 || vals[1] != 3 || vals[2] != 4 {
		t.Errorf("Values() returned %v, want [1 3 4]", vals)
	}
}

func TestSetOperations(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(3, 4, 5)

	union := Union(s1, s2)
	if len(union) != 5 || !union.Contains(1) || !union.Contains(5) {
		t.Errorf("Union failed: %v", union)
	}

	intersection := Intersection(s1, s2)
	if len(intersection) != 1 || !intersection.Contains(3) {
		t.Errorf("Intersection failed: %v", intersection)
	}

	diff := Difference(s1, s2)
	if len(diff) != 2 || !diff.Contains(1) || !diff.Contains(2) {
		t.Errorf("Difference failed: %v", diff)
	}

	symDiff := SymmetricDifference(s1, s2)
	if len(symDiff) != 4 || symDiff.Contains(3) || !symDiff.Contains(1) || !symDiff.Contains(5) {
		t.Errorf("SymmetricDifference failed: %v", symDiff)
	}
}

func TestSetSubSuperset(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(1, 2, 3)

	if !s1.IsSubset(s2) {
		t.Error("s1 should be subset of s2")
	}
	if s2.IsSubset(s1) {
		t.Error("s2 should not be subset of s1")
	}
	if !s2.IsSuperset(s1) {
		t.Error("s2 should be superset of s1")
	}
}

func ExampleSet() {
	s := New("apple", "banana")
	s.Add("cherry")

	if s.Contains("banana") {
		fmt.Println("Has banana")
	}
	// Output: Has banana
}
