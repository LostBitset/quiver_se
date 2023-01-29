package qse

func (TrieValueNode[NODE, LEAF, META]) IsTrieLeaf() (is bool) {
	is = false
	return
}

func (TrieLeafNode[NODE, LEAF, META]) IsTrieLeaf() (is bool) {
	is = true
	return
}

func NewTrie[NODE hashable, LEAF comparable]() (t Trie[NODE, LEAF, struct{}]) {
	t = Trie[NODE, LEAF, struct{}]{
		NewTrieRootNode[NODE, LEAF](),
		make(map[LEAF][]*TrieLeafNode[NODE, LEAF, struct{}]),
	}
	return
}

func NewTrieRootNode[NODE hashable, LEAF comparable]() (node TrieValueNode[NODE, LEAF, struct{}]) {
	pm := NewPHashMap[NODE, struct{}]()
	node = TrieValueNode[NODE, LEAF, struct{}]{
		&pm,
		nil,
		make([]TrieNode[NODE, LEAF], 0),
		struct{}{},
	}
	return
}

func (node *TrieValueNode[NODE, LEAF, META]) CutPrefix(shared PHashMap[NODE, struct{}]) (parent *TrieValueNode[NODE, LEAF, META]) {
	parent = &TrieValueNode[NODE, LEAF, META]{
		&shared,
		node.parent,
		[]TrieNode[NODE, LEAF]{
			node,
		},
		node.meta,
	}
	for itr := shared.inner.Iterator(); itr.HasElem(); itr.Next() {
		shared_key_any, _ := itr.Elem()
		shared_key := shared_key_any.(NODE)
		*node.value = node.value.Dissoc(shared_key)
	}
	node.parent = parent
	return
}

func (node *TrieValueNode[NODE, LEAF, META]) PrepChild(seq *PHashMap[NODE, struct{}], leaf LEAF) (r_child TrieNode[NODE, LEAF]) {
	var closest TrieNode[NODE, LEAF]
	var closest_shared *PHashMap[NODE, struct{}]
	var closest_index *int
	exact_match := false
	leaf_count := 0
	for index, child := range node.children {
		if child.IsTrieLeaf() {
			leaf_count++
			continue
		}
		child_nc := child
		child := child_nc.(*TrieValueNode[NODE, LEAF, META])
		shared := NewPHashMap[NODE, struct{}]()
		for itr := child.value.inner.Iterator(); itr.HasElem(); itr.Next() {
			key_any, _ := itr.Elem()
			key := key_any.(NODE)
			if (*seq).HasKey(key) {
				shared = shared.Assoc(key, struct{}{})
			}
		}
		length := shared.length
		if length == 0 {
			continue
		}
		if closest == nil || length > (*closest_shared).length {
			closest = child
			closest_shared = new(PHashMap[NODE, struct{}])
			*closest_shared = shared
			closest_index = new(int)
			*closest_index = index
			exact_match = shared.length == child.value.length
		}
	}
	if leaf_count != 0 && leaf_count == len(node.children) {
		r_child = &TrieLeafNode[NODE, LEAF, META]{
			leaf,
			new(TrieValueNode[NODE, LEAF, META]),
		}
		seq_copy := NewPHashMap[NODE, struct{}]()
		for itr := (*seq).inner.Iterator(); itr.HasElem(); itr.Next() {
			k_any, _ := itr.Elem()
			k := k_any.(NODE)
			seq_copy = seq_copy.Assoc(k, struct{}{})

		}
		extension_node := &TrieValueNode[NODE, LEAF, META]{
			&seq_copy,
			new(TrieValueNode[NODE, LEAF, META]),
			[]TrieNode[NODE, LEAF]{
				r_child,
			},
			node.meta,
		}
		*seq = NewPHashMap[NODE, struct{}]()
		node.children = append(node.children, extension_node)
		return
	}
	if exact_match {
		skip_ref := closest.(*TrieValueNode[NODE, LEAF, META])
		r_child = skip_ref
		for itr := skip_ref.value.inner.Iterator(); itr.HasElem(); itr.Next() {
			k_any, _ := itr.Elem()
			k := k_any.(NODE)
			*seq = (*seq).Dissoc(k)
		}
		return
	}
	if closest == nil {
		r_child = &TrieLeafNode[NODE, LEAF, META]{
			leaf,
			new(TrieValueNode[NODE, LEAF, META]),
		}
		seq_copy := NewPHashMap[NODE, struct{}]()
		for itr := (*seq).inner.Iterator(); itr.HasElem(); itr.Next() {
			k_any, _ := itr.Elem()
			k := k_any.(NODE)
			seq_copy = seq_copy.Assoc(k, struct{}{})
		}
		r_child_inner := &TrieValueNode[NODE, LEAF, META]{
			&seq_copy,
			new(TrieValueNode[NODE, LEAF, META]),
			[]TrieNode[NODE, LEAF]{
				r_child,
			},
			node.meta,
		}
		node.children = append(node.children, r_child_inner)
		*seq = NewPHashMap[NODE, struct{}]()
		return
	} else {
		r_child = &TrieLeafNode[NODE, LEAF, META]{
			leaf,
			new(TrieValueNode[NODE, LEAF, META]),
		}
		target := closest.(*TrieValueNode[NODE, LEAF, META])
		parent_ref := target.CutPrefix(*closest_shared)
		node.children[*closest_index] = parent_ref
		rem_seq := NewPHashMap[NODE, struct{}]()
		for itr := (*seq).inner.Iterator(); itr.HasElem(); itr.Next() {
			item_any, _ := itr.Elem()
			item := item_any.(NODE)
			if !closest_shared.HasKey(item) {
				rem_seq = rem_seq.Assoc(item, struct{}{})
			}
		}
		r_child_inner := &TrieValueNode[NODE, LEAF, META]{
			&rem_seq,
			new(TrieValueNode[NODE, LEAF, META]),
			[]TrieNode[NODE, LEAF]{
				r_child,
			},
			node.meta,
		}
		*seq = NewPHashMap[NODE, struct{}]()
		parent_ref.children = append(parent_ref.children, r_child_inner)
		return
	}
}

