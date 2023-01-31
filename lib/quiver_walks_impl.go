package qse

func NewPhantomQuiverAssociation[N any, E any, C ReversibleAssoc[E, QuiverIndex]]() (
	phantom_association PhantomQuiverAssociation[N, E, C],
) {
	phantom_association = PhantomQuiverAssociation[N, E, C]{
		PhantomData[N]{},
		PhantomData[E]{},
		PhantomData[C]{},
	}
	return
}

func TrustingParameterizeQuiverIndex[N any, E any, C ReversibleAssoc[E, QuiverIndex]](
	index QuiverIndex,
) (
	indexp QuiverIndexParameterized[N, E, C],
) {
	indexp = QuiverIndexParameterized[N, E, C]{
		index,
		NewPhantomQuiverAssociation[N, E, C](),
	}
	return
}

func (indexp QuiverIndexParameterized[N, E, C]) Unparameterize() (index QuiverIndex) {
	index = indexp.QuiverIndex
	return
}

func (Quiver[N, E, C]) ParameterizeIndex(index QuiverIndex) (indexp QuiverIndexParameterized[N, E, C]) {
	indexp = TrustingParameterizeQuiverIndex[N, E, C](index)
	return
}

func NewQuiverIntendedNode[N any, E any, C ReversibleAssoc[E, QuiverIndex], UNUSED_BOUNDARY any](
	node N, container C,
) (
	intended_node QuiverIntendedNode[N, E, C],
) {
	intended_node = QuiverIntendedNode[N, E, C]{
		node,
		container,
		NewPhantomQuiverAssociation[N, E, C](),
	}
	return
}

func (Quiver[N, E, C]) NewIntendedNode(node N, container C) (intended_node QuiverIntendedNode[N, E, C]) {
	intended_node = NewQuiverIntendedNode[N, E, C, any](node, container)
	return
}

func (indexp QuiverIndexParameterized[N, E, C]) ResolveAsQuiverUpdateDst(q_ptr *Quiver[N, E, C]) (
	index QuiverIndex,
) {
	index = indexp.Unparameterize()
	return
}

func (intended_node QuiverIntendedNode[N, E, C]) ResolveAsQuiverUpdateDst(q_ptr *Quiver[N, E, C]) (
	index QuiverIndex,
) {
	index = q_ptr.InsertNode(
		intended_node.node,
		intended_node.container,
	)
	return
}

func (q *Quiver[N, E, C]) ApplyUpdate(update QuiverUpdate[N, E, C]) (src, dst QuiverIndex) {
	src = update.src
	dst = update.dst.ResolveAsQuiverUpdateDst(q)
	q.InsertEdge(src, dst, update.edge)
	return
}

func (q *Quiver[N, E, C]) ApplyUpdateAndEmitWalks(
	out_walks chan QuiverWalk[N, E],
	update QuiverUpdate[N, E, C],
	start_unresolved QuiverIndexParameterized[N, E, C],
) {
	update_src, update_dst := q.ApplyUpdate(update)
	start := start_unresolved.ResolveAsQuiverUpdateDst(q)
	update_walk_chunk := []*E{&update.edge}
	walk_prefixes := make(chan []*E)
	walk_suffixes := make(chan []*E)
	go q.EmitSimpleWalksFromToRev(walk_prefixes, start, update_src)
	go q.EmitSimpleWalksFromFwd(walk_suffixes, update_dst)
	go func() {
		known_suffixes := make([][]*E, 0)
		known_prefixes := make([][]*E, 0)
		for {
			select {
			case prefix := <-walk_prefixes:
				for _, known_suffix := range known_suffixes {
					out_walks <- QuiverWalk[N, E]{
						start,
						[]*[]*E{&prefix, &update_walk_chunk, &known_suffix},
					}
				}
				known_prefixes = append(known_prefixes, prefix)
			case suffix := <-walk_suffixes:
				for _, known_prefix := range known_prefixes {
					out_walks <- QuiverWalk[N, E]{
						start,
						[]*[]*E{&known_prefix, &update_walk_chunk, &suffix},
					}
				}
			}
		}
	}()
}

const QUIVER_EMIT_SIMPLE_WALKS_DFS_MEMOIZATION_MAX_SIZE = 10

func (q Quiver[N, E, C]) EmitSimpleWalksFromFwd(Aout_simple_walks chan []*E, src QuiverIndex) {
	prefix := make([]*E, 0)
	index_stack := make([]uint, 0)
	cursor := src
	outneighbors_memo := make(map[QuiverIndex]*[]Neighbor[E])
dfsLoop:
	for {
		var outneighbors []Neighbor[E]
		if memoized, ok := outneighbors_memo[cursor]; ok {
			outneighbors = *memoized
		} else {
			outneighbors = q.AllOutneighbors(cursor)
			for len(outneighbors_memo) >= (QUIVER_EMIT_SIMPLE_WALKS_DFS_MEMOIZATION_MAX_SIZE - 1) {
			destroyOneKeyLoop:
				for key := range outneighbors_memo {
					delete(outneighbors_memo, key)
					break destroyOneKeyLoop
				}
			}
			outneighbors_memo[cursor] = &outneighbors
		}
		index_stack[len(index_stack)-1]++
		top_index := index_stack[len(index_stack)-1]
		if top_index >= uint(len(outneighbors)) {
			prefix[len(prefix)-1] = nil
			prefix = prefix[:(len(prefix) - 1)]
			var zero_uint uint
			index_stack[len(index_stack)-1] = zero_uint
			index_stack = index_stack[:(len(index_stack) - 1)]
			if len(index_stack) == 0 {
				break dfsLoop
			} else {
				continue dfsLoop
			}
		}
		new_neighbor := outneighbors[top_index]
		prefix[len(prefix)-1] = &new_neighbor.via_edge
		cursor = new_neighbor.dst
	}
	// TODO
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRev(
	out_simple_walks chan []*E,
	src QuiverIndex,
	dst QuiverIndex,
) {
	// TODO
}
