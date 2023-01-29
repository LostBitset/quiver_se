package qse

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
)

func (n uint32_H) Hash() (digest digest_t) {
	pseudo_digest := binary.LittleEndian.AppendUint32([]byte{}, n.uint32)
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{0xCE, 0xD1, 0x32})
	digest = hasher.Sum([]byte{})
	return
}

func (n uint32_H) Hash32() (fixed_digest uint32) {
	pseudo_digest := binary.LittleEndian.AppendUint32([]byte{}, n.uint32)
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{0xCE, 0xD1, 0x32})
	fixed_digest = hasher.Sum32()
	return
}

func (lit Literal[NODE]) Hash() (digest digest_t) {
	digest = lit.value.Hash()
	if !lit.eq {
		digest = WrapInvert(digest)
	}
	return
}

func (lit Literal[NODE]) Hash32() (fixed_digest uint32) {
	fixed_digest = lit.value.Hash32()
	if !lit.eq {
		fixed_digest = WrapInvert32(fixed_digest)
	}
	return
}

func WrapInvert(pseudo_digest digest_t) (digest digest_t) {
	digest = FixDigest(pseudo_digest, 0xDD)
	return
}

func WrapInvert32(pseudo_fixed_digest uint32) (fixed_digest uint32) {
	fixed_digest = FixDigest32(pseudo_fixed_digest, 0xDD)
	return
}

func ZeroDigest() (digest digest_t) {
	digest = make([]byte, 4)
	return
}

func FixDigest(pseudo_digest digest_t, fix_seed byte) (digest digest_t) {
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{fix_seed})
	digest = hasher.Sum([]byte{})
	return
}

func FixDigest32(pseudo_fixed_digest uint32, fix_seed byte) (fixed_digest uint32) {
	hasher := fnv.New32a()
	hasher.Write(
		binary.LittleEndian.AppendUint32(
			[]byte{},
			pseudo_fixed_digest,
		),
	)
	hasher.Write([]byte{fix_seed})
	fixed_digest = hasher.Sum32()
	return
}

func MixDigests(pseudo_a []byte, pseudo_b []byte) (digest digest_t) {
	hasher := fnv.New32a()
	hasher.Write([]byte{0x80})
	hasher.Write(pseudo_a)
	hasher.Write(pseudo_b)
	hasher.Write([]byte{0x80})
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
	root_value_pm := NewPHashMap[Literal[NODE], struct{}]()
	t = DMT[NODE, LEAF]{
		Trie[Literal[NODE], LEAF, digest_t]{
			TrieValueNode[Literal[NODE], LEAF, digest_t]{
				&root_value_pm,
				nil,
				[]TrieNode[Literal[NODE], LEAF]{},
				ZeroDigest(),
			},
			map[LEAF][]*TrieLeafNode[Literal[NODE], LEAF, []byte]{},
		},
	}
	return
}

func (t DMT[NODE, LEAF]) FwdLookup(a PHashMap[Literal[NODE], struct{}]) (item *LEAF) {
	item = t.trie.FwdLookup(a)
	return
}

func (t DMT[NODE, LEAF]) RevLookup(b LEAF) (items []PHashMap[Literal[NODE], struct{}]) {
	items = t.trie.RevLookup(b)
	return
}

func (t DMT[NODE, LEAF]) ForEachPair(fn func(PHashMap[Literal[NODE], struct{}], LEAF)) {
	t.trie.ForEachPair(fn)
}

func (t *DMT[NODE, LEAF]) Insert(
	seq PHashMap[Literal[NODE], struct{}], leaf LEAF,
) {
	t.InsertReturn(seq, leaf)
	return
}

