package poolx

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTypedPool(t *testing.T) {
	// Must initialize by storing pointers to properly modify and save attributes in pool
	p := NewTypedPool[*bytes.Buffer](func() *bytes.Buffer {
		return new(bytes.Buffer) // Only allocates strictly dynamically if empty
	})

	buf := p.Get()
	if buf.Len() != 0 {
		t.Error("Expected an intrinsically empty buffer array slice wrapper.")
	}

	buf.WriteString("hello gravity")

	// Extremely common edge rule: MUST reset explicitly BEFORE placing it back in the pool structurally.
	buf.Reset()
	p.Put(buf)

	buf2 := p.Get()

	// Structurally prove that buffer string matches cleanly memory sizes (in other words, 0 characters length!)
	if buf2.Len() != 0 {
		t.Error("Failed to provide a safe, scrubbed buffer struct slice size natively.")
	}

	// Make sure we actually can rewrite locally.
	buf2.WriteString("secondary logic")
	if buf2.String() != "secondary logic" {
		t.Error("Failed buffer translation structure logic.")
	}
}

func ExampleTypedPool() {
	// A strictly designated string slice pool
	pool := NewTypedPool[[]string](func() []string {
		return make([]string, 0, 100) // Pre-allocate precisely structural size strings lengths mathematically
	})

	slice := pool.Get()
	slice = append(slice, "hello")
	slice = append(slice, "world")

	fmt.Println("Initial slice assignment length logically:", len(slice))

	// Reusing safely! Must reset arrays down to exactly 0 to guarantee slice capacities dynamically!
	slice = slice[:0]
	pool.Put(slice)

	// Output:
	// Initial slice assignment length logically: 2
}
