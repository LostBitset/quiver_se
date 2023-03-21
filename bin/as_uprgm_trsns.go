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
	if len(tree.children) == 0 {
		// Leaf case
		if tree.id >= 0 {
			panic("Invalid tree. All leafs must be non-negative. This was not the case. ")
		}
		dst_state := dst_states[(-tree.id)-1]
		transitions = []MicroprogramTransition{
			{
				dst_state:   dst_state,
				constraints: constraint_prefix,
			},
		}
	} else {
		// Non-leaf case
		transitions = make([]MicroprogramTransition, 0)
		for _, child := range tree.children {
			new_constraint := constraintgen.Generate(BoolSort)
			new_constraint_prefix := make([]string, len(constraint_prefix)+1)
			copy(new_constraint_prefix, constraint_prefix)
			new_constraint_prefix[len(constraint_prefix)] = new_constraint
			transitions = append(transitions, child.AsMicroprogramTransitionsWithPrefix(
				dst_states,
				constraintgen,
				new_constraint_prefix,
			)...)
		}
	}
	return
}
