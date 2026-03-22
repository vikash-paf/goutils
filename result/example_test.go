package result_test

import (
	"errors"
	"fmt"
	"github.com/vikash-paf/goutils/result"
)

func ExampleResult_UnwrapOr() {
	r1 := result.Ok(42)
	r2 := result.Err[int](errors.New("fail"))
	
	fmt.Println(r1.UnwrapOr(0))
	fmt.Println(r2.UnwrapOr(0))
	// Output:
	// 42
	// 0
}

func ExampleMap() {
	r := result.Ok(10)
	mapped := result.Map(r, func(n int) int { return n * 2 })
	fmt.Println(mapped.Unwrap())
	// Output: 20
}

func ExampleAndThen() {
	r := result.Ok(10)
	stepped := result.AndThen(r, func(n int) result.Result[int] {
		return result.Ok(n + 5)
	})
	fmt.Println(stepped.Unwrap())
	// Output: 15
}

func ExampleResult_Match() {
	r := result.Ok("success")
	r.Match(
		func(s string) { fmt.Println("OK:", s) },
		func(err error) { fmt.Println("ERR:", err) },
	)
	// Output: OK: success
}
