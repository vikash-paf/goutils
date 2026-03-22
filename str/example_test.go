package str_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/str"
)

func ExampleIsBlank() {
	fmt.Println(str.IsBlank(""))
	fmt.Println(str.IsBlank("  "))
	fmt.Println(str.IsBlank("hello"))
	// Output:
	// true
	// true
	// false
}

func ExampleReverse() {
	fmt.Println(str.Reverse("hello"))
	// Output: olleh
}

func ExampleTruncate() {
	fmt.Println(str.Truncate("hello world", 8, "..."))
	// Output: hello...
}

func ExampleToCamelCase() {
	fmt.Println(str.ToCamelCase("hello_world"))
	// Output: helloWorld
}

func ExampleToSnakeCase() {
	fmt.Println(str.ToSnakeCase("helloWorld"))
	// Output: hello_world
}

func ExampleLevenshtein() {
	fmt.Println(str.Levenshtein("kitten", "sitting"))
	// Output: 3
}

func ExampleJaroWinkler() {
	// JaroWinkler similarity is between 0.0 and 1.0.
	sim := str.JaroWinkler("martha", "marhta")
	fmt.Println(sim > 0.9)
	// Output: true
}
