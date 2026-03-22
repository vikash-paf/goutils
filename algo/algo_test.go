package algo

import "testing"

func TestBinarySearch(t *testing.T) {
	type user struct {
		id   int
		name string
	}
	users := []user{
		{1, "Alice"},
		{5, "Bob"},
		{10, "Charlie"},
	}

	idx := BinarySearch(users, 5, func(u user) int { return u.id })
	if idx != 1 {
		t.Errorf("BinarySearch() = %d, want 1", idx)
	}

	idx2 := BinarySearch(users, 7, func(u user) int { return u.id })
	if idx2 != -1 {
		t.Errorf("BinarySearch() = %d, want -1", idx2)
	}
}
