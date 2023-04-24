package libsynthetic

type SimpleQuiverAdjList struct {
	adj_list []SimpleEdgeDesc
}

type SimpleEdgeDesc struct {
	src int
	dst int
}

type ConnectedOrNot int

const (
	Connected ConnectedOrNot = iota
	NotConnected
)
