package qse

type ReversibleAssoc[A any, B any] interface {
}

type QuiverIndex uint

type QuiverNode[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	value   N
	parents []QuiverIndex
	edges   C
}
