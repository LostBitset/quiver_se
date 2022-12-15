package qse

func (TrieValueNode[N, L]) IsTrieLeaf() (is bool) {
	is = false
	return
}

func (TrieLeafNode[L]) IsTrieLeaf() (is bool) {
	is = true
	return
}

func NewTrie[N any, L comparable]() (t Trie[N, L]) {
	t = Trie[N, L]{
		NewTrieValueNode[N, L](),
		make(map[L]struct{}),
	}
	return
}

func NewTrieValueNode[N any, L comparable]() (node TrieValueNode[N, L]) {
	node = TrieValueNode[N, L]{
		make([]N, 0),
		make([]TrieValueNode[N, L], 0),
		make([]TrieNode[N, L], 0),
	}
	return
}
