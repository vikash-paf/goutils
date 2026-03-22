package opt

import (
	"fmt"
	"strconv"
	"testing"
)

func TestOpt(t *testing.T) {
	val := Some(42)
	if !val.IsSome() || val.IsNone() {
		t.Error("Some(42) should be Some")
	}
	if got := val.Unwrap(); got != 42 {
		t.Errorf("Unwrap() = %v, want 42", got)
	}

	none := None[int]()
	if !none.IsNone() || none.IsSome() {
		t.Error("None() should be None")
	}
	if got := none.UnwrapOr(10); got != 10 {
		t.Errorf("UnwrapOr() = %v, want 10", got)
	}
}

func TestOptUnwrapPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Unwrap() on None should panic")
		}
	}()
	none := None[string]()
	none.Unwrap()
}

func TestMap(t *testing.T) {
	val := Some(10)
	mapped := Map(val, func(v int) string {
		return strconv.Itoa(v * 2)
	})

	if !mapped.IsSome() || mapped.Unwrap() != "20" {
		t.Errorf("Map() = %v, want Some(20)", mapped.UnwrapOr(""))
	}

	none := None[int]()
	mappedNone := Map(none, strconv.Itoa)
	if !mappedNone.IsNone() {
		t.Error("Map() on None should return None")
	}
}

func ExampleOpt() {
	var user map[string]string = map[string]string{"name": "Alice"}

	// A safe lookup that returns an Optional
	lookup := func(key string) Opt[string] {
		if val, ok := user[key]; ok {
			return Some(val)
		}
		return None[string]()
	}

	nameOpt := lookup("name")
	fmt.Println(nameOpt.UnwrapOr("Anonymous"))

	ageOpt := lookup("age")
	fmt.Println(ageOpt.UnwrapOr("Unknown"))
	// Output:
	// Alice
	// Unknown
}
