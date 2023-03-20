package main

type PruferSequence struct {
	occurences []int
	sequence   []int
}

type SimpleTree struct {
	id    int
	left  *SimpleTree
	right *SimpleTree
}
