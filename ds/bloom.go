// Package ds provides common data structures.
package ds

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter is a space-efficient probabilistic data structure that is used to test
// whether an element is a member of a set. False positive matches are possible,
// but false negatives are not.
type BloomFilter struct {
	bitset []uint64
	k      uint // Number of hash functions
	m      uint // Number of bits in the bitset
}

// NewBloomFilter creates a new Bloom filter tailored to store 'n' elements
// with a target false positive probability 'p'.
func NewBloomFilter(n uint, p float64) *BloomFilter {
	// m = ceil((n * log(p)) / log(1 / pow(2, log(2))))
	m := uint(math.Ceil(float64(n) * math.Log(p) / math.Log(1.0/math.Pow(2, math.Log(2)))))
	// k = round((m / n) * log(2))
	k := uint(math.Round(float64(m) / float64(n) * math.Log(2)))

	return &BloomFilter{
		bitset: make([]uint64, (m+63)/64),
		k:      k,
		m:      m,
	}
}

// Add inserts data into the Bloom filter.
func (bf *BloomFilter) Add(data []byte) {
	h1, h2 := hashData(data)
	for i := uint(0); i < bf.k; i++ {
		bitIndex := (h1 + uint64(i)*h2) % uint64(bf.m)
		bf.bitset[bitIndex/64] |= 1 << (bitIndex % 64)
	}
}

// AddString inserts a string into the Bloom filter.
func (bf *BloomFilter) AddString(s string) {
	bf.Add([]byte(s))
}

// Contains checks whether the data might be in the set.
// It returns true if the data might be in the set (with a false positive probability),
// or false if the data is definitely not in the set.
func (bf *BloomFilter) Contains(data []byte) bool {
	h1, h2 := hashData(data)
	for i := uint(0); i < bf.k; i++ {
		bitIndex := (h1 + uint64(i)*h2) % uint64(bf.m)
		if bf.bitset[bitIndex/64]&(1<<(bitIndex%64)) == 0 {
			return false
		}
	}
	return true
}

// ContainsString checks whether a string might be in the set.
func (bf *BloomFilter) ContainsString(s string) bool {
	return bf.Contains([]byte(s))
}

// hashData returns two independent 64-bit hashes using FNV-1a.
// We use Kirsch-Mitzenmacher optimization to simulate k hash functions from two.
func hashData(data []byte) (uint64, uint64) {
	h := fnv.New64a()
	_, _ = h.Write(data)
	hash1 := h.Sum64()

	// Slightly modify the hash for the second one
	h2 := fnv.New64a()
	_, _ = h2.Write(append(data, 1))
	hash2 := h2.Sum64()

	return hash1, hash2
}
