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
		make([]TrieValueNode[N, L], 0),
		make([]TrieNode[N, L], 0),
	}
	return
}

func (t *Trie[N, L]) Insert(seq map[N]struct{}, leaf L) {
	node := &t.root
	already_found := make(map[N]struct{})
	for elem := range seq {
		for _, child := range node.children {
			if child.IsTrieLeaf() {
				// TODO CASE is a leaf
			} else {
				child := child.(TrieValueNode[N, L])
				if _, ok := child.value[elem]; ok {
					if 
					// TODO CASE found match
				} else {
					// TODO CASE no match found
				}
			}
		}
	}
}
