package libsynthetic

import "fmt"

func (tree SimpleTree) AsMicroprogramTransitions(
	dst_states []MicroprogramState,
	constraintgen ConstraintGenerator,
) (
	transitions []MicroprogramTransition,
) {
	transitions = tree.AsMicroprogramTransitionsWithPrefix(
		dst_states,
		constraintgen,
		[]string{},
		[]string{},
	)
	return
}

func (tree SimpleTree) AsMicroprogramTransitionsWithPrefix(
	dst_states []MicroprogramState,
	constraintgen ConstraintGenerator,
	constraint_prefix []string,
	invertible_constraint_prefix []string,
) (
	transitions []MicroprogramTransition,
) {
	if len(tree.children) == 0 {
		// Leaf case
		if tree.id >= 0 {
			tree.id = -1
		}
		/*fmt.Printf("tree.id = %#+v\n", tree.id)
		fmt.Printf("dst_states index = %#+v\n", (-tree.id)-1)
		fmt.Printf("dst_states length = %#+v\n", len(dst_states))
		/**/
		dst_state := dst_states[(-tree.id)-1]
		transitions = []MicroprogramTransition{
			{
				StateDst:    dst_state,
				Constraints: constraint_prefix,
			},
		}
		return
	}
	// Non-leaf case
	transitions = make([]MicroprogramTransition, 0)
	if len(tree.children) == 1 {
		panic("(Tree) degree one is disallowed. Remember to call tree.CoerceForbidDegreeOne(). ")
	} else if len(tree.children) == 2 {
		// Binary case
		constraint_size := len(constraint_prefix) + len(invertible_constraint_prefix) + 1
		left_child, right_child := tree.children[0], tree.children[1]
		left_constraint := constraintgen.Generate(BoolSort)
		left_constraint_prefix := make([]string, constraint_size)
		copy(left_constraint_prefix, constraint_prefix)
		left_constraint_prefix[len(constraint_prefix)] = left_constraint
		for i, invertible_prefix_constraint := range invertible_constraint_prefix {
			left_constraint_prefix[constraint_size-1-i] = invertible_prefix_constraint
		}
		transitions = append(transitions, left_child.AsMicroprogramTransitionsWithPrefix(
			dst_states,
			constraintgen,
			left_constraint_prefix,
			[]string{},
		)...)
		right_constraint := InvertedConstraintForMicroprogram(left_constraint)
		right_constraint_prefix := make([]string, constraint_size)
		copy(right_constraint_prefix, constraint_prefix)
		right_constraint_prefix[len(constraint_prefix)] = right_constraint
		for i, invertible_prefix_constraint := range invertible_constraint_prefix {
			inverted_constraint := InvertedConstraintForMicroprogram(invertible_prefix_constraint)
			left_constraint_prefix[constraint_size-1-i] = inverted_constraint
		}
		transitions = append(transitions, right_child.AsMicroprogramTransitionsWithPrefix(
			dst_states,
			constraintgen,
			right_constraint_prefix,
			[]string{},
		)...)
	} else {
		panic(
			fmt.Sprintf(
				"(Tree) invalid. Maximum degree must be 2. Got %d. ",
				len(tree.children),
			),
		)
	}
	return
}

func InvertedConstraintForMicroprogram(orig string) (inverted_form string) {
	inverted_form = "@__INVERTED__" + orig
	return
}
