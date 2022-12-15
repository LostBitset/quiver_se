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
		make([]N, 0),
		make([]TrieValueNode[N, L], 0),
		make([]TrieNode[N, L], 0),
	}
	return
}

func (t Trie[N, L]) Lookup(seq []N) (ptr *L) {
	node := t.root
	cursor := 0
	for {
		for _, child := range node.children {
			if child.IsTrieLeaf() {
				leaf := child.(TrieLeafNode[L])
				if cursor == len(seq)-1 {
					ptr = &leaf.value
					return
				}
			} else {
				value_node := child.(TrieValueNode[N, L])
				for _, expected := range value_node.value {
					if cursor >= len(seq) || seq[cursor] != expected {
						break
					}
					cursor++
				}
			}
		}
	}
}
