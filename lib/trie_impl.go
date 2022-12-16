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

func (node *TrieValueNode[N, L]) PrepChild(seq map[N]struct{}, leaf L) (child TrieNode[N, L]) {
	var closest TrieNode[N, L]
	var closest_shared *map[N]struct{}
	var closest_index *int
	for index, child := range node.children {
		if child.IsTrieLeaf() {
			continue
		}
		child := child.(TrieValueNode[N, L])
		shared := make(map[N]struct{})
		for key := range child.value {
			if _, ok := seq[key]; ok {
				shared[key] = struct{}{}
			}
		}
		length := len(shared)
		if length > len(*closest_shared) {
			closest = &child
			closest_shared = &shared
			closest_index = &index
		}
	}
	if closest == nil {
		child = &TrieLeafNode[L]{
			leaf,
		}
	} else {
		target := closest.(TrieValueNode[N, L])
		node.children[*closest_index] = target.CutPrefix(*closest_shared)
	}
	node.children = append(node.children, child)
	return
}
