package qse

import (
	"encoding/binary"
	"hash/fnv"
)

func (lit Literal[NODE]) Hash() (digest []byte) {
	hasher := fnv.New32a()
	hasher.Write(lit.value.Hash())
	if !lit.eq {
		hasher.Write([]byte{0xDD})
	}
	digest = hasher.Sum([]byte{})
	return
}

func FixDigest(digest digest_t) (digest_fixed digest_fixed_t) {
	digest_fixed = binary.LittleEndian.Uint32(digest)
	return
}

func (lit Literal[NODE]) Merkleify() (mlit MerkleLiteral[NODE]) {
	mlit = MerkleLiteral[NODE]{
		lit,
		FixDigest(lit.Hash()),
	}
	return
}

func BufferingLiteral[NODE hashable](value NODE) (lit MerkleLiteral[NODE]) {
	lit = Literal[NODE]{value, true}.Merkleify()
	return
}

func InvertingLiteral[NODE hashable](value NODE) (lit MerkleLiteral[NODE]) {
	lit = Literal[NODE]{value, false}.Merkleify()
	return
}

func NewDMT[NODE hashable, LEAF comparable]() (t DMT[NODE, LEAF]) {
	t = DMT[NODE, LEAF]{
		NewTrie[MerkleLiteral[NODE], LEAF](),
	}
	return
}

func (t *DMT[NODE, LEAF]) Insert(seq map[MerkleLiteral[NODE]]struct{}, leaf LEAF) (leaf_ptr *TrieLeafNode[MerkleLiteral[NODE], LEAF]) {
	leaf_ptr = t.trie.Insert(seq, leaf)
	t.UpdateHashes(leaf_ptr)
	return
}

func (t *DMT[NODE, LEAF]) UpdateHashes(leaf_ptr *TrieLeafNode[MerkleLiteral[NODE], LEAF]) {
}
