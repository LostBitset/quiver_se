package qse

func (t *DMT[NODE, LEAF]) Insert(seq map[Literal[NODE]]struct{}, leaf LEAF) {
	t.trie.Insert(seq, leaf)
	t.UpdateHashes(seq)
}

func (t *DMT[NODE, LEAF]) UpdateHashes(seq map[Literal[NODE]]struct{}) {
	query_copy := make(map[Literal[NODE]]struct{})
	for k := range seq {
		query_copy[k] = struct{}{}
	}
	node := &t.trie.root
}
