package qse

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
)

func (n uint32_H) Hash() (digest []byte) {
	pseudo_digest := binary.LittleEndian.AppendUint32([]byte{}, n.uint32)
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{0xCE, 0xD1, 0x32})
	digest = hasher.Sum([]byte{})
	return
}

func (n uint64_H) Hash() (digest []byte) {
	pseudo_digest := binary.LittleEndian.AppendUint64([]byte{}, n.uint64)
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{0xCE, 0xD1, 0x64})
	digest = hasher.Sum([]byte{})
	return
}

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

func NewDMT[NODE hashable, LEAF hashable]() (t DMT[NODE, LEAF]) {
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

func (t DMT[NODE, LEAF]) FwdLookup(a map[Literal[NODE]]struct{}) (item *LEAF) {
	item = t.trie.FwdLookup(a)
	return
}

func (t DMT[NODE, LEAF]) RevLookup(b LEAF) (items []map[Literal[NODE]]struct{}) {
	items = t.trie.RevLookup(b)
	return
}

func (t DMT[NODE, LEAF]) ForEachPair(fn func(map[Literal[NODE]]struct{}, LEAF)) {
	t.trie.ForEachPair(fn)
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
		for _, raw_child := range node.children {
			var child_meta []byte
			switch c := raw_child.(type) {
			case TrieLeafNode[Literal[NODE], LEAF, digest_t]:
				child_meta = FixDigest(c.value.Hash(), 0xA4)
			case *TrieLeafNode[Literal[NODE], LEAF, digest_t]:
				child_meta = FixDigest(c.value.Hash(), 0xA4)
			case TrieValueNode[Literal[NODE], LEAF, digest_t]:
				child_meta = c.meta
			case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
				child_meta = c.meta
			}
			for i, byte_value := range child_meta {
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
				buf_i, inv_i := edge_indices[buf_ii], edge_indices[inv_ii]
				var buf_subtrie_hash digest_t
				var inv_subtrie_hash digest_t
				switch c := node.children[buf_i].(type) {
				case TrieValueNode[Literal[NODE], LEAF, digest_t]:
					buf_subtrie_hash = c.meta
				case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
					buf_subtrie_hash = c.meta
				}
				switch c := node.children[inv_i].(type) {
				case TrieValueNode[Literal[NODE], LEAF, digest_t]:
					inv_subtrie_hash = c.meta
				case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
					inv_subtrie_hash = c.meta
				}
				if DigestEqual(buf_subtrie_hash, inv_subtrie_hash) {
					unwanted_children = append(unwanted_children, buf_i)
					unwanted_children = append(unwanted_children, inv_i)
					t.ShiftChildren(node, buf_i)
				}
			}
		}
	}
	InsertionSortInPlace(unwanted_children)
	DedupSortedInPlace(&unwanted_children)
	offset := 0
	for _, child_i := range unwanted_children {
		// Drop the child
		child_i := child_i - offset
		SpliceOutReclaim(&node.children, child_i)
		offset++
	}
}

func (t *DMT[NODE, LEAF]) ShiftChildren(node *TrieValueNode[Literal[NODE], LEAF, digest_t], child_i int) {
	var child *TrieValueNode[Literal[NODE], LEAF, digest_t]
	switch c := node.children[child_i].(type) {
	case TrieValueNode[Literal[NODE], LEAF, digest_t]:
		child = &c
	case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
		child = c
	}
	bypass_children := make([]TrieNode[Literal[NODE], LEAF], len(child.children))
	for i, bypass_child := range child.children {
		bypass_children[i] = bypass_child
	}
	for _, bypass_child := range bypass_children {
		node.children = append(node.children, bypass_child)
	}
}

func IsInvertedEdge[NODE hashable](maybe_buf map[Literal[NODE]]struct{}, maybe_inv map[Literal[NODE]]struct{}) (match bool) {
	maybe_inv_copy := make(map[Literal[NODE]]struct{})
	for k := range maybe_inv {
		maybe_inv_copy[k] = struct{}{}
	}
	for key := range maybe_buf {
		inverted := key.Invert()
		if _, ok := maybe_inv_copy[inverted]; ok {
			delete(maybe_inv_copy, inverted)
		} else {
			match = false
			return
		}
	}
	match = (len(maybe_inv_copy) == 0)
	return
}

func InsertionSortInPlace(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
	findTargetPosition:
		for {
			if j == 0 || arr[j-1] <= arr[j] {
				break findTargetPosition
			}
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
}

func DedupSortedInPlace(arr *[]int) {
	for i := 0; i < (len(*arr) - 1); i++ {
		if (*arr)[i] == (*arr)[i+1] {
			SpliceOutReclaim(arr, i+1)
			i--
		}
	}
}

func SpliceOutReclaim[T any](arr *[]T, index int) {
	var zero T
	copy((*arr)[index:], (*arr)[(index+1):])
	(*arr)[len(*arr)-1] = zero
	*arr = (*arr)[:(len(*arr) - 1)]
}

func (dmt DMT[NODE, LEAF]) EntryList() (out []TrieEntry[Literal[NODE], LEAF]) {
	out = dmt.trie.EntryList()
	return
}
