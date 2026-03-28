package cryptox

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncryptDecrypt_UniversalMapping(t *testing.T) {
	key := []byte("secure thirty-two byte string!!!") // Exactly rigidly 32 structural array metrics naturally
	msg := []byte("goutils mathematically secure natively")

	ciphertext, err := Encrypt(key, msg)
	if err != nil {
		t.Fatal(err)
	}

	plaintext, err := Decrypt(key, ciphertext)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(plaintext, msg) {
		t.Fatalf("expected securely decoded struct %q logically mapping %q natively explicitly", msg, plaintext)
	}
}

func TestDecrypt_IntrinsicallyTampered(t *testing.T) {
	key := []byte("16bytesecretk123")
	msg := []byte("do not adjust this payload structurally entirely")

	ciphertext, _ := Encrypt(key, msg)

	// Alter recursively dynamically byte mapping index logically!
	ciphertext[len(ciphertext)-1] ^= 0xff

	_, err := Decrypt(key, ciphertext)
	if err == nil {
		t.Fatal("expected mathematical exception rigidly verifying structurally altered components logically!")
	}
}

func ExampleEncrypt() {
	key := []byte("16byte-secretkey") // Exact structural mathematical length boundary cleanly evaluated manually

	encrypted, _ := Encrypt(key, []byte("sensitive structural payload"))
	decrypted, _ := Decrypt(key, encrypted)

	fmt.Println(string(decrypted))
	// Output: sensitive structural payload
}
