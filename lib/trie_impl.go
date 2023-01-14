package qse

func (TrieValueNode[NODE, LEAF]) IsTrieLeaf() (is bool) {
	is = false
	return
}

func (TrieLeafNode[NODE, LEAF]) IsTrieLeaf() (is bool) {
	is = true
	return
}

func NewTrie[NODE comparable, LEAF comparable]() (t Trie[NODE, LEAF]) {
	t = Trie[NODE, LEAF]{
		NewTrieRootNode[NODE, LEAF](),
		make(map[LEAF][]*TrieLeafNode[NODE, LEAF]),
	}
	return
}

func NewTrieRootNode[NODE comparable, LEAF comparable]() (node TrieValueNode[NODE, LEAF]) {
	node = TrieValueNode[NODE, LEAF]{
		make(map[NODE]struct{}),
		nil,
		make([]TrieNode[NODE, LEAF], 0),
	}
	return
}

func (node *TrieValueNode[NODE, LEAF]) CutPrefix(shared map[NODE]struct{}) (parent *TrieValueNode[NODE, LEAF]) {
	parent = &TrieValueNode[NODE, LEAF]{
		shared,
		node.parent,
		[]TrieNode[NODE, LEAF]{
			node,
		},
	}
	for shared_key := range shared {
		delete(node.value, shared_key)
	}
	node.parent = parent
	return
}

func (node *TrieValueNode[NODE, LEAF]) PrepChild(seq *map[NODE]struct{}, leaf LEAF) (r_child TrieNode[NODE, LEAF]) {
	var closest TrieNode[NODE, LEAF]
	var closest_shared *map[NODE]struct{}
	var closest_index *int
	exact_match := false
	leaf_count := 0
	for index, child := range node.children {
		if child.IsTrieLeaf() {
			leaf_count++
			continue
		}
		child_nc := child
		child := child_nc.(*TrieValueNode[NODE, LEAF])
		shared := make(map[NODE]struct{})
		for key := range child.value {
			if _, ok := (*seq)[key]; ok {
				shared[key] = struct{}{}
			}
		}
		length := len(shared)
		if length == 0 {
			continue
		}
		if closest == nil || length > len(*closest_shared) {
			closest = child
			closest_shared = new(map[NODE]struct{})
			*closest_shared = shared
			closest_index = new(int)
			*closest_index = index
			exact_match = len(shared) == len(child.value)
		}
	}
	if leaf_count != 0 && leaf_count == len(node.children) {
		r_child = &TrieLeafNode[NODE, LEAF]{
			leaf,
			new(TrieValueNode[NODE, LEAF]),
		}
		seq_copy := make(map[NODE]struct{})
		for k := range *seq {
			seq_copy[k] = struct{}{}
		}
		extension_node := &TrieValueNode[NODE, LEAF]{
			seq_copy,
			new(TrieValueNode[NODE, LEAF]),
			[]TrieNode[NODE, LEAF]{
				r_child,
			},
		}
		*seq = make(map[NODE]struct{})
		node.children = append(node.children, extension_node)
		return
	}
	if exact_match {
		skip_ref := closest.(*TrieValueNode[NODE, LEAF])
		r_child = skip_ref
		for k := range skip_ref.value {
			delete(*seq, k)
		}
		return
	}
	if closest == nil {
		r_child = &TrieLeafNode[NODE, LEAF]{
			leaf,
			new(TrieValueNode[NODE, LEAF]),
		}
		seq_copy := make(map[NODE]struct{})
		for k := range *seq {
			seq_copy[k] = struct{}{}
		}
		r_child_inner := &TrieValueNode[NODE, LEAF]{
			seq_copy,
			new(TrieValueNode[NODE, LEAF]),
			[]TrieNode[NODE, LEAF]{
				r_child,
			},
		}
		node.children = append(node.children, r_child_inner)
		*seq = make(map[NODE]struct{})
		return
	} else {
		r_child = &TrieLeafNode[NODE, LEAF]{
			leaf,
			new(TrieValueNode[NODE, LEAF]),
		}
		target := closest.(*TrieValueNode[NODE, LEAF])
		parent_ref := target.CutPrefix(*closest_shared)
		node.children[*closest_index] = parent_ref
		rem_seq := make(map[NODE]struct{})
		for item := range *seq {
			if _, ok := (*closest_shared)[item]; !ok {
				rem_seq[item] = struct{}{}
			}
		}
		r_child_inner := &TrieValueNode[NODE, LEAF]{
			rem_seq,
			new(TrieValueNode[NODE, LEAF]),
			[]TrieNode[NODE, LEAF]{
				r_child,
			},
		}
		*seq = make(map[NODE]struct{})
		parent_ref.children = append(parent_ref.children, r_child_inner)
		return
	}
}

func (t *Trie[NODE, LEAF]) Insert(seq map[NODE]struct{}, leaf LEAF) (leaf_ptr *TrieLeafNode[NODE, LEAF]) {
	node := &t.root
	seq_copy := make(map[NODE]struct{})
	for k := range seq {
		seq_copy[k] = struct{}{}
	}
	for {
		child := node.PrepChild(&seq_copy, leaf)
		if child.IsTrieLeaf() {
			child := child.(*TrieLeafNode[NODE, LEAF])
			if _, ok := t.leaves[leaf]; !ok {
				t.leaves[leaf] = make([]*TrieLeafNode[NODE, LEAF], 0)
			}
			t.leaves[leaf] = append(t.leaves[leaf], child)
			leaf_ptr = child
			break
		} else {
			child := child.(*TrieValueNode[NODE, LEAF])
			node = child
		}
	}
	t.LookupRepair(seq)
	return
}

