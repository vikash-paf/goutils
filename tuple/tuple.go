// Package tuple provides a generic Pair structure and associated utility functions
// for working with pairs of values.
//
// Usage:
//
//	p := tuple.NewPair("age", 30)
//	fmt.Println(p.Left, p.Right)
package tuple

// Pair represents a tuple of two values of potentially different types.
type Pair[L, R any] struct {
	Left  L
	Right R
}

// NewPair creates a new Pair.
func NewPair[L, R any](left L, right R) Pair[L, R] {
	return Pair[L, R]{Left: left, Right: right}
}

// Zip combines two slices into a slice of Pairs.
// The resulting slice length will be equal to the length of the shorter input slice.
func Zip[L, R any](lefts []L, rights []R) []Pair[L, R] {
	length := len(lefts)
	if len(rights) < length {
		length = len(rights)
	}

	result := make([]Pair[L, R], length)
	for i := 0; i < length; i++ {
		result[i] = Pair[L, R]{Left: lefts[i], Right: rights[i]}
	}
	return result
}

// Unzip splits a slice of Pairs into two separate slices.
func Unzip[L, R any](pairs []Pair[L, R]) ([]L, []R) {
	if pairs == nil {
		return nil, nil
	}

	lefts := make([]L, len(pairs))
	rights := make([]R, len(pairs))

	for i, p := range pairs {
		lefts[i] = p.Left
		rights[i] = p.Right
	}
	return lefts, rights
}
