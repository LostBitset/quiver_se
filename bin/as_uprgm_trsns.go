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
		if len(tree.children) == 2 {
			left_child, right_child := tree.children[0], tree.children[1]
			left_constraint := constraintgen.Generate(BoolSort)
			left_constraint_prefix := make([]string, len(constraint_prefix)+1)
			copy(left_constraint_prefix, constraint_prefix)
			left_constraint_prefix[len(constraint_prefix)] = left_constraint
			transitions = append(transitions, left_child.AsMicroprogramTransitionsWithPrefix(
				dst_states,
				constraintgen,
				left_constraint_prefix,
			)...)
			right_constraint := "(not " + left_constraint + ")"
			right_constraint_prefix := make([]string, len(constraint_prefix)+1)
			copy(right_constraint_prefix, constraint_prefix)
			right_constraint_prefix[len(constraint_prefix)] = right_constraint
			transitions = append(transitions, right_child.AsMicroprogramTransitionsWithPrefix(
				dst_states,
				constraintgen,
				right_constraint_prefix,
			)...)
		} else {
			panic("Invalid. Must be a binary tree. ")
		}
	}
	return
}
