package str

import (
	"crypto/rand"
	"math/big"
)

// Levenshtein calculates the Levenshtein distance between two strings.
// It returns the minimum number of single-character edits (insertions, deletions or substitutions)
// required to change one word into the other.
func Levenshtein(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)
	n, m := len(runesA), len(runesB)

	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, m+1)
		matrix[i][0] = i
	}
	for j := 0; j <= m; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cost := 1
			if runesA[i-1] == runesB[j-1] {
				cost = 0
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[n][m]
}

// SecureRandom generates a cryptographically secure random alphanumeric string of the given length.
func SecureRandom(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}
