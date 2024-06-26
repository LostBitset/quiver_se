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

func (obj SimpleReversibleAssoc[A, B]) FwdLookup(a A) (item *B) {
	item_value, ok := obj.backing_map[a]
	if ok {
		item = &item_value
	}
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

func (obj SimpleReversibleAssoc[A, B]) ForEachPair(fn func(A, B)) {
	for k, v := range obj.backing_map {
		fn(k, v)
	}
}

func (q *Quiver[N, E, C]) InsertNode(node_value N, container C) (idx QuiverIndex) {
	node := QuiverNode[N, E, C]{
		node_value,
		make(map[QuiverIndex]struct{}),
		container,
	}
	q.arena = append(q.arena, node)
	idx = QuiverIndex(len(q.arena) - 1)
	return
}

func (q *Quiver[N, E, C]) InsertEdge(src, dst QuiverIndex, edge_value E) {
	src_node, dst_node := q.arena[src], q.arena[dst]
	src_node.edges.Insert(edge_value, dst)
	dst_node.parents[src] = struct{}{}
}

type Neighbor[E any] struct {
	via_edge E
	dst      QuiverIndex
}

func (q *Quiver[N, E, C]) ForEachOutneighbor(src QuiverIndex, fn func(Neighbor[E])) {
	src_node := q.arena[src]
	src_node.edges.ForEachPair(func(edge E, dst QuiverIndex) {
		neighbor := Neighbor[E]{
			edge,
			dst,
		}
		fn(neighbor)
	})
}

func (q *Quiver[N, E, C]) ForEachInneighbor(src QuiverIndex, fn func(Neighbor[E])) {
	src_node := q.arena[src]
	for parent := range src_node.parents {
		parent_node := q.arena[parent]
		for _, edge := range parent_node.edges.RevLookup(src) {
			neighbor := Neighbor[E]{
				edge,
				parent,
			}
			fn(neighbor)
		}
	}
}

func (q *Quiver[N, E, C]) AllOutneighbors(src QuiverIndex) (outneighbors []Neighbor[E]) {
	q.ForEachOutneighbor(
		src,
		func(neighbor Neighbor[E]) {
			outneighbors = append(outneighbors, neighbor)
		},
	)
	return
}

func (q *Quiver[N, E, C]) AllInneighbors(src QuiverIndex) (inneighbors []Neighbor[E]) {
	q.ForEachInneighbor(
		src,
		func(neighbor Neighbor[E]) {
			inneighbors = append(inneighbors, neighbor)
		},
	)
	return
}

func (q *SimpleQuiver[N, E]) InsertNodeSimple(node_value N) (idx QuiverIndex) {
	container := NewSimpleRA[E, QuiverIndex]()
	idx = q.InsertNode(node_value, &container)
	return
}