func (t Trie[NODE, LEAF]) LookupRepair(query map[NODE]struct{}) {
	query_copy := make(map[NODE]struct{})
	for k := range query {
		query_copy[k] = struct{}{}
	}
	node := &t.root
searchLoop:
	for {
		for key := range node.value {
			delete(query_copy, key)
		}
		if len(query_copy) == 0 {
			for _, child := range node.children {
				if child.IsTrieLeaf() {
					child := child.(*TrieLeafNode[NODE, LEAF])
					child.parent = node
					return
				}
			}
		}
	checkChildrenLoop:
		for _, child := range node.children {
			if child.IsTrieLeaf() {
				continue checkChildrenLoop
			}
			child := child.(*TrieValueNode[NODE, LEAF])
			has_match := false
			for key := range child.value {
				if _, ok := query_copy[key]; ok {
					has_match = true
				} else {
					continue checkChildrenLoop
				}
			}
			if has_match {
				child.parent = node
				node = child
				continue searchLoop
			}
		}
	}
}

func (e TrieEntry[NODE, LEAF]) PrefixWith(prefix map[NODE]struct{}) (mod TrieEntry[NODE, LEAF]) {
	key := make(map[NODE]struct{})
	for k, v := range e.key {
		key[k] = v
	}
	for k, v := range prefix {
		key[k] = v
	}
	mod = TrieEntry[NODE, LEAF]{
		key,
		e.value,
	}
	return
}

// Recursively call a method on all key-value pairs defined in a trie
func (t Trie[NODE, LEAF]) ForEachEntry(fn func(TrieEntry[NODE, LEAF])) {
	t.root.ForEachNodeEntry(fn)
}

func (node TrieValueNode[NODE, LEAF]) ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF])) {
	prefix := node.value
	for _, child := range node.children {
		child.ForEachNodeEntry(func(entry TrieEntry[NODE, LEAF]) {
			fn(entry.PrefixWith(prefix))
		})
	}
}

func (node TrieLeafNode[NODE, LEAF]) ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF])) {
	fn(TrieEntry[NODE, LEAF]{
		make(map[NODE]struct{}),
		node.value,
	})
}

// Enumerate all mappings contained within a trie
func (t Trie[NODE, LEAF]) EntryList() (out []TrieEntry[NODE, LEAF]) {
	t.ForEachEntry(func(entry TrieEntry[NODE, LEAF]) {
		out = append(out, entry)
	})
	return
}

func (t Trie[NODE, LEAF]) Lookup(query map[NODE]struct{}) (leaf *LEAF) {
	query_copy := make(map[NODE]struct{})
	for k := range query {
		query_copy[k] = struct{}{}
	}
	node := &t.root
searchLoop:
	for {
		for key := range node.value {
			delete(query_copy, key)
		}
		if len(query_copy) == 0 {
			for _, child := range node.children {
				if child.IsTrieLeaf() {
					leaf = &child.(*TrieLeafNode[NODE, LEAF]).value
					return
				}
			}
			leaf = nil
			return
		}
	checkChildrenLoop:
		for _, child := range node.children {
			if child.IsTrieLeaf() {
				continue checkChildrenLoop
			}
			child := child.(*TrieValueNode[NODE, LEAF])
			has_match := false
			for key := range child.value {
				if _, ok := query_copy[key]; ok {
					has_match = true
				} else {
					continue checkChildrenLoop
				}
			}
			if has_match {
				node = child
				continue searchLoop
			}
		}
		leaf = nil
		return
	}
}

func (t Trie[NODE, LEAF]) LookupLeaf(leaf LEAF) (seqs []map[NODE]struct{}) {
	leaf_nodes, ok := t.leaves[leaf]
	if !ok {
		return
	}
	for _, leaf_node := range leaf_nodes {
		seqs = append(seqs, t.LookupLeafByNode(leaf_node))
	}
	return
}

func (t Trie[NODE, LEAF]) LookupLeafByNode(leaf_node *TrieLeafNode[NODE, LEAF]) (seq map[NODE]struct{}) {
	seq = make(map[NODE]struct{})
	node := leaf_node.parent
	for {
		if node == nil {
			break
		}
		for k := range node.value {
			seq[k] = struct{}{}
		}
		node = node.parent
	}
	return
}

func (t Trie[NODE, LEAF]) FwdLookup(a map[NODE]struct{}) (item *LEAF) {
	item = t.Lookup(a)
	return
}

func (t Trie[NODE, LEAF]) RevLookup(b LEAF) (items []map[NODE]struct{}) {
	items = t.LookupLeaf(b)
	return
}

func (t Trie[NODE, LEAF]) ForEachPair(fn func(map[NODE]struct{}, LEAF)) {
	t.ForEachEntry(func(entry TrieEntry[NODE, LEAF]) {
		fn(entry.key, entry.value)
	})
}
