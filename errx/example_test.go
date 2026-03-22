package errx_test

import (
	"errors"
	"fmt"
	"github.com/vikash-paf/goutils/errx"
)

func ExampleMultiError() {
	m := &errx.MultiError{}
	m.Append(errors.New("error 1"))
	m.Append(errors.New("error 2"))
	
	if err := m.AsError(); err != nil {
		fmt.Println(m.HasErrors())
	}
	// Output: true
}
