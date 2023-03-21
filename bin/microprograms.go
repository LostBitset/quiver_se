package main

import (
	qse "LostBitset/quiver_se/lib"
)

type MicroprogramState int

type Microprogram struct {
	top_state     MicroprogramState
	fail_state    MicroprogramState
	transitions   map[MicroprogramState][]MicroprogramTransition
	smt_free_funs []qse.SMTFreeFun[string, string]
}

type MicroprogramTransition struct {
	dst_state  MicroprogramState
	constraint string
}

type MicroprogramGenerator struct {
	n_states          int
	p_transition      float64
	avg_n_transitions int
	p_fallible        float64
	n_entry_samples   int
	n_tree_nonleaf    int
	constraintgen     ConstraintGenerator
	smt_free_funs     []qse.SMTFreeFun[string, string]
	next_state_id     MicroprogramState
}
