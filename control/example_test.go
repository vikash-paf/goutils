package control_test

import (
	"errors"
	"fmt"
	"github.com/vikash-paf/goutils/control"
)

func ExampleIf() {
	status := control.If(len("OK") == 2, "Healthy", "Dead")
	fmt.Println(status)
	// Output: Healthy
}

func ExampleMust() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	_ = control.Must(10, errors.New("critical error"))
	// Output: Recovered from panic
}

func ExampleCoalesce() {
	val := control.Coalesce("", "hello", "world")
	fmt.Println(val)
	// Output: hello
}

func ExampleTry() {
	val := control.Try("fallback", func() (string, error) {
		return "", errors.New("fail")
	})
	fmt.Println(val)
	// Output: fallback
}
