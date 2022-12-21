package qse

// A merkle trie that maps sets of NODE to values of LEAF.
type MerkleTrie[NODE comparable, LEAF comparable] struct {
	root   MTrieValueNode[NODE, LEAF]
	leaves map[LEAF][]*MTrieLeafNode[NODE, LEAF]
}

type MTrieValueNode[NODE comparable, LEAF comparable] struct {
	value              map[NODE]struct{}
	parent             *TrieValueNode[NODE, LEAF]
	children           []TrieNode[NODE, LEAF]
	child_subtree_hash int
}

type MTrieLeafNode[NODE comparable, LEAF comparable] struct {
	value  LEAF
	parent *TrieValueNode[NODE, LEAF]
}
