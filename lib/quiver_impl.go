package qse

func (obj SimpleReversibleAssoc[A, B]) Insert(a A, b B) {
	obj.backing_map[a] = b
}

func (obj SimpleReversibleAssoc[A, B]) FwdLookup(a A) (item B) {
	item = obj.backing_map[a]
	return
}

func (obj SimpleReversibleAssoc[A, B]) RevLookup(b B) (items []A) {
	for k, v := range obj.backing_map {
		if v == b {
			items = append(items, k)
		}
	}
	return
}
