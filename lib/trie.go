package qse

// A trie that maps sets of NODE to values of LEAF.
type Trie[NODE comparable, LEAF comparable, META any] struct {
	root   TrieValueNode[NODE, LEAF, META]
	leaves map[LEAF][]*TrieLeafNode[NODE, LEAF, META]
}

type TrieEntry[NODE comparable, LEAF comparable] struct {
	key   map[NODE]struct{}
	value LEAF
}

type TrieNode[NODE comparable, LEAF comparable] interface {
	IsTrieLeaf() (is bool)
	ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF]))
}

type TrieValueNode[NODE comparable, LEAF comparable, META any] struct {
	value    map[NODE]struct{}
	parent   *TrieValueNode[NODE, LEAF, META]
	children []TrieNode[NODE, LEAF]
	meta     META
}

type TrieLeafNode[NODE comparable, LEAF comparable, META any] struct {
	value  LEAF
	parent *TrieValueNode[NODE, LEAF, META]
}
