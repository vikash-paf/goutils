package ds

// UnionFind is a generic disjoint-set data structure.
// It tracks a set of elements partitioned into a number of disjoint (non-overlapping) subsets.
// This implementation uses a map to support arbitrary comparable elements, offering path
// compression and union-by-rank for near O(1) time complexity.
type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	count  int // Number of connected components
}

// NewUnionFind creates a new, empty UnionFind data structure.
func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		count:  0,
	}
}

// Add inserts a new element into the disjoint set as a standalone component.
// If the element already exists, it does nothing.
func (uf *UnionFind[T]) Add(item T) {
	if _, exists := uf.parent[item]; !exists {
		uf.parent[item] = item
		uf.rank[item] = 1
		uf.count++
	}
}

// Find returns the representative (root) element of the set containing item.
// It uses path compression to keep the tree shallow.
func (uf *UnionFind[T]) Find(item T) (T, bool) {
	var zero T
	p, exists := uf.parent[item]
	if !exists {
		return zero, false
	}

	if p == item {
		return p, true
	}

	// Path compression: make every node on the path point directly to the root.
	root, _ := uf.Find(p)
	uf.parent[item] = root
	return root, true
}

// Union merges the sets containing item1 and item2.
// It uses union-by-rank. Returns true if they were merged, false if they were
// already in the same set or if one of the items does not exist.
func (uf *UnionFind[T]) Union(item1, item2 T) bool {
	root1, ok1 := uf.Find(item1)
	root2, ok2 := uf.Find(item2)

	if !ok1 || !ok2 || root1 == root2 {
		return false
	}

	if uf.rank[root1] > uf.rank[root2] {
		uf.parent[root2] = root1
	} else if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
	} else {
		uf.parent[root2] = root1
		uf.rank[root1]++
	}

	uf.count--
	return true
}

// Connected returns true if item1 and item2 belong to the same set.
func (uf *UnionFind[T]) Connected(item1, item2 T) bool {
	root1, ok1 := uf.Find(item1)
	root2, ok2 := uf.Find(item2)
	return ok1 && ok2 && root1 == root2
}

// Count returns the current number of independent disjoint sets.
func (uf *UnionFind[T]) Count() int {
	return uf.count
}
