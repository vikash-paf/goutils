package str

import (
	"fmt"
	"testing"
)

func TestIsBlank(t *testing.T) {
	if !IsBlank("") {
		t.Error("IsBlank(` `) should be true")
	}
	if !IsBlank("   ") {
		t.Error("IsBlank(`   `) should be true")
	}
	if IsBlank("a") {
		t.Error("IsBlank(`a`) should be false")
	}
}

func TestReverse(t *testing.T) {
	got := Reverse("hello")
	want := "olleh"
	if got != want {
		t.Errorf("Reverse() = %v, want %v", got, want)
	}

	gotRune := Reverse("こんにちは")
	wantRune := "はちにんこ"
	if gotRune != wantRune {
		t.Errorf("Reverse() = %v, want %v", gotRune, wantRune)
	}
}

func TestTruncate(t *testing.T) {
	got := Truncate("hello world", 8, "...")
	want := "hello..."
	if got != want {
		t.Errorf("Truncate() = %v, want %v", got, want)
	}

	got2 := Truncate("hi", 5, "...")
	want2 := "hi"
	if got2 != want2 {
		t.Errorf("Truncate() = %v, want %v", got2, want2)
	}
}

func TestToCamelCase(t *testing.T) {
	got := ToCamelCase("Hello world_Test-case")
	want := "helloWorldTestCase"
	if got != want {
		t.Errorf("ToCamelCase() = %v, want %v", got, want)
	}
}

func TestToSnakeCase(t *testing.T) {
	got := ToSnakeCase("HelloWorldTestCase")
	want := "hello_world_test_case"
	if got != want {
		t.Errorf("ToSnakeCase() = %v, want %v", got, want)
	}

	got2 := ToSnakeCase("hello_world")
	want2 := "hello_world"
	if got2 != want2 {
		t.Errorf("ToSnakeCase() = %v, want %v", got2, want2)
	}
}

func ExampleToSnakeCase() {
	s := ToSnakeCase("CamelCaseToSnakeCase")
	fmt.Println(s)
	// Output: camel_case_to_snake_case
}

func ExampleToCamelCase() {
	s := ToCamelCase("snake_case_to_camel_case")
	fmt.Println(s)
	// Output: snakeCaseToCamelCase
}
