package ptr

import (
	"fmt"
	"testing"
)

func TestOf(t *testing.T) {
	val := 42
	p := Of(val)
	if p == nil || *p != val {
		t.Errorf("Of() = %v, want pointer to %v", p, val)
	}
}

func TestVal(t *testing.T) {
	var nilPtr *int
	if got := Val(nilPtr); got != 0 {
		t.Errorf("Val(nil) = %v, want 0", got)
	}

	val := 42
	if got := Val(&val); got != 42 {
		t.Errorf("Val(&val) = %v, want 42", got)
	}
}

func TestValOrDefault(t *testing.T) {
	var nilPtr *string
	if got := ValOrDefault(nilPtr, "default"); got != "default" {
		t.Errorf("ValOrDefault(nil, default) = %v, want 'default'", got)
	}

	val := "hello"
	if got := ValOrDefault(&val, "default"); got != "hello" {
		t.Errorf("ValOrDefault(&val, default) = %v, want 'hello'", got)
	}
}

func TestEqual(t *testing.T) {
	v1, v2 := 1, 1
	v3 := 2

	if !Equal(&v1, &v2) {
		t.Errorf("Equal(&v1, &v2) should be true")
	}
	if Equal(&v1, &v3) {
		t.Errorf("Equal(&v1, &v3) should be false")
	}

	var nilPtr1 *int
	var nilPtr2 *int
	if !Equal(nilPtr1, nilPtr2) {
		t.Errorf("Equal(nil, nil) should be true")
	}
	if Equal(&v1, nilPtr1) {
		t.Errorf("Equal(&v1, nil) should be false")
	}
}

func ExampleOf() {
	p := Of("hello")
	fmt.Println(*p)
	// Output: hello
}

func ExampleVal() {
	var p *int // nil pointer
	fmt.Println(Val(p))

	val := 42
	fmt.Println(Val(&val))
	// Output:
	// 0
	// 42
}
