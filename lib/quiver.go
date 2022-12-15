package qse

// An interface that represents a data structure that maps values in one direction, but allows you
// to query in both. Each A is mapped to one B, but you're free to look up all A values associated
// with a specific B.
type ReversibleAssoc[A any, B any] interface {
	Insert(a A, b B)
	FwdLookup(a A) (item B)
	RevLookup(b B) (items []A)
	ForEachKey(fn func(A))
}

// A simple implementation of a ReversibleAssoc data structure backed by a map. The RevLookup
// method runs in O(n).
type SimpleReversibleAssoc[A comparable, B comparable] struct {
	backing_map map[A]B
}

// A named type representing an index into the internal data structure of a quiver. It refers to a
// specific node, and is only usable with a specific Quiver[N, E, C] object.
type QuiverIndex uint

// A node of a quiver. Parent references are done with QuiverIndex objects and not references.
type QuiverNode[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	value   N
	parents []QuiverIndex
	edges   C
}

// A simple quiver using SimpleReversibleAssoc as the edge container.
type SimpleQuiver[N any, E comparable] struct {
	Quiver[N, E, *SimpleReversibleAssoc[E, QuiverIndex]]
}

// A doubly-linked arena-based quiver. Abstracted over arbitrary edge container types, which store
// all of the edges for a particular node.
type Quiver[N any, E any, C ReversibleAssoc[E, QuiverIndex]] struct {
	arena []QuiverNode[N, E, C]
}
