package dict

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Keys(m)
	sort.Strings(got)
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Values(m)
	sort.Ints(got)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values() = %v, want %v", got, want)
	}
}

func TestMerge(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	got := Merge(m1, m2)
	want := map[string]int{"a": 1, "b": 3, "c": 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge() = %v, want %v", got, want)
	}
}

func TestInvert(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := Invert(m)
	want := map[int]string{1: "a", 2: "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Invert() = %v, want %v", got, want)
	}
}

func TestOmit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Omit(m, "a", "c")
	want := map[string]int{"b": 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Omit() = %v, want %v", got, want)
	}
}

func ExampleMerge() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := Merge(m1, m2)
	fmt.Println(merged["b"])
	fmt.Println(merged["c"])
	// Output:
	// 3
	// 4
}
