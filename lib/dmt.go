package qse

type digest_t = []byte
type digest_fixed_t = uint32

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
	subtree_hash digest_fixed_t
}

type DMT[NODE hashable, LEAF comparable] struct {
	trie Trie[MerkleLiteral[NODE], LEAF]
}
