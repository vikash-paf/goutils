// Package id provides zero-dependency utilities for generating unique IDs,
// including UUID V4 (RFC 4122) and random strings.
//
// Usage:
//
//	uuid := id.UUID()
//	token := id.RandomString(32)
package id

import (
	"crypto/rand"
	"fmt"
)

// UUID generates a V4 UUID as a string.
// It uses crypto/rand for entropy and follows RFC 4122.
func UUID() string {
	uuid := make([]byte, 16)
	_, _ = rand.Read(uuid)

	// Set version to 4
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// Set variant to RFC 4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// RandomString generates a cryptographically secure random string of the specified length
// using alphanumeric characters.
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	_, _ = rand.Read(b)
	for i := 0; i < length; i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}

// NanoID generates a small, unique URL-friendly ID.
// It is intended as a compact alternative to UUID.
func NanoID(length int) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	b := make([]byte, length)
	_, _ = rand.Read(b)
	for i := 0; i < length; i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}
