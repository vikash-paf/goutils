package tuple

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPair(t *testing.T) {
	p := NewPair(1, "hello")
	if p.Left != 1 || p.Right != "hello" {
		t.Errorf("NewPair() = %v, want {1 hello}", p)
	}
}

func TestZip(t *testing.T) {
	lefts := []int{1, 2, 3}
	rights := []string{"a", "b"}

	got := Zip(lefts, rights)
	want := []Pair[int, string]{
		{1, "a"},
		{2, "b"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Zip() = %v, want %v", got, want)
	}
}

func TestUnzip(t *testing.T) {
	pairs := []Pair[int, string]{
		{1, "a"},
		{2, "b"},
	}

	lefts, rights := Unzip(pairs)
	wantLefts := []int{1, 2}
	wantRights := []string{"a", "b"}

	if !reflect.DeepEqual(lefts, wantLefts) {
		t.Errorf("Unzip lefts = %v, want %v", lefts, wantLefts)
	}
	if !reflect.DeepEqual(rights, wantRights) {
		t.Errorf("Unzip rights = %v, want %v", rights, wantRights)
	}
}

func ExampleZip() {
	names := []string{"Alice", "Bob"}
	ages := []int{30, 25, 40} // 40 will be ignored since names is shorter

	pairs := Zip(names, ages)
	fmt.Println(pairs)
	// Output: [{Alice 30} {Bob 25}]
}