func (t *DMT[NODE, LEAF]) InsertReturn(
	seq PHashMap[Literal[NODE], struct{}], leaf LEAF,
) (
	leaf_ptr *TrieLeafNode[Literal[NODE], LEAF, digest_t],
) {
	leaf_ptr = t.trie.InsertReturn(seq, leaf)
	t.UpdateHashes(leaf_ptr)
	final_leaves := t.trie.leaves[leaf]
	t.MergeIdenticalLeaves(final_leaves)
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
				value_hash := ZeroDigest()
				for itr := c.value.inner.Iterator(); itr.HasElem(); itr.Next() {
					key_any, _ := itr.Elem()
					key := key_any.(Literal[NODE])
					for i, byte_value := range key.Hash() {
						value_hash[i] ^= byte_value
					}
				}
				child_meta = MixDigests(child_meta, value_hash)
			case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
				child_meta = c.meta
				value_hash := ZeroDigest()
				for itr := c.value.inner.Iterator(); itr.HasElem(); itr.Next() {
					key_any, _ := itr.Elem()
					key := key_any.(Literal[NODE])
					for i, byte_value := range key.Hash() {
						value_hash[i] ^= byte_value
					}
				}
				child_meta = MixDigests(child_meta, value_hash)
			}
			for i, byte_value := range child_meta {
				sum[i] ^= byte_value
			}
		}
		node.meta = FixDigest(sum, 0x3E) // Update the hash
		t.SimplifyNode(node)             // Try and locally simplify
		node = node.parent
		if node == nil {
			break correctLoop
		}
	}
}

func (t *DMT[NODE, LEAF]) SimplifyNode(node *TrieValueNode[Literal[NODE], LEAF, digest_t]) {
	edge_indices := make([]int, 0)
	edge_values := make([]PHashMap[Literal[NODE], struct{}], 0)
getChildEdgesLoop:
	for i, child := range node.children {
		switch c := child.(type) {
		case TrieLeafNode[Literal[NODE], LEAF, digest_t]:
			continue getChildEdgesLoop
		case *TrieLeafNode[Literal[NODE], LEAF, digest_t]:
			continue getChildEdgesLoop
		case TrieValueNode[Literal[NODE], LEAF, digest_t]:
			edge_indices = append(edge_indices, i)
			edge_values = append(edge_values, *c.value)
		case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
			edge_indices = append(edge_indices, i)
			edge_values = append(edge_values, *c.value)
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

func (t *DMT[NODE, LEAF]) MergeIdenticalLeaves(leaves []*TrieLeafNode[Literal[NODE], LEAF, digest_t]) {
	for i := 1; i < len(leaves); i++ {
		unwanted_leaf := leaves[i]
		parent := unwanted_leaf.parent
		var index int
		for i, option := range parent.children {
		findLeafTypeSwitch:
			switch c := option.(type) {
			case TrieLeafNode[Literal[NODE], LEAF, digest_t]:
				if unwanted_leaf == &c {
					index = i
				}
			case *TrieLeafNode[Literal[NODE], LEAF, digest_t]:
				if unwanted_leaf == c {
					index = i
				}
			default:
				break findLeafTypeSwitch
			}
		}
		SpliceOutReclaim(&parent.children, index)
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
	copy(bypass_children, child.children)
	node.children = append(node.children, bypass_children...)
	for _, bypass_child := range bypass_children {
		switch c := bypass_child.(type) {
		case TrieLeafNode[Literal[NODE], LEAF, digest_t]:
			*c.parent = *node
		case *TrieLeafNode[Literal[NODE], LEAF, digest_t]:
			*c.parent = *node
		case TrieValueNode[Literal[NODE], LEAF, digest_t]:
			*c.parent = *node
		case *TrieValueNode[Literal[NODE], LEAF, digest_t]:
			*c.parent = *node
		}
	}
}

func IsInvertedEdge[NODE hashable](maybe_buf PHashMap[Literal[NODE], struct{}], maybe_inv PHashMap[Literal[NODE], struct{}]) (match bool) {
	maybe_inv_copy := NewPHashMap[Literal[NODE], struct{}]()
	for itr := maybe_inv.inner.Iterator(); itr.HasElem(); itr.Next() {
		k_any, _ := itr.Elem()
		k := k_any.(Literal[NODE])
		maybe_inv_copy = maybe_inv_copy.Assoc(k, struct{}{})
	}
	for itr := maybe_buf.inner.Iterator(); itr.HasElem(); itr.Next() {
		key_any, _ := itr.Elem()
		key := key_any.(Literal[NODE])
		inverted := key.Invert()
		if maybe_inv_copy.HasKey(inverted) {
			maybe_inv_copy = maybe_inv_copy.Dissoc(inverted)
		} else {
			match = false
			return
		}
	}
	match = (maybe_inv_copy.length == 0)
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
