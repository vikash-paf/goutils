package cryptox_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/cryptox"
)

func ExampleEncrypt() {
	key := []byte("secure-thirty-two-byte-string!!!")
	data := []byte("super secret payload")

	encrypted, err := cryptox.Encrypt(key, data)
	if err != nil {
		fmt.Println("Encryption failed:", err)
		return
	}

	decrypted, err := cryptox.Decrypt(key, encrypted)
	if err != nil {
		fmt.Println("Decryption failed:", err)
		return
	}

	fmt.Println(string(decrypted))
	// Output: super secret payload
}
