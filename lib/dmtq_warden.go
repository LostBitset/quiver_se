package qse

type DMTQWardenConfig[N any, ATOM hashable, AUG any] struct {
	in_updates chan Augmented[
		QuiverUpdate[N, PHashMap[Literal[ATOM], struct{}], *DMT[ATOM, QuiverIndex]],
		AUG,
	]
	out_walks chan Augmented[
		QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]],
		AUG,
	]
	walk_src QuiverIndex
	walk_dst QuiverIndex
	dmtq     Quiver[N, PHashMap[Literal[ATOM], struct{}], *DMT[ATOM, QuiverIndex]]
}

type Augmented[A any, B any] struct {
	Value   A
	Augment B
}
