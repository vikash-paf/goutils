package str

import (
	"math"
	"testing"
)

const epsilon = 0.0001

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestJaro(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   float64
	}{
		{"", "", 1.0},
		{"a", "", 0.0},
		{"", "a", 0.0},
		{"martha", "marhta", 0.944444},
		{"dwayne", "duane", 0.822222},
		{"dixon", "dicksonx", 0.766666},
		{"MARTHA", "MARHTA", 0.944444},
	}

	for _, tt := range tests {
		t.Run(tt.s1+"_"+tt.s2, func(t *testing.T) {
			got := Jaro(tt.s1, tt.s2)
			if !almostEqual(got, tt.want) {
				t.Errorf("Jaro(%q, %q) = %v; want ≈ %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestJaroWinkler(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   float64
	}{
		{"", "", 1.0},
		{"martha", "marhta", 0.961111},
		{"dwayne", "duane", 0.840000},
		{"dixon", "dicksonx", 0.813333},
	}

	for _, tt := range tests {
		t.Run(tt.s1+"_"+tt.s2, func(t *testing.T) {
			got := JaroWinkler(tt.s1, tt.s2)
			if !almostEqual(got, tt.want) {
				t.Errorf("JaroWinkler(%q, %q) = %v; want ≈ %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func ExampleJaroWinkler() {
	sim := JaroWinkler("martha", "marhta")
	// Thresholding helps handle typos effectively
	if sim > 0.9 {
		println("High similarity")
	}
	// Output: High similarity
}
