// Package encodingx provides high-level utilities for encoding and decoding data,
// specifically focusing on JSON and map transformations.
//
// Usage:
//
//	jsonStr := encodingx.MustMarshalJSON(config)
//	if m, err := encodingx.ToMap(structObj); err == nil {
//	    // access map
//	}
package encodingx

import (
	"encoding/json"
	"fmt"
)

// MustMarshalJSON marshals the provided value to JSON and returns the resulting string.
// It panics if marshaling fails. This is intended for use with known structures or in tests.
func MustMarshalJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("MustMarshalJSON failed: %v", err))
	}
	return string(b)
}

// ToMap converts any struct or value into a map[string]any by marshaling it to JSON
// and then unmarshaling it back into a map.
func ToMap(v any) (map[string]any, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrettyPrint returns a formatted, indented JSON string for the given value.
func PrettyPrint(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error pretty-printing: %v", err)
	}
	return string(b)
}