func (t *Trie[NODE, LEAF, META]) Insert(seq PHashMap[NODE, struct{}], leaf LEAF) {
	t.InsertReturn(seq, leaf)
	return
}

func (t *Trie[NODE, LEAF, META]) InsertReturn(seq PHashMap[NODE, struct{}], leaf LEAF) (leaf_ptr *TrieLeafNode[NODE, LEAF, META]) {
	node := &t.root
	seq_copy := NewPHashMap[NODE, struct{}]()
	for itr := seq.inner.Iterator(); itr.HasElem(); itr.Next() {
		k_any, _ := itr.Elem()
		k := k_any.(NODE)
		seq_copy = seq_copy.Assoc(k, struct{}{})
	}
	for {
		child := node.PrepChild(&seq_copy, leaf)
		if child.IsTrieLeaf() {
			child := child.(*TrieLeafNode[NODE, LEAF, META])
			if _, ok := t.leaves[leaf]; !ok {
				t.leaves[leaf] = make([]*TrieLeafNode[NODE, LEAF, META], 0)
			}
			t.leaves[leaf] = append(t.leaves[leaf], child)
			leaf_ptr = child
			break
		} else {
			child := child.(*TrieValueNode[NODE, LEAF, META])
			node = child
		}
	}
	t.LookupRepair(seq)
	return
}

