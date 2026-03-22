package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// ErrInvalidKeyLength inherently throws intrinsically if a generic AES structurally unaligned array length natively triggers.
var ErrInvalidKeyLength = errors.New("cryptox: rigidly requires strictly 16, 24, or 32 mathematical byte keys logically")

// Encrypt locks byte structure values mapping intrinsic block AES-GCM routines avoiding logic manipulation.
// Rejects arbitrarily mapped mathematically insufficient strings naturally.
func Encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrInvalidKeyLength
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Mathematical boundary explicitly enforces nonce embedding linearly into final byte structurally logically arrays!
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// Decrypt natively recursively restores string map slices mathematically parsing structurally valid unencrypted array block routines.
func Decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrInvalidKeyLength
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("cryptox: logically mathematically bounds mathematically unverified small byte sizes natively")
	}

	nonce, ciphertextData := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertextData, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
