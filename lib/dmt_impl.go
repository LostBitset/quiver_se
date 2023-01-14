package qse

func (t *DMT[NODE, LEAF]) Insert(seq map[Literal[NODE]]struct{}, leaf LEAF) (leaf_ptr *TrieLeafNode[Literal[NODE], LEAF]) {
	leaf_ptr = t.trie.Insert(seq, leaf)
	t.UpdateHashes(leaf_ptr)
	return
}

func (t *DMT[NODE, LEAF]) UpdateHashes(leaf_ptr *TrieLeafNode[Literal[NODE], LEAF]) {
}
