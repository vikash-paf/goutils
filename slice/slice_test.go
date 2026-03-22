package slice

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		iteratee func(int) string
		want     []string
	}{
		{
			name:     "nil slice",
			input:    nil,
			iteratee: strconv.Itoa,
			want:     nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			iteratee: strconv.Itoa,
			want:     []string{},
		},
		{
			name:     "valid slice",
			input:    []int{1, 2, 3},
			iteratee: strconv.Itoa,
			want:     []string{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.input, tt.iteratee)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "nil slice",
			input:     nil,
			predicate: func(x int) bool { return x > 1 },
			want:      nil,
		},
		{
			name:      "empty slice",
			input:     []int{},
			predicate: func(x int) bool { return x > 1 },
			want:      []int{},
		},
		{
			name:      "valid slice filter out",
			input:     []int{1, 2, 3},
			predicate: func(x int) bool { return x > 1 },
			want:      []int{2, 3},
		},
		{
			name:      "none pass",
			input:     []int{1, 2, 3},
			predicate: func(x int) bool { return x > 3 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.input, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	got := Reduce([]int{1, 2, 3, 4}, func(acc int, val int) int {
		return acc + val
	}, 10)
	want := 20
	if got != want {
		t.Errorf("Reduce() = %v, want %v", got, want)
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 3, 1, 4})
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unique() = %v, want %v", got, want)
	}
}

func TestUniqueBy(t *testing.T) {
	type user struct {
		id   int
		name string
	}
	input := []user{
		{1, "Alice"},
		{2, "Bob"},
		{1, "Alice Dupe"},
	}
	got := UniqueBy(input, func(u user) int { return u.id })
	want := []user{
		{1, "Alice"},
		{2, "Bob"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("UniqueBy() = %v, want %v", got, want)
	}
}

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunk() = %v, want %v", got, want)
	}
}

func TestGroupBy(t *testing.T) {
	got := GroupBy([]int{1, 2, 3, 4, 5}, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	want := map[string][]int{
		"odd":  {1, 3, 5},
		"even": {2, 4},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func ExampleMap() {
	numbers := []int{1, 2, 3, 4}
	strings := Map(numbers, func(n int) string {
		return strconv.Itoa(n * 2)
	})
	fmt.Println(strings)
	// Output: [2 4 6 8]
}

func ExampleFilter() {
	words := []string{"hello", "world", "goutils", "go"}
	filtered := Filter(words, func(w string) bool {
		return len(w) > 4
	})
	fmt.Println(filtered)
	// Output: [hello world goutils]
}

func ExampleReduce() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(numbers, func(acc, val int) int {
		return acc + val
	}, 0)
	fmt.Println(sum)
	// Output: 15
}

func ExampleUnique() {
	items := []string{"a", "b", "c", "a", "b"}
	uniq := Unique(items)
	fmt.Println(uniq)
	// Output: [a b c]
}

func ExampleChunk() {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	chunks := Chunk(items, 3)
	fmt.Println(chunks)
	// Output: [[1 2 3] [4 5 6] [7]]
}

func ExampleGroupBy() {
	items := []string{"one", "two", "three"}
	grouped := GroupBy(items, func(s string) int {
		return len(s)
	})
	fmt.Println(grouped[3])
	// Output: [one two]
}
