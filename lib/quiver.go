package qse

type ReversibleAssoc[A any, B any] interface {
	Insert(a A, b B)
	FwdLookup(a A) (item B)
	RevLookup(b B) (items []A)
}

type SimpleReversibleAssoc[A comparable, B comparable] struct {
	backing_map map[A]B
}

type QuiverIndex uint

type SimpleQuiverNode[N any, E comparable] struct {
	QuiverNode[N, E, SimpleReversibleAssoc[E, QuiverIndex]]
}

type QuiverNode[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	value   N
	parents []QuiverIndex
	edges   C
}
