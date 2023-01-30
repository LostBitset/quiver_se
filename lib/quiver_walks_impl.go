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

func MarkIntentToBeQuiverNode[N any, E any, C ReversibleAssoc[E, QuiverIndex]](node N) (
	intended_node QuiverIntendedNode[N, E, C],
) {
	intended_node = QuiverIntendedNode[N, E, C]{
		node,
		NewPhantomQuiverAssociation[N, E, C](),
	}
	return
}

func (Quiver[N, E, C]) MarkNodeIntent(node N) (intended_node QuiverIntendedNode[N, E, C]) {
	intended_node = MarkIntentToBeQuiverNode[N, E, C](node)
	return
}

func (indexp QuiverIndexParameterized[N, E, C]) ResolveAsQuiverUpdateDst(q_ptr *Quiver[N, E, C]) (
	index QuiverIndex,
) {
	// TODO
}

func (intended_node QuiverIntendedNode[N, E, C]) ResolveAsQuiverUpdateDst(q_ptr *Quiver[N, E, C]) (
	index QuiverIndex,
) {
	// TODO
}
