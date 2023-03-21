package main

import qse "LostBitset/quiver_se/lib"

func BuildTestingMicroprogramGenerator() (uprgm_gen MicroprogramGenerator) {
	ops, vals := GetStandardItems()
	constraint_gen := ConstraintGenerator{
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
	constraint_gen.AddVariables(4, var_sorts_distr, 0.75)
	uprgm_gen = MicroprogramGenerator{
		n_states:          7,
		p_transition:      0.75,
		avg_n_transitions: 2.0,
		p_fallible:        0.4,
		n_entry_samples:   3,
		n_tree_nonleaf:    4,
		constraintgen:     constraint_gen,
		smt_free_funs:     []qse.SMTFreeFun[string, string]{},
	}
	return
}
