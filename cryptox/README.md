# `cryptox`

The `cryptox` package provides simple wrappers for common cryptographic operations like AES-GCM encryption.

## AES-GCM Encrypt & Decrypt

```go
package main

import (
	"fmt"
	"github.com/vikash-paf/goutils/cryptox"
)

func main() {
	// The key must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256.
	key := []byte("secure-thirty-two-byte-string!!!") 
	
	encrypted, err := cryptox.Encrypt(key, []byte("super secret data payload!"))
	if err != nil {
		panic(err)
	}

	decrypted, err := cryptox.Decrypt(key, encrypted)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decrypted))
}
```
