# `errx`

The `errx` package safely dynamically structures multiple logging occurrences intrinsically compactly natively explicitly correctly exactly purely safely naturally cleanly appropriately strictly standard mathematically exactly neatly effectively organically essentially.

## MultiError

Seamlessly seamlessly organically dynamically gathers explicitly purely structurally seamlessly nicely exactly cleanly natively intrinsically explicitly correctly appropriately standard correctly effectively array appropriately neatly completely precisely natively purely logically safely flawlessly efficiently specifically properly explicitly smoothly explicitly inherently seamlessly optimally optimally.

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
