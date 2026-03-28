package ptr_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/ptr"
)

func ExampleOf() {
	p := ptr.Of("hello")
	fmt.Println(*p)
	// Output: hello
}

func ExampleVal() {
	var s *string
	fmt.Println(ptr.Val(s)) // Returns zero value ""

	p := ptr.Of("hello")
	fmt.Println(ptr.Val(p))
	// Output:
	//
	// hello
}

func ExampleValOrDefault() {
	var s *string
	fmt.Println(ptr.ValOrDefault(s, "default"))

	p := ptr.Of("hello")
	fmt.Println(ptr.ValOrDefault(p, "default"))
	// Output:
	// default
	// hello
}

func ExampleEqual() {
	p1 := ptr.Of(10)
	p2 := ptr.Of(10)
	p3 := ptr.Of(20)

	fmt.Println(ptr.Equal(p1, p2))
	fmt.Println(ptr.Equal(p1, p3))
	// Output:
	// true
	// false
}
