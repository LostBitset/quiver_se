package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops, vals := GetStandardItems()
	gtor := ConstraintGenerator{
		n_depth_mean:   2.0,
		n_depth_stddev: 1.5,
		ops:            ops,
		vals:           vals,
		next_var_id:    pto(0),
	}
	var_sorts := SimpleDDistr[Sort]{
		map[Sort]float64{
			RealSort: 0.7,
			BoolSort: 0.3,
		},
	}
	var_sorts_distr := BakeDDistr[Sort](var_sorts)
	gtor.AddVariables(4, var_sorts_distr, 0.75)
	for i := 0; i < 5; i++ {
		test := gtor.Generate(BoolSort)
		fmt.Printf(strconv.Itoa(i)+": %#+v\n", test)
	}
	seq := RandomPruferSequence(5)
	fmt.Printf("seq: %#+v\n", seq.sequence)
	tree := seq.ToTree()
	fmt.Printf("size: %#+v\n", tree.ComputeSize())
}
