package opt_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/opt"
)

func ExampleSome() {
	o := opt.Some(42)
	fmt.Println(o.IsSome())
	fmt.Println(o.Unwrap())
	// Output:
	// true
	// 42
}

func ExampleNone() {
	o := opt.None[int]()
	fmt.Println(o.IsNone())
	// Output: true
}

func ExampleOpt_UnwrapOr() {
	o := opt.None[int]()
	fmt.Println(o.UnwrapOr(100))
	// Output: 100
}

func ExampleMap() {
	o := opt.Some(10)
	mapped := opt.Map(o, func(n int) int { return n * 2 })
	fmt.Println(mapped.Unwrap())
	// Output: 20
}
