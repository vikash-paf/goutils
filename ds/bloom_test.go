package ds

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	// Create a Bloom filter for 1000 items with a 1% false positive rate
	bf := NewBloomFilter(1000, 0.01)

	// Add some items
	items := []string{"apple", "banana", "orange", "grape"}
	for _, item := range items {
		bf.AddString(item)
	}

	// Verify added items are found (no false negatives)
	for _, item := range items {
		if !bf.ContainsString(item) {
			t.Errorf("Expected to contain %s, but didn't", item)
		}
	}

	// Verify items not added are likely not found
	notAdded := []string{"pear", "mango", "pineapple", "watermelon"}
	falsePositives := 0
	for _, item := range notAdded {
		if bf.ContainsString(item) {
			falsePositives++
		}
	}

	// Given a 1% false positive rate and 4 checks, we expect very few false positives.
	if falsePositives > 1 {
		t.Errorf("Got too many false positives: %d", falsePositives)
	}
}

func ExampleBloomFilter() {
	bf := NewBloomFilter(10000, 0.01)
	bf.AddString("user_123")

	if bf.ContainsString("user_123") {
		fmt.Println("user_123 might be in the set")
	}

	if !bf.ContainsString("user_456") {
		fmt.Println("user_456 is definitely not in the set")
	}
	// Output:
	// user_123 might be in the set
	// user_456 is definitely not in the set
}
