package libsynthetic

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

type MicroprogramState int

type Microprogram struct {
	StateTop      MicroprogramState
	StateFail     MicroprogramState
	Transitions   map[MicroprogramState][]MicroprogramTransition
	smt_free_funs []qse.SMTFreeFun[string, string]
}

type MicroprogramTransition struct {
	StateDst    MicroprogramState
	Constraints []string
}

type MicroprogramGenerator struct {
	P_n_states        int
	P_p_transition    float64
	P_n_merged_graphs int
	P_p_fallible      float64
	P_n_entry_samples int
	P_n_tree_nonleaf  int
	P_constraintgen   ConstraintGenerator
	next_state_id     MicroprogramState
}
