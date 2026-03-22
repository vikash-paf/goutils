# `poolx`

The `poolx` package provides robust generic object memory pooling functions to limit arbitrary memory allocations and enhance garbage collection metrics intrinsically.

## TypedPool

A highly optimized memory array generic wrapper enforcing strong mathematical type definitions over the notoriously frustrating Go `sync.Pool`.

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/your-org/goutils/poolx"
)

func main() {
	// A dedicated string pool array returning exact []string structures safely without `any` casting.
	p := poolx.NewTypedPool[[]string](func() []string {
		return make([]string, 0, 32)
	})

	allocatedSlice := p.Get()
	allocatedSlice = append(allocatedSlice, "safely pooled item")
	
	// Reset mathematically BEFORE safely returning
	allocatedSlice = allocatedSlice[:0]
	p.Put(allocatedSlice)
}
```
