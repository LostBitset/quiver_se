package qse

type digest_t = []byte

// A dangerous version of hashable that doesn't require comparable
// This is because you cannot cast to an interface containing
// comparable for some reason
type unsafe_relaxed_hashable interface {
	Hash() (digest digest_t)
	Hash32() (fixed_digest uint32)
}

// This is the version of Hashable to use
type Hashable interface {
	unsafe_relaxed_hashable
	comparable
}

type uint32_H struct {
	uint32
}

type Literal[NODE Hashable] struct {
	Value NODE
	Eq    bool
}

type DMT[NODE Hashable, LEAF Hashable] struct {
	trie Trie[Literal[NODE], LEAF, digest_t]
}
