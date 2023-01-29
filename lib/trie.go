package qse

// A trie that maps sets of NODE to values of LEAF.
type Trie[NODE hashable, LEAF comparable, META any] struct {
	root   TrieValueNode[NODE, LEAF, META]
	leaves map[LEAF][]*TrieLeafNode[NODE, LEAF, META]
}

type TrieEntry[NODE hashable, LEAF comparable] struct {
	key   map[NODE]struct{}
	value LEAF
}

type TrieNode[NODE hashable, LEAF comparable] interface {
	IsTrieLeaf() (is bool)
	ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF]))
}

type TrieValueNode[NODE hashable, LEAF comparable, META any] struct {
	value    map[NODE]struct{}
	parent   *TrieValueNode[NODE, LEAF, META]
	children []TrieNode[NODE, LEAF]
	meta     META
}

type TrieLeafNode[NODE hashable, LEAF comparable, META any] struct {
	value  LEAF
	parent *TrieValueNode[NODE, LEAF, META]
}
