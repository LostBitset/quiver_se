package libsynthetic

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConstraintGeneration(t *testing.T) {
	ops, vals := GetStandardItems()
	gtor := ConstraintGenerator{
		P_n_depth_mean:   2.0,
		P_n_depth_stddev: 1.5,
		P_ops:            ops,
		P_vals:           vals,
		NextVarId:        Pto(0),
	}
	var_sorts := SimpleDDistr[Sort]{
		map[Sort]float64{
			RealSort: 0.7,
			BoolSort: 0.3,
		},
	}
	var_sorts_distr := BakeDDistr[Sort](var_sorts)
	gtor.AddVariables(4, var_sorts_distr, 0.75)
	fmt.Printf("Generating %d random SMTLib-v2 constraints.\n", TEST_ITERATIONS)
	for i := 0; i < TEST_ITERATIONS; i++ {
		test := gtor.Generate(BoolSort)
		fmt.Printf(strconv.Itoa(i)+": %#+v\n", test)
	}

}
