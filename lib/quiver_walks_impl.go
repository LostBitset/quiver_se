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
	update_walk_chunk := []E{update.edge}
	walk_prefixes := make(chan []E)
	walk_suffixes := make(chan []E)
	go func() {
		defer close(out_walks)
		prefixes := make([][]E, 0)
		for prefix_flipped := range walk_prefixes {
			prefix := make([]E, len(prefix_flipped))
			for i := range prefix {
				prefix[i] = prefix_flipped[len(prefix)-(i+1)]
			}
			prefixes = append(prefixes, prefix)
		}
		for suffix := range walk_suffixes {
			for _, prefix := range prefixes {
				l_prefix := prefix
				l_suffix := suffix
				out_walks <- QuiverWalk[N, E]{
					start,
					[]*[]E{
						&l_prefix,
						&update_walk_chunk,
						&l_suffix,
					},
				}
			}
		}
	}()
	go func() {
		q.EmitSimpleWalksFromToRev(walk_prefixes, start, update_src)
		close(walk_prefixes)
	}()
	go func() {
		q.EmitSimpleWalksFromFwd(walk_suffixes, update_dst)
		close(walk_suffixes)
	}()
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromFwd(out_simple_walks chan []E, src QuiverIndex) {
	backing_prefix := make([]E, 0)
	seen := NewPHashMap[QuiverIndex, struct{}]()
	q.EmitSimpleWalksFromFwdPrefix(out_simple_walks, src, &backing_prefix, seen)
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromFwdPrefix(
	out_simple_walks chan []E,
	src QuiverIndex,
	prefix *[]E,
	seen PHashMap[QuiverIndex, struct{}],
) {
	curr := *prefix
	out_simple_walks <- curr
	q.ForEachOutneighbor(
		src,
		func(neighbor Neighbor[E]) {
			if seen.HasKey(neighbor.dst) {
				return
			}
			curr := *prefix
			curr = append(curr, neighbor.via_edge)
			q.EmitSimpleWalksFromFwdPrefix(
				out_simple_walks,
				neighbor.dst,
				&curr,
				seen.Clone().Assoc(neighbor.dst, struct{}{}),
			)
		},
	)
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRev(
	out_simple_walks chan []E,
	src QuiverIndex,
	dst QuiverIndex,
) {
	backing_prefix := make([]E, 0)
	seen := NewPHashMap[QuiverIndex, struct{}]()
	q.EmitSimpleWalksFromToRevPrefix(out_simple_walks, src, dst, &backing_prefix, seen)
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRevPrefix(
	out_simple_walks chan []E,
	src QuiverIndex,
	true_dst QuiverIndex,
	prefix *[]E,
	seen PHashMap[QuiverIndex, struct{}],
) {
	if src == true_dst {
		curr := *prefix
		out_simple_walks <- curr
	}
	q.ForEachInneighbor(
		src,
		func(neighbor Neighbor[E]) {
			if seen.HasKey(neighbor.dst) {
				return
			}
			curr := *prefix
			curr = append(curr, neighbor.via_edge)
			q.EmitSimpleWalksFromToRevPrefix(
				out_simple_walks,
				neighbor.dst,
				true_dst,
				&curr,
				seen.Clone().Assoc(neighbor.dst, struct{}{}),
			)
		},
	)
}
