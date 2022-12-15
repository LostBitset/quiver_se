package qse

type Trie[N any, L comparable] struct {
	root   TrieNode[N, L]
	leaves map[L]struct{}
}

type TrieNode[N any, L any] struct {
	value    []N
	parents  []TrieNode[N, L]
	children []NT
}
