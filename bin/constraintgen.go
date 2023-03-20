package main

type ConstraintGenerator struct {
	n_depth_mean   float64
	n_depth_stddev float64
	ops            map[Sort][]Op
	vals           map[Sort][]Val
	next_var_id    *int
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