func (t Trie[NODE, LEAF, META]) LookupRepair(query PHashMap[NODE, struct{}]) {
	query_copy := NewPHashMap[NODE, struct{}]()
	for itr := query.inner.Iterator(); itr.HasElem(); itr.Next() {
		k_any, _ := itr.Elem()
		k := k_any.(NODE)
		query_copy = query_copy.Assoc(k, struct{}{})
	}
	node := &t.root
searchLoop:
	for {
		for itr := node.value.inner.Iterator(); itr.HasElem(); itr.Next() {
			key_any, _ := itr.Elem()
			key := key_any.(NODE)
			query_copy = query_copy.Dissoc(key)
		}
		if query_copy.length == 0 {
			for _, child := range node.children {
				if child.IsTrieLeaf() {
					child := child.(*TrieLeafNode[NODE, LEAF, META])
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
			child := child.(*TrieValueNode[NODE, LEAF, META])
			has_match := false
			for itr := child.value.inner.Iterator(); itr.HasElem(); itr.Next() {
				key_any, _ := itr.Elem()
				key := key_any.(NODE)
				if query_copy.HasKey(key) {
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

func (e TrieEntry[NODE, LEAF]) PrefixWith(prefix PHashMap[NODE, struct{}]) (mod TrieEntry[NODE, LEAF]) {
	key := map[NODE]struct{}{}
	for k, v := range e.key {
		key[k] = v
	}
	for itr := prefix.inner.Iterator(); itr.HasElem(); itr.Next() {
		k_any, v_any := itr.Elem()
		k := k_any.(NODE)
		v := v_any.(struct{})
		key[k] = v
	}
	mod = TrieEntry[NODE, LEAF]{
		key,
		e.value,
	}
	return
}

// Recursively call a method on all key-value pairs defined in a trie
func (t Trie[NODE, LEAF, META]) ForEachEntry(fn func(TrieEntry[NODE, LEAF])) {
	t.root.ForEachNodeEntry(fn)
}

func (node TrieValueNode[NODE, LEAF, META]) ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF])) {
	prefix := node.value
	for _, child := range node.children {
		child.ForEachNodeEntry(func(entry TrieEntry[NODE, LEAF]) {
			fn(entry.PrefixWith(*prefix))
		})
	}
}

func (node TrieLeafNode[NODE, LEAF, META]) ForEachNodeEntry(fn func(TrieEntry[NODE, LEAF])) {
	fn(TrieEntry[NODE, LEAF]{
		map[NODE]struct{}{},
		node.value,
	})
}

// Enumerate all mappings contained within a trie
func (t Trie[NODE, LEAF, META]) EntryList() (out []TrieEntry[NODE, LEAF]) {
	t.ForEachEntry(func(entry TrieEntry[NODE, LEAF]) {
		out = append(out, entry)
	})
	return
}

func (t Trie[NODE, LEAF, META]) Lookup(query PHashMap[NODE, struct{}]) (leaf *LEAF) {
	query_copy := NewPHashMap[NODE, struct{}]()
	for itr := query.inner.Iterator(); itr.HasElem(); itr.Next() {
		k_any, _ := itr.Elem()
		k := k_any.(NODE)
		query_copy = query_copy.Assoc(k, struct{}{})
	}
	node := &t.root
searchLoop:
	for {
		for itr := node.value.inner.Iterator(); itr.HasElem(); itr.Next() {
			key_any, _ := itr.Elem()
			key := key_any.(NODE)
			query_copy = query_copy.Dissoc(key)
		}
		if query_copy.length == 0 {
			for _, child := range node.children {
				if child.IsTrieLeaf() {
					leaf = &child.(*TrieLeafNode[NODE, LEAF, META]).value
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
			child := child.(*TrieValueNode[NODE, LEAF, META])
			has_match := false
			for itr := child.value.inner.Iterator(); itr.HasElem(); itr.Next() {
				key_any, _ := itr.Elem()
				key := key_any.(NODE)
				if query_copy.HasKey(key) {
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

func (t Trie[NODE, LEAF, META]) LookupLeaf(leaf LEAF) (seqs []PHashMap[NODE, struct{}]) {
	leaf_nodes, ok := t.leaves[leaf]
	if !ok {
		return
	}
	for _, leaf_node := range leaf_nodes {
		seqs = append(seqs, t.LookupLeafByNode(leaf_node))
	}
	return
}

func (t Trie[NODE, LEAF, META]) LookupLeafByNode(leaf_node *TrieLeafNode[NODE, LEAF, META]) (seq PHashMap[NODE, struct{}]) {
	seq = NewPHashMap[NODE, struct{}]()
	node := leaf_node.parent
	for {
		if node == nil {
			break
		}
		for itr := node.value.inner.Iterator(); itr.HasElem(); itr.Next() {
			k_any, _ := itr.Elem()
			k := k_any.(NODE)
			seq = seq.Assoc(k, struct{}{})
		}
		node = node.parent
	}
	return
}

func (t Trie[NODE, LEAF, META]) FwdLookup(a PHashMap[NODE, struct{}]) (item *LEAF) {
	item = t.Lookup(a)
	return
}

func (t Trie[NODE, LEAF, META]) RevLookup(b LEAF) (items []PHashMap[NODE, struct{}]) {
	items = t.LookupLeaf(b)
	return
}

func (t Trie[NODE, LEAF, META]) ForEachPair(fn func(map[NODE]struct{}, LEAF)) {
	t.ForEachEntry(func(entry TrieEntry[NODE, LEAF]) {
		fn(entry.key, entry.value)
	})
}
