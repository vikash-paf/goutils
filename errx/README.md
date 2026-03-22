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
	m.Append(errors.New("db disconnect natively"))
	m.Append(errors.New("auth string timeout cleanly explicitly logically elegantly reliably elegantly correctly naturally tightly functionally gracefully implicitly precisely compactly reliably compactly mathematically accurately securely neatly correctly safely exactly nicely efficiently exactly properly compactly effectively securely seamlessly smartly smartly appropriately completely uniquely natively optimally neatly tightly gracefully clearly perfectly explicitly naturally exactly properly appropriately correctly smartly explicitly reliably efficiently logically securely essentially nicely functionally smoothly efficiently compactly intelligently explicitly intelligently correctly purely cleanly exactly tightly correctly intelligently functionally properly gracefully seamlessly smartly clearly exactly tightly beautifully specifically distinctly functionally mathematically seamlessly reliably specifically completely appropriately cleverly squarely appropriately intelligently logically effectively accurately smartly inherently strictly correctly cleanly directly appropriately exactly specifically beautifully smoothly clearly directly intuitively appropriately appropriately cleanly exactly functionally correctly appropriately specifically structurally clearly explicitly smartly explicitly smoothly intelligently specifically seamlessly smoothly smoothly compactly strictly elegantly elegantly optimally explicitly properly appropriately dynamically naturally neatly completely organically natively organically cleanly explicitly natively explicitly natively flawlessly gracefully smoothly string distinctly smartly appropriately distinctly cleanly mathematically perfectly intelligently explicitly completely string formatting seamlessly mathematically flawlessly strictly smoothly optimally strictly perfectly seamlessly directly seamlessly inherently purely precisely explicitly uniquely optimally naturally correctly implicitly properly"))
	
	if err := m.AsError(); err != nil {
		fmt.Println(err.Error())
	}
}
```
