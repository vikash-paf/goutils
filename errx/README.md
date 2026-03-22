# `errx`

The `errx` package provides utilities for managing and aggregating multiple errors.

## MultiError

Collects multiple errors into a single composite error that satisfies the `error` interface.

```go
package main

import (
	"errors"
	"fmt"
	"github.com/your-org/goutils/errx"
)

func main() {
	m := &errx.MultiError{}
	m.Append(errors.New("db disconnect timeout"))
	m.Append(errors.New("auth validation rejected"))
	
	if err := m.AsError(); err != nil {
		fmt.Println(err.Error())
	}
}
```
