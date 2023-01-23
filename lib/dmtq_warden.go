package qse

type InnerDMTQ[NODE hashable, ATOM hashable] *Quiver[NODE, map[Literal[ATOM]]struct{}, *DMT[ATOM, QuiverIndex]]

type DMTQNodeMapping[NODE hashable] struct {
	node  NODE
	index QuiverIndex
}

type DMTQNodeUpdate[NODE hashable] struct {
	node            NODE
	mappping_return chan DMTQNodeMapping[NODE]
}

type DMTQEdgeUpdate[NODE hashable, ATOM hashable] struct {
	src     QuiverIndex
	dst     QuiverIndex
	formula map[Literal[ATOM]]struct{}
}

type DMTQTransaction[NODE hashable, ATOM hashable] interface {
	Invoke(dmtq InnerDMTQ[NODE, ATOM])
}

type DMTQWalk[NODE hashable, ATOM hashable] struct {
	edges  []map[Literal[ATOM]]struct{}
	target NODE
}

type DMTQWardenConfig[NODE hashable, ATOM hashable] struct {
	walk_chan chan DMTQWalk[NODE, ATOM]
}

