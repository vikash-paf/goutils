package dict_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/dict"
)

func ExampleKeys() {
	m := map[string]int{"a": 1, "b": 2}
	keys := dict.Keys(m)
	fmt.Println(len(keys))
	// Output: 2
}

func ExampleValues() {
	m := map[string]int{"a": 1, "b": 2}
	values := dict.Values(m)
	fmt.Println(len(values))
	// Output: 2
}

func ExampleMerge() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 99, "c": 3}
	merged := dict.Merge(m1, m2)
	fmt.Println(merged["b"])
	// Output: 99
}

func ExampleInvert() {
	m := map[string]int{"a": 1, "b": 2}
	inverted := dict.Invert(m)
	fmt.Println(inverted[1])
	// Output: a
}

func ExampleOmit() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	omitted := dict.Omit(m, "a", "c")
	fmt.Println(len(omitted))
	fmt.Println(omitted["b"])
	// Output:
	// 1
	// 2
}
