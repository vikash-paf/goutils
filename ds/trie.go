package ds

// TrieNode represents a node in the Trie.
type TrieNode[V any] struct {
	children map[rune]*TrieNode[V]
	isEnd    bool
	value    V
}

// Trie is a generic prefix tree.
type Trie[V any] struct {
	root *TrieNode[V]
}

// NewTrie creates a new Trie.
func NewTrie[V any]() *Trie[V] {
	return &Trie[V]{
		root: &TrieNode[V]{
			children: make(map[rune]*TrieNode[V]),
		},
	}
}

// Insert adds a word and its associated value to the Trie.
func (t *Trie[V]) Insert(word string, value V) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &TrieNode[V]{
				children: make(map[rune]*TrieNode[V]),
			}
		}
		node = node.children[r]
	}
	node.isEnd = true
	node.value = value
}

// Search looks for a word in the Trie and returns its value and existence.
func (t *Trie[V]) Search(word string) (V, bool) {
	node := t.root
	for _, r := range word {
		if next, ok := node.children[r]; ok {
			node = next
		} else {
			var zero V
			return zero, false
		}
	}
	if node.isEnd {
		return node.value, true
	}
	var zero V
	return zero, false
}

// StartsWith returns true if there is any word in the Trie that starts with the given prefix.
func (t *Trie[V]) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		if next, ok := node.children[r]; ok {
			node = next
		} else {
			return false
		}
	}
	return true
}

// PrefixSearch returns all values whose keys start with the given prefix.
func (t *Trie[V]) PrefixSearch(prefix string) []V {
	node := t.root
	for _, r := range prefix {
		if next, ok := node.children[r]; ok {
			node = next
		} else {
			return nil
		}
	}

	var results []V
	var collect func(*TrieNode[V])
	collect = func(n *TrieNode[V]) {
		if n.isEnd {
			results = append(results, n.value)
		}
		for _, child := range n.children {
			collect(child)
		}
	}
	collect(node)
	return results
}
