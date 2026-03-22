package control

import (
	"errors"
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	if got := If(true, "yes", "no"); got != "yes" {
		t.Errorf("If(true) = %v, want 'yes'", got)
	}
	if got := If(false, 1, 0); got != 0 {
		t.Errorf("If(false) = %v, want 0", got)
	}
}

func TestMust(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must() did not panic on error")
		}
	}()

	// Should not panic
	val := Must("success", nil)
	if val != "success" {
		t.Errorf("Must() returned %v, want 'success'", val)
	}

	// Should panic
	Must("", errors.New("fatal"))
}

func TestCoalesce(t *testing.T) {
	if got := Coalesce("", "", "hello", "world"); got != "hello" {
		t.Errorf("Coalesce() = %v, want 'hello'", got)
	}
	if got := Coalesce(0, 0, 0); got != 0 {
		t.Errorf("Coalesce() = %v, want 0", got)
	}
	if got := Coalesce(0, 42); got != 42 {
		t.Errorf("Coalesce() = %v, want 42", got)
	}
}

func TestTry(t *testing.T) {
	successFn := func() (int, error) { return 42, nil }
	errorFn := func() (int, error) { return 0, errors.New("fail") }

	if got := Try(10, successFn); got != 42 {
		t.Errorf("Try(successFn) = %v, want 42", got)
	}

	if got := Try(10, errorFn); got != 10 {
		t.Errorf("Try(errorFn) = %v, want 10", got)
	}
}

func ExampleIf() {
	statusCode := 404
	statusMsg := If(statusCode == 200, "OK", "Not Found")
	fmt.Println(statusMsg)
	// Output: Not Found
}

func ExampleCoalesce() {
	var input1, input2 string
	input3 := "default_value"

	result := Coalesce(input1, input2, input3)
	fmt.Println(result)
	// Output: default_value
}
