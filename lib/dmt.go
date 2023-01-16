package qse

type digest_t = []byte

type hashable interface {
	Hash() (digest digest_t)
	comparable
}

type uint32_H struct {
	uint32
}

type Literal[NODE hashable] struct {
	value NODE
	eq    bool
}

type DMT[NODE hashable, LEAF hashable] struct {
	trie Trie[Literal[NODE], LEAF, digest_t]
}
