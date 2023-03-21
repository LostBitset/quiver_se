package main

func (tree SimpleTree) AsMicroprogramTransitions(
	dst_states []MicroprogramState,
	constraintgen ConstraintGenerator,
) (
	transitions []MicroprogramTransition,
) {
	transitions = tree.AsMicroprogramTransitionsWithPrefix(dst_states, constraintgen, []string{})
	return
}

func (tree SimpleTree) AsMicroprogramTransitionsWithPrefix(
	dst_states []MicroprogramState,
	constraintgen ConstraintGenerator,
	constraint_prefix []string,
) (
	transitions []MicroprogramTransition,
) {
	transitions = make([]MicroprogramTransition, len(tree.children))
	for i, child := range tree.children {

		transitions[i] = new_transition
	}
}
