package tuple_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/tuple"
)

func ExampleNewPair() {
	p := tuple.NewPair("age", 30)
	fmt.Println(p.Left, p.Right)
	// Output: age 30
}

func ExampleZip() {
	names := []string{"Alice", "Bob"}
	ages := []int{30, 25}
	zipped := tuple.Zip(names, ages)
	fmt.Println(zipped[0].Left, zipped[0].Right)
	// Output: Alice 30
}

func ExampleUnzip() {
	pairs := []tuple.Pair[string, int]{
		{"Alice", 30},
		{"Bob", 25},
	}
	names, ages := tuple.Unzip(pairs)
	fmt.Println(names, ages)
	// Output: [Alice Bob] [30 25]
}
