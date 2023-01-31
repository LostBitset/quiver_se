package qse

import (
	"fmt"
	"sync"
)

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
	var wg sync.WaitGroup
	wg.Add(2)
	walk_prefixes := make(chan []*E)
	walk_suffixes := make(chan []*E)
	go func() {
		defer wg.Done()
		q.EmitSimpleWalksFromToRev(walk_prefixes, start, update_src)
	}()
	go func() {
		defer wg.Done()
		q.EmitSimpleWalksFromFwd(walk_suffixes, update_dst)
	}()
	wgsig := StartWaitGroupSignal(&wg)
	go func() {
		known_suffixes := make([][]*E, 0)
		known_prefixes := make([][]*E, 0)
	sendNewWalksLoop:
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
			case <-wgsig:
				break sendNewWalksLoop
			}
		}
		close(out_walks)
	}()
}

func StartWaitGroupSignal(wg *sync.WaitGroup) (wgsig chan struct{}) {
	wgsig = make(chan struct{})
	go func() {
		wg.Wait()
		fmt.Println("WAIT GROUP ENDED")
		wgsig <- struct{}{}
	}()
	return
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromFwd(out_simple_walks chan []*E, src QuiverIndex) {
	backing_prefix := make([]*E, 0)
	q.EmitSimpleWalksFromFwdMutPrefix(out_simple_walks, src, &backing_prefix)
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromFwdMutPrefix(
	out_simple_walks chan []*E,
	src QuiverIndex,
	prefix *[]*E,
) {
	curr := *prefix
	out_simple_walks <- curr
	fmt.Printf("FromFwd ! %v\n", curr)
	q.ForEachOutneighbor(
		src,
		func(neighbor Neighbor[E]) {
			fmt.Printf("somehow found neighbor %v\n", neighbor)
			*prefix = append(*prefix, &neighbor.via_edge)
			curr_prime := *prefix
			q.EmitSimpleWalksFromFwdMutPrefix(
				out_simple_walks,
				neighbor.dst,
				&curr_prime,
			)
			(*prefix)[len(*prefix)-1] = nil
			*prefix = (*prefix)[:len(*prefix)-1]
		},
	)
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRev(
	out_simple_walks chan []*E,
	src QuiverIndex,
	dst QuiverIndex,
) {
	backing_prefix := make([]*E, 0)
	q.EmitSimpleWalksFromToRevMutPrefix(out_simple_walks, src, dst, &backing_prefix)
}

func (q Quiver[N, E, C]) EmitSimpleWalksFromToRevMutPrefix(
	out_simple_walks chan []*E,
	src QuiverIndex,
	dst QuiverIndex,
	prefix *[]*E,
) {
	curr := *prefix
	out_simple_walks <- curr
	fmt.Printf("FromToRev ! %v\n", curr)
	q.ForEachInneighbor(
		src,
		func(neighbor Neighbor[E]) {
			*prefix = append(*prefix, &neighbor.via_edge)
			curr_prime := *prefix
			if neighbor.dst == dst {
				fmt.Printf("FromToRev (terminal) ! %v\n", curr)
				out_simple_walks <- curr_prime
				return
			}
			q.EmitSimpleWalksFromFwdMutPrefix(
				out_simple_walks,
				neighbor.dst,
				&curr_prime,
			)
			(*prefix)[len(*prefix)-1] = nil
			*prefix = (*prefix)[:len(*prefix)-1]
		},
	)
}
