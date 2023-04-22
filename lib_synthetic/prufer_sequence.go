package libsynthetic

type PruferSequence struct {
	degrees  []int
	sequence []int
}

type SimpleTree struct {
	id       int
	children []*SimpleTree
}
