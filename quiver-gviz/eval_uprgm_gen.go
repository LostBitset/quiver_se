package main

import (
	s "github.com/LostBitset/quiver_se/lib_synthetic"
)

func GenerateEvaluationMicroprogram() (uprgm s.Microprogram) {
	uprgm_gen := BuildEvaluationMicroprogramGenerator()
	uprgm = uprgm_gen.RandomMicroprogram()
	return
}

func BuildEvaluationMicroprogramGenerator() (uprgm_gen s.MicroprogramGenerator) {
	ops, vals := s.GetStandardItems()
	constraint_gen := s.ConstraintGenerator{
		P_n_depth_mean:   2.0,
		P_n_depth_stddev: 1.5,
		P_ops:            ops,
		P_vals:           vals,
		NextVarId:        s.Pto(0),
	}
	var_sorts := s.SimpleDDistr[s.Sort]{
		Outcomes: map[s.Sort]float64{
			s.RealSort: 0.7,
			s.BoolSort: 0.3,
		},
	}
	var_sorts_distr := s.BakeDDistr[s.Sort](var_sorts)
	constraint_gen.AddVariables(4, var_sorts_distr, 0.75)
	uprgm_gen = s.MicroprogramGenerator{
		P_n_states:          30,
		P_p_transition:      0.8,
		P_avg_n_transitions: 40.0,
		P_p_fallible:        0.9,
		P_n_entry_samples:   7,
		P_n_tree_nonleaf:    3,
		P_constraintgen:     constraint_gen,
	}
	return
}
