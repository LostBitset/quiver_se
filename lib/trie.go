package qse

// A trie that maps sequences of N to values of L
type Trie[N any, L comparable] struct {
	root   TrieNode[N, L]
	leaves map[L]struct{}
}

type TrieNode[N any, L comparable] interface {
	IsTrieLeaf() (is bool)
}

type TrieValueNode[N any, L comparable] struct {
	value    []N
	parents  []TrieValueNode[N, L]
	children []TrieNode[N, L]
}

type TrieLeafNode[L comparable] struct {
	value L
}
