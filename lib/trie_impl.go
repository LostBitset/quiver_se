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

func (t *Trie[N, L]) Insert(seq []N, leaf L) {
	node := &t.root
	cursor := 0
	for {
		var closest_node *TrieValueNode[N, L]
		var closest_by *int
		var closest_exact_match *bool
		var closest_index *int
		for i, child := range node.children {
			if child.IsTrieLeaf() {
				continue
			}
			value_node := child.(*TrieValueNode[N, L])
			speculative_cursor := cursor
			exact_match := true
			for _, expected := range value_node.value {
				if speculative_cursor >= len(seq) || expected != seq[speculative_cursor] {
					exact_match = false
					break
				}
				speculative_cursor++
			}
			if closest_by == nil || (speculative_cursor-cursor) > *closest_by {
				closest_by_value := speculative_cursor - cursor
				closest_node = value_node
				closest_by = &closest_by_value
				closest_exact_match = &exact_match
				closest_index = &i
			}
		}
		if closest_node == nil {
			var leaf_part TrieNode[N, L]
			leaf_part = TrieLeafNode[L]{
				leaf,
			}
			if cursor != len(seq)-1 {
				leaf_part = TrieValueNode[N, L]{
					seq[cursor:],
					make([]TrieValueNode[N, L], 0), // TODO parent pointers
					[]TrieNode[N, L]{
						leaf_part,
					},
				}
			}
			node.children = append(node.children, leaf_part)
			return
		}
		if *closest_exact_match {
			cursor += *closest_by
			node = closest_node
		} else {
			new_node := TrieValueNode[N, L]{
				seq[cursor : cursor+*closest_by],
				make([]TrieValueNode[N, L], 0), // TODO parent pointers
				[]TrieNode[N, L]{
					closest_node,
				},
			}
			node.children[*closest_index] = new_node
			cursor += *closest_by
			node = &new_node
		}
	}
}

func (t Trie[N, L]) Lookup(seq []N) (ptr *L) {
	node := t.root
	cursor := 0
	for {
		moved_cursor := false
		for _, child := range node.children {
			if child.IsTrieLeaf() {
				leaf := child.(TrieLeafNode[L])
				if cursor == len(seq)-1 {
					ptr = &leaf.value
					return
				} else {
					break
				}
			} else {
				value_node := child.(TrieValueNode[N, L])
				for _, expected := range value_node.value {
					if cursor >= len(seq) || seq[cursor] != expected {
						break
					}
					moved_cursor = true
					cursor++
				}
			}
		}
		if !moved_cursor {
			return
		}
	}
}
