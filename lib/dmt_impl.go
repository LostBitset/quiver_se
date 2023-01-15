package qse

import (
	"bytes"
	"hash/fnv"
)

func (lit Literal[NODE]) Hash() (digest []byte) {
	digest = lit.value.Hash()
	if !lit.eq {
		digest = WrapInvert(digest)
	}
	return
}

func WrapInvert(pseudo_digest []byte) (digest []byte) {
	digest = FixDigest(pseudo_digest, 0xDD)
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

func DigestEqual(a digest_t, b digest_t) (equal bool) {
	equal = bytes.Equal(a, b)
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

func (lit Literal[NODE]) Invert() (inverted Literal[NODE]) {
	inverted = Literal[NODE]{
		lit.value,
		!lit.eq,
	}
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

func (t *DMT[NODE, LEAF]) SimplifyNode(node *TrieValueNode[Literal[NODE], LEAF, digest_t]) {
	edge_indices := make([]int, 0)
	edge_values := make([]map[Literal[NODE]]struct{}, 0)
getChildEdgesLoop:
	for i, child := range node.children {
		switch c := child.(type) {
		case TrieLeafNode[Literal[NODE], LEAF, digest_t]:
			continue getChildEdgesLoop
		case *TrieLeafNode[Literal[NODE], LEAF, digest_t]:
			continue getChildEdgesLoop
		case TrieValueNode[Literal[NODE], LEAF, digest_t]:
			edge_indices = append(edge_indices, i)
			edge_values = append(edge_values, c.value)
		case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
			edge_indices = append(edge_indices, i)
			edge_values = append(edge_values, c.value)
		}
	}
	unwanted_children := make([]int, 0)
	for buf_ii := range edge_indices {
		for inv_ii := range edge_indices {
			if IsInvertedEdge(edge_values[buf_ii], edge_values[inv_ii]) {
				var buf_subtrie_hash digest_t
				var inv_subtrie_hash digest_t
				switch c := node.children[edge_indices[buf_ii]].(type) {
				case TrieValueNode[Literal[NODE], LEAF, digest_t]:
					buf_subtrie_hash = c.meta
				case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
					buf_subtrie_hash = c.meta
				}
				switch c := node.children[edge_indices[inv_ii]].(type) {
				case TrieValueNode[Literal[NODE], LEAF, digest_t]:
					inv_subtrie_hash = c.meta
				case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
					inv_subtrie_hash = c.meta
				}
				if DigestEqual(buf_subtrie_hash, inv_subtrie_hash) {
					unwanted_i := edge_indices[inv_ii]
					unwanted_children = append(unwanted_children, unwanted_i)
					// TODO splice out the buf node
				}
			}
		}
	}
	for _, child := range unwanted_children {
		// Drop the child
		copy(node.children[child:], node.children[(child+1):])
		node.children = node.children[:(len(node.children)-1)]
	}
}

func IsInvertedEdge[NODE hashable](maybe_buf map[Literal[NODE]]struct{}, maybe_inv map[Literal[NODE]]struct{}) (match bool) {
	for key := range maybe_buf {
		inverted := key.Invert()
		if _, ok := maybe_inv[inverted]; !ok {
			match = false
			return
		}
	}
	match = true
	return
}
