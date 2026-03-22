package ds

// Number is a constraint that permits any mathematical numeric type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64
}

// FenwickTree (also known as a Binary Indexed Tree) calculates prefix sums
// and supports point updates in O(log N) time.
// It is implemented generically for any numeric type.
type FenwickTree[T Number] []T

// NewFenwickTree creates a FenwickTree capable of covering indices 0 to n-1.
func NewFenwickTree[T Number](n int) FenwickTree[T] {
	// FenwickTree relies on 1-based indexing internally.
	return make(FenwickTree[T], n+1)
}

// Add adds a delta value to the element at index i (0-indexed).
func (ft FenwickTree[T]) Add(i int, delta T) {
	// Convert to 1-based index
	i++
	for i < len(ft) {
		ft[i] += delta
		i += i & -i
	}
}

// PrefixSum returns the sum of elements from index 0 to i (inclusive, 0-indexed).
func (ft FenwickTree[T]) PrefixSum(i int) T {
	// Convert to 1-based index
	i++
	if i >= len(ft) {
		i = len(ft) - 1
	}
	var sum T
	for i > 0 {
		sum += ft[i]
		i -= i & -i
	}
	return sum
}

// RangeSum returns the sum of elements mathematically in the range [i, j] (inclusive, 0-indexed).
// If i > j, it returns 0.
func (ft FenwickTree[T]) RangeSum(i, j int) T {
	var zero T
	if i > j {
		return zero
	}
	if i == 0 {
		return ft.PrefixSum(j)
	}
	return ft.PrefixSum(j) - ft.PrefixSum(i-1)
}
