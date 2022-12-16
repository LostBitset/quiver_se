package qse

// A trie that maps sets of N to values of L
type Trie[N comparable, L comparable] struct {
	root   TrieValueNode[N, L]
	leaves map[L]struct{}
}

type TrieNode[N comparable, L comparable] interface {
	IsTrieLeaf() (is bool)
}

type TrieValueNode[N comparable, L comparable] struct {
	value    map[N]struct{}
	parents  []TrieValueNode[N, L]
	children []TrieNode[N, L]
}

type TrieLeafNode[L comparable] struct {
	value L
}
