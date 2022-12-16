package qse

func (TrieValueNode[N, L]) IsTrieLeaf() (is bool) {
	is = false
	return
}

func (TrieLeafNode[L]) IsTrieLeaf() (is bool) {
	is = true
	return
}

func NewTrie[N comparable, L comparable]() (t Trie[N, L]) {
	t = Trie[N, L]{
		NewTrieValueNode[N, L](),
		make(map[L]struct{}),
	}
	return
}

func NewTrieValueNode[N comparable, L comparable]() (node TrieValueNode[N, L]) {
	node = TrieValueNode[N, L]{
		make(map[N]struct{}),
		make([]*TrieValueNode[N, L], 0),
		make([]TrieNode[N, L], 0),
	}
	return
}

func (node *TrieValueNode[N, L]) CutPrefix(shared map[N]struct{}) (parent *TrieValueNode[N, L]) {
	parent = &TrieValueNode[N, L]{
		shared,
		node.parents,
		[]TrieNode[N, L]{
			node,
		},
	}
	for shared_key := range shared {
		delete(node.value, shared_key)
	}
	node.parents = append(node.parents, parent)
	return
}

func (node *TrieValueNode[N, L]) PrepChild(seq map[N]struct{}) (child *TrieNode[N, L]) {
	var closest *TrieNode[N, L]
	var closest_by *int
	for _, child := range node.children {
		if child.IsTrieLeaf() {
			continue
		}
		child := child.(TrieValueNode[N, L])
		matches := 0
		for key := range child.value {
			if _, ok := seq[key]; ok {
				matches++
			}
		}
		if matches > *closest_by {
			// TODO
		}
		// TODO
	}
	// TODO
}
