package qse

type digest_t = []byte

type hashable interface {
	Hash() (digest digest_t)
	comparable
}

type Literal[NODE hashable] struct {
	value NODE
	eq bool
}

type MerkleLiteral[NODE hashable] struct {
	Literal[NODE]
	subtree_hash digest_t
}

type DMT[NODE hashable, LEAF comparable] struct {
	trie Trie[MerkleLiteral[NODE], LEAF]
}
