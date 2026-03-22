package ds

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie[int]()
	trie.Insert("apple", 1)
	trie.Insert("app", 2)
	trie.Insert("applied", 3)

	val, ok := trie.Search("apple")
	if !ok || val != 1 {
		t.Errorf("Search(apple) = %v, %v, want 1, true", val, ok)
	}

	val, ok = trie.Search("app")
	if !ok || val != 2 {
		t.Errorf("Search(app) = %v, %v, want 2, true", val, ok)
	}

	if !trie.StartsWith("appl") {
		t.Error("StartsWith(appl) should be true")
	}

	results := trie.PrefixSearch("app")
	if len(results) != 3 {
		t.Errorf("PrefixSearch(app) length = %d, want 3", len(results))
	}
}
