package qse

import "fmt"

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
	dst_final QuiverIndex,
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
			fmt.Println("got prefix")
			prefixes = append(prefixes, prefix)
		}
		for suffix_flipped := range walk_suffixes {
			suffix := make([]E, len(suffix_flipped))
			for i := range suffix {
				suffix[i] = suffix_flipped[len(suffix)-(i+1)]
			}
			fmt.Println("got suffix")
			for _, prefix := range prefixes {
				l_prefix := prefix
				l_suffix := suffix
				fmt.Println("SEND")
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
		q.EmitSimpleWalksFromToRev(walk_prefixes, update_src, start)
		close(walk_prefixes)
	}()
	go func() {
		q.EmitSimpleWalksFromToRev(walk_suffixes, dst_final, update_dst)
		close(walk_suffixes)
	}()
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRev(
	out_simple_walks chan []E,
	src QuiverIndex,
	dst QuiverIndex,
) {
	backing_prefix := make([]E, 0)
	seen := NewPHashMap[QuiverIndex, uint8]()
	q.EmitSimpleWalksFromToRevPrefix(out_simple_walks, src, dst, &backing_prefix, seen)
}

const QUIVER_SIMPLE_WALKS_MAX_TRAVERSAL_CYCLE_COUNT = 1

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRevPrefix(
	out_simple_walks chan []E,
	src QuiverIndex,
	true_dst QuiverIndex,
	prefix *[]E,
	seen PHashMap[QuiverIndex, uint8],
) {
	if src == true_dst {
		curr := *prefix
		out_simple_walks <- curr
	}
	q.ForEachInneighbor(
		src,
		func(neighbor Neighbor[E]) {
			prev_count := uint8(0)
			if count, ok := seen.Index(neighbor.dst); ok {
				if count >= QUIVER_SIMPLE_WALKS_MAX_TRAVERSAL_CYCLE_COUNT {
					return
				}
				prev_count = count
			}
			curr := *prefix
			curr = append(curr, neighbor.via_edge)
			q.EmitSimpleWalksFromToRevPrefix(
				out_simple_walks,
				neighbor.dst,
				true_dst,
				&curr,
				seen.Clone().Assoc(neighbor.dst, prev_count+1),
			)
		},
	)
}
