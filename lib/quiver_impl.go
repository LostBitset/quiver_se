package qse

func (obj SimpleReversibleAssoc[A, B]) EmptyOut() {}

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

func (q Quiver[N, E, C]) insert_node(node_value N) (idx QuiverIndex) {
	container := new(C)
	(*container).EmptyOut()
	node := QuiverNode[N, E, C]{
		node_value,
		make([]QuiverIndex, 0),
		*container,
	}
	q.arena = append(q.arena, node)
	idx = QuiverIndex(len(q.arena) - 1)
	return
}

func (q Quiver[N, E, C]) insert_edge(src, dst QuiverIndex, edge_value E) {
	q.arena[src].edges.Insert(edge_value, dst)
	q.arena[dst].parents = append(q.arena[dst].parents, src)
}
