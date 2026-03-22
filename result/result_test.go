package result

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestResult(t *testing.T) {
	r := Ok(42)
	if !r.IsOk() || r.IsErr() {
		t.Error("Ok(42) should be Ok")
	}
	if r.Unwrap() != 42 {
		t.Errorf("Unwrap() = %v, want 42", r.Unwrap())
	}

	err := errors.New("fail")
	re := Err[int](err)
	if !re.IsErr() || re.IsOk() {
		t.Error("Err should be Err")
	}
	if re.UnwrapOr(10) != 10 {
		t.Errorf("UnwrapOr() = %v, want 10", re.UnwrapOr(10))
	}
}

func TestChaining(t *testing.T) {
	r := Ok("10")

	// Chain Map and AndThen
	res := AndThen(Map(r, func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}), func(i int) Result[int] {
		if i > 5 {
			return Ok(i * 2)
		}
		return Err[int](errors.New("too small"))
	})

	if !res.IsOk() || res.Unwrap() != 20 {
		t.Errorf("Chain failed: %v", res.Value())
	}
}

func ExampleResult() {
	divide := func(a, b float64) Result[float64] {
		if b == 0 {
			return Err[float64](errors.New("division by zero"))
		}
		return Ok(a / b)
	}

	res := divide(10, 2)
	if res.IsOk() {
		fmt.Println("Result:", res.Unwrap())
	}
	// Output: Result: 5
}
