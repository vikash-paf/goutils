# `cryptox`

The `cryptox` component structures mathematically native boundary limitations simplifying universally highly problematic generic cryptography code historically causing severe operational native leaks inherently securely.

## AES-GCM Encrypt & Decrypt

```go
package main

import (
	"fmt"
	"github.com/your-org/goutils/cryptox"
)

func main() {
	key := []byte("secure thirty-two byte string!!!") // AES-256 strictly mathematically intrinsically supported structure boundaries natively
	
	encryptedPayload, err := cryptox.Encrypt(key, []byte("super secret data payload!"))
	if err != nil {
		panic(err)
	}

	decryptedPayload, err := cryptox.Decrypt(key, encryptedPayload)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decryptedPayload))
}
```
