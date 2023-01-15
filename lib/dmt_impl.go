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

func ZeroDigest() (digest []byte) {
	digest = make([]byte, 4)
	return
}

func FixDigest(pseudo_digest []byte, fix_seed byte) (digest []byte) {
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{fix_seed})
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
				ZeroDigest(),
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
	var node *TrieValueNode[Literal[NODE], LEAF, digest_t]
	node = leaf_ptr.parent
correctLoop:
	for {
		sum := ZeroDigest()
	xorChildrenLoop:
		for _, raw_child := range node.children {
			var child *TrieValueNode[Literal[NODE], LEAF, digest_t]
			switch c := raw_child.(type) {
			case TrieLeafNode[Literal[NODE], LEAF, digest_t]:
				continue xorChildrenLoop
			case *TrieLeafNode[Literal[NODE], LEAF, digest_t]:
				continue xorChildrenLoop
			case TrieValueNode[Literal[NODE], LEAF, digest_t]:
				child = &c
			case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
				child = c
			}
			for i, byte_value := range child.meta {
				sum[i] ^= byte_value
			}
		}
		*&node.meta = FixDigest(sum, 0x3E) // Update the hash
		t.SimplifyNode(node) // Try and locally simplify
		node = node.parent
		if node == nil {
			break correctLoop
		}
	}
}

func (t *DMT[NODE, LEAF]) SimplifyNode(node *TrieValueNode[Literal[NODE], LEAF, digest_t]) {}
