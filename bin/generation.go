package main

type Generator struct {
	n_callbacks    int
	n_depth_mean   float64
	n_depth_stddev float64
	n_operators    []Op
	n_values       []Val
}

type Op struct {
	name string
	args []Sort
	ret  Sort
}

type Val struct {
	name string
	sort Sort
}

type Sort uint8

const (
	RealSort Sort = iota
	BoolSort
)
