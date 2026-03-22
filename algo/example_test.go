package algo_test

import (
	"fmt"
	"sort"
	"github.com/vikash-paf/goutils/algo"
)

func ExampleTopK() {
	items := []int{10, 4, 25, 8, 3, 1, 15}

	// Find the 3 largest numbers. Less function returns a < b to keep larger elements.
	largest3 := algo.TopK(items, 3, func(a, b int) bool { return a < b })
	sort.Ints(largest3)
	fmt.Println(largest3)
	// Output: [10 15 25]
}

func ExampleBinarySearch() {
	items := []int{1, 3, 5, 7, 9}
	idx := algo.BinarySearch(items, 5, func(v int) int { return v })
	fmt.Println(idx)
	// Output: 2
}
