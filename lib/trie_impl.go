package qse

import "fmt"

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

func (t Trie[N, L]) String() (repr string) {
	repr = fmt.Sprintf("Trie{Leaves = %v, Root = [@Root]%v}", t.leaves, t.root)
	return
}

func (node TrieValueNode[N, L]) String() (repr string) {
	repr = fmt.Sprintf(
		"V[%v]{..., Children = %v}",
		node.value, node.children,
	)
	return
}

func (node TrieLeafNode[L]) String() (repr string) {
	repr = fmt.Sprintf("L(%v)", node.value)
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

func (node *TrieValueNode[N, L]) PrepChild(seq map[N]struct{}, leaf L) (r_child TrieNode[N, L]) {
	var closest TrieNode[N, L]
	var closest_shared *map[N]struct{}
	var closest_index *int
	exact_match := false
	for index, child := range node.children {
		if child.IsTrieLeaf() {
			continue
		}
		child := child.(*TrieValueNode[N, L])
		shared := make(map[N]struct{})
		for key := range child.value {
			if _, ok := seq[key]; ok {
				shared[key] = struct{}{}
			}
		}
		length := len(shared)
		if closest == nil || length > len(*closest_shared) {
			closest = child
			closest_shared = &shared
			closest_index = &index
			exact_match = len(shared) == len(child.value)
		}
	}
	if exact_match {
		r_child = *closest.(*TrieValueNode[N, L])
		return
	}
	if closest == nil {
		if len(seq) == 0 {
			fmt.Println("new branch, leaf node")
			r_child = &TrieLeafNode[L]{
				leaf,
			}
		} else {
			fmt.Println("new branch, value node")
			seq_copy := make(map[N]struct{})
			for k := range seq {
				seq_copy[k] = struct{}{}
			}
			r_child = &TrieValueNode[N, L]{
				seq_copy,
				[]*TrieValueNode[N, L]{
					node,
				},
				make([]TrieNode[N, L], 0),
			}
		}
	} else {
		target := closest.(*TrieValueNode[N, L])
		parent_ref := target.CutPrefix(*closest_shared)
		node.children[*closest_index] = *parent_ref
		rem_seq := make(map[N]struct{})
		for item := range seq {
			if _, ok := (*closest_shared)[item]; !ok {
				rem_seq[item] = struct{}{}
			}
		}
		r_child = TrieValueNode[N, L]{
			rem_seq,
			[]*TrieValueNode[N, L]{
				parent_ref,
			},
			make([]TrieNode[N, L], 0),
		}
	}
	node.children = append(node.children, r_child)
	return
}

func (t *Trie[N, L]) Insert(seq map[N]struct{}, leaf L) {
	node := &t.root
	seq_copy := make(map[N]struct{})
	for k := range seq {
		seq_copy[k] = struct{}{}
	}
	for {
		child := node.PrepChild(seq_copy, leaf)
		if child.IsTrieLeaf() {
			break
		} else {
			switch child := child.(type) {
			case *TrieValueNode[N, L]:
				node = child
			case TrieValueNode[N, L]:
				node = &child
			}
		}
		for taken := range node.value {
			delete(seq_copy, taken)
		}
	}
}