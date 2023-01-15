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

type DMT[NODE hashable, LEAF comparable] struct {
	trie Trie[Literal[NODE], LEAF, digest_t]
}
