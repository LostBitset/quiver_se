package libsynthetic

func BuildTestingMicroprogramGenerator() (uprgm_gen MicroprogramGenerator) {
	ops, vals := GetStandardItems()
	constraint_gen := ConstraintGenerator{
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
	constraint_gen.AddVariables(4, var_sorts_distr, 0.75)
	uprgm_gen = MicroprogramGenerator{
		P_n_states:        30,
		P_p_transition:    0.3,
		P_n_merged_graphs: 2,
		P_p_fallible:      0.5,
		P_n_entry_samples: 7,
		P_n_tree_nonleaf:  4,
		P_constraintgen:   constraint_gen,
	}
	return
}
