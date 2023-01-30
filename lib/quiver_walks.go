package qse

type QuiverWalk[N any, E any] struct {
	start N
	edges []E
}

// A new edge (and possibly a new node) being added to a quiver
// The new_dst key indicates whether dst is already a part of the quiver or not
type QuiverUpdate[N any, E any] struct {
	src     N
	dst     N
	edge    E
	new_dst bool
}

