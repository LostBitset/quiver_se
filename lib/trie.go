package qse

// A trie that maps sets of NODE to values of LEAF.
type Trie[NODE comparable, LEAF comparable] struct {
	root   TrieValueNode[NODE, LEAF]
	leaves map[LEAF][]*TrieLeafNode[NODE, LEAF]
}

type TrieEntry[NODE comparable, LEAF comparable] struct {
	key   map[NODE]struct{}
	value LEAF
}

type TrieNode[NODE comparable, LEAF comparable] interface {
	IsTrieLeaf() (is bool)
	ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF]))
}

type TrieValueNode[NODE comparable, LEAF comparable] struct {
	value    map[NODE]struct{}
	parent   *TrieValueNode[NODE, LEAF]
	children []TrieNode[NODE, LEAF]
}

type TrieLeafNode[NODE comparable, LEAF comparable] struct {
	value  LEAF
	parent *TrieValueNode[NODE, LEAF]
}
