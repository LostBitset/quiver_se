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

// This invalidates the reciever (node)
func (node *TrieValueNode[N, L]) CutPrefix(shared map[N]struct{}) (
	parent *TrieValueNode[N, L], child *TrieValueNode[N, L],
) {
	parent = &TrieValueNode[N, L]{
		shared,
		node.parents,
		append(node.children, child),
	}
}
