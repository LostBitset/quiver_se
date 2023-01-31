package qse

// A nicely packaged way to represent association with a particular quiver
type PhantomQuiverAssociation[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	phantom_node PhantomData[N]
	phantom_edge PhantomData[E]
	phantom_cont PhantomData[C]
}

// An alternative to the QuiverIndex type that carries type parameters
type QuiverIndexParameterized[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	QuiverIndex
	phantom_association PhantomQuiverAssociation[N, E, C]
}

// Allow edges to be stored in pointers to chunks to limit (to some extent) duplication
type QuiverWalk[N any, E any] struct {
	start         QuiverIndex
	edges_chunked []*[]*E
}

// A new edge (and possibly a new node) being added to a quiver
// The new_dst key indicates whether dst is already a part of the quiver or not
type QuiverUpdate[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	src  QuiverIndex
	dst  QuiverUpdateDst[N, E, C]
	edge E
}

// A non-trivial QuiverUpdateDst - a node itself that has to be added to the quiver
// But this must be wrapped with it's intent (mostly for type parameterization)
type QuiverIntendedNode[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	node                N
	container           C
	phantom_association PhantomQuiverAssociation[N, E, C]
}

type QuiverUpdateDst[N any, E any, C ReversibleAssoc[E, QuiverIndex]] interface {
	ResolveAsQuiverUpdateDst(q_ptr *Quiver[N, E, C]) (index QuiverIndex)
}
