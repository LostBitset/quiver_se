package qse

type DMTQWardenConfig[N any, ATOM hashable] struct {
	in_updates chan QuiverUpdate[N, PHashMap[Literal[ATOM], struct{}], *DMT[ATOM, QuiverIndex]]
	out_walks  chan QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]]
	walk_src   QuiverIndex
	dmtq       Quiver[N, PHashMap[Literal[ATOM], struct{}], *DMT[ATOM, QuiverIndex]]
}
