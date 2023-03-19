package qse

type digest_t = []byte

// A dangerous version of hashable that doesn't require comparable
// This is because you cannot cast to an interface containing
// comparable for some reason
type unsafe_relaxed_hashable interface {
	Hash() (digest digest_t)
	Hash32() (fixed_digest uint32)
}

// This is the version of hashable to use
type hashable interface {
	unsafe_relaxed_hashable
	comparable
}

type uint32_H struct {
	uint32
}

type Literal[NODE hashable] struct {
	Value NODE
	Eq    bool
}

type DMT[NODE hashable, LEAF hashable] struct {
	trie Trie[Literal[NODE], LEAF, digest_t]
}
