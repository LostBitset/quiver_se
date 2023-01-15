package qse

import (
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

func BufferingLiteral[NODE hashable](value NODE) (lit Literal[NODE]) {
	lit = Literal[NODE]{value, true}
	return
}

func InvertingLiteral[NODE hashable](value NODE) (lit Literal[NODE]) {
	lit = Literal[NODE]{value, false}
	return
}

func NewDMT[NODE hashable, LEAF comparable]() (t DMT[NODE, LEAF]) {
	t = DMT[NODE, LEAF]{
		Trie[Literal[NODE], LEAF, digest_t]{
			TrieValueNode[Literal[NODE], LEAF, digest_t]{
				map[Literal[NODE]]struct{}{},
				nil,
				[]TrieNode[Literal[NODE], LEAF]{},
				make([]byte, 4),
			},
			map[LEAF][]*TrieLeafNode[Literal[NODE], LEAF, []byte]{},
		},
	}
	return
}

func (t *DMT[NODE, LEAF]) Insert(
	seq map[Literal[NODE]]struct{}, leaf LEAF,
) (
	leaf_ptr *TrieLeafNode[Literal[NODE], LEAF, digest_t],
) {
	leaf_ptr = t.trie.Insert(seq, leaf)
	t.UpdateHashes(leaf_ptr)
	return
}

func (t *DMT[NODE, LEAF]) UpdateHashes(leaf_ptr *TrieLeafNode[Literal[NODE], LEAF, digest_t]) {
}
