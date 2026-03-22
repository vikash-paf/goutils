package iterx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/iterx"
)

func ExampleFromSlice() {
	nums := []int{1, 2, 3}
	seq := iterx.FromSlice(nums)
	for n := range seq {
		fmt.Println(n)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleToSlice() {
	nums := []int{1, 2, 3}
	seq := iterx.FromSlice(nums)
	slice := iterx.ToSlice(seq)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func ExampleMap() {
	nums := []int{1, 2, 3}
	seq := iterx.FromSlice(nums)
	mapped := iterx.Map(seq, func(n int) int { return n * 10 })
	fmt.Println(iterx.ToSlice(mapped))
	// Output: [10 20 30]
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4}
	seq := iterx.FromSlice(nums)
	filtered := iterx.Filter(seq, func(n int) bool { return n%2 == 0 })
	fmt.Println(iterx.ToSlice(filtered))
	// Output: [2 4]
}

func ExampleReduce() {
	nums := []int{1, 2, 3, 4}
	seq := iterx.FromSlice(nums)
	sum := iterx.Reduce(seq, func(acc, n int) int { return acc + n }, 0)
	fmt.Println(sum)
	// Output: 10
}

func ExampleTake() {
	nums := []int{1, 2, 3, 4, 5}
	seq := iterx.FromSlice(nums)
	taken := iterx.Take(seq, 3)
	fmt.Println(iterx.ToSlice(taken))
	// Output: [1 2 3]
}

func ExampleGroupBy() {
	nums := []int{1, 2, 3, 4, 5, 6}
	seq := iterx.FromSlice(nums)
	grouped := iterx.GroupBy(seq, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println(len(grouped["even"]), len(grouped["odd"]))
	// Output: 3 3
}
