package str

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"kitten", "sitting", 3},
		{"flaw", "lawn", 2},
		{"goutils", "goutils", 0},
		{"", "abc", 3},
		{"café", "cafe", 1},
	}

	for _, tt := range tests {
		if got := Levenshtein(tt.a, tt.b); got != tt.want {
			t.Errorf("Levenshtein(%s, %s) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSecureRandom(t *testing.T) {
	s1, err := SecureRandom(32)
	if err != nil {
		t.Fatalf("SecureRandom failed: %v", err)
	}
	if len(s1) != 32 {
		t.Errorf("Expected length 32, got %d", len(s1))
	}

	s2, _ := SecureRandom(32)
	if s1 == s2 {
		t.Error("SecureRandom produced identical strings")
	}
}
