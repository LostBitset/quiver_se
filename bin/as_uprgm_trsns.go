package main

func (tree SimpleTree) AsMicroprogramTransitions(
	leaf_allocation MicroprogramState,
	constraintgen ConstraintGenerator,
) (
	transitions []MicroprogramTransition,
) {
	transitions = tree.AsMicroprogramTransitionsWithPrefix(leaf_allocation, constraintgen, []string{})
	return
}

func (tree SimpleTree) AsMicroprogramTransitionsWithPrefix(
	leaf_allocation MicroprogramState,
	constraintgen ConstraintGenerator,
	constraint_prefix []string,
) (
	transitions []MicroprogramTransition,
) {
	// TODO
}
