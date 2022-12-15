package qse

func NewSimpleRA[A comparable, B comparable]() (obj SimpleReversibleAssoc[A, B]) {
	obj = SimpleReversibleAssoc[A, B]{
		make(map[A]B),
	}
	return
}

func (obj *SimpleReversibleAssoc[A, B]) Insert(a A, b B) {
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

func (obj SimpleReversibleAssoc[A, B]) ForEachKey(fn func(A)) {
	for k := range obj.backing_map {
		fn(k)
	}
}

func (q *Quiver[N, E, C]) insert_node(node_value N, container C) (idx QuiverIndex) {
	node := QuiverNode[N, E, C]{
		node_value,
		make([]QuiverIndex, 0),
		container,
	}
	q.arena = append(q.arena, node)
	idx = QuiverIndex(len(q.arena) - 1)
	return
}

func (q *Quiver[N, E, C]) insert_edge(src, dst QuiverIndex, edge_value E) {
	src_node, dst_node := q.arena[src], q.arena[dst]
	src_node.edges.Insert(edge_value, dst)
	dst_node.parents = append(dst_node.parents, src)
}

type Neighbor[E any] struct {
	via_edge E
	dst      QuiverIndex
}

func (q *Quiver[N, E, C]) all_outneighbors(src QuiverIndex) (outneighbors []Neighbor[E]) {
	src_node := q.arena[src]
	src_node.edges.ForEachKey(func(edge E) {
		neighbor := Neighbor[E]{
			edge,
			src_node.edges.FwdLookup(edge),
		}
		outneighbors = append(outneighbors, neighbor)
	})
	return
}

/*func (q *Quiver[N, E, C]) all_inneighbors(src QuiverIndex) (inneighbors []Neighbor[E]) {
	src_node := q.arena[src]
	for parent := range src_node.parents {}
}*/

func (q *SimpleQuiver[N, E]) insert_node_simple(node_value N) (idx QuiverIndex) {
	container := NewSimpleRA[E, QuiverIndex]()
	idx = q.insert_node(node_value, &container)
	return
}
