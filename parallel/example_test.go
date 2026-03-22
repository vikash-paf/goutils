package parallel_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/parallel"
)

func ExampleMap() {
	inputs := []int{1, 2, 3}
	results := parallel.Map(inputs, func(n int) int { return n * 2 })
	fmt.Println(results)
	// Output: [2 4 6]
}

func ExampleForEach() {
	inputs := []int{1, 2, 3}
	// Note: ForEach is concurrent, so we'd need a mutex or atomic for a real sum,
	// but here we just show the call.
	parallel.ForEach(inputs, func(n int) {
		_ = n // process n
	})
	fmt.Println("Done")
	// Output: Done
}

func ExampleMapBatched() {
	inputs := []int{1, 2, 3, 4, 5}
	// Max 2 concurrent goroutines
	results := parallel.MapBatched(inputs, 2, func(n int) int { return n * 2 })
	fmt.Println(results)
	// Output: [2 4 6 8 10]
}
