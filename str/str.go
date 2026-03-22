// Package str provides utility functions for strings.
package str

import (
	"strings"
	"unicode"
)

// IsBlank returns true if the string is empty or contains only whitespace characters.
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// Reverse reverses the characters in a string. It works correctly with unicode runes.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Truncate shortens a string to the given length and appends the omission string (e.g., "...").
// If the original string is already shorter than or equal to the specified length, it is returned unchanged.
func Truncate(s string, length int, omission string) string {
	if length <= 0 {
		return ""
	}

	runes := []rune(s)
	if len(runes) <= length {
		return s
	}

	omissionRunes := []rune(omission)
	if length <= len(omissionRunes) {
		return string(omissionRunes[:length])
	}

	return string(runes[:length-len(omissionRunes)]) + omission
}

// ToCamelCase converts a string to camelCase.
// It removes non-alphanumeric characters and capitalizes words.
func ToCamelCase(s string) string {
	var builder strings.Builder
	var nextUpper bool
	var isFirstWord = true

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			nextUpper = true
			continue
		}

		if isFirstWord {
			builder.WriteRune(unicode.ToLower(r))
			isFirstWord = false
		} else if nextUpper {
			builder.WriteRune(unicode.ToUpper(r))
			nextUpper = false
		} else {
			builder.WriteRune(unicode.ToLower(r))
		}
	}

	return builder.String()
}

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(s string) string {
	var builder strings.Builder
	var previousLower bool

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if builder.Len() > 0 && !strings.HasSuffix(builder.String(), "_") {
				builder.WriteRune('_')
			}
			previousLower = false
			continue
		}

		if unicode.IsUpper(r) {
			if previousLower && builder.Len() > 0 {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(r))
			previousLower = false
		} else {
			builder.WriteRune(unicode.ToLower(r))
			previousLower = true
		}
	}

	// Remove trailing underscore if present
	return strings.TrimSuffix(builder.String(), "_")
}
