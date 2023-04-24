package libsynthetic

type ConstraintGenerator struct {
	P_n_depth_mean   float64
	P_n_depth_stddev float64
	P_ops            map[Sort][]Op
	P_vals           map[Sort][]Val
	NextVarId        *int
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
