package main

import (
	"fmt"
	"strings"

	qse "github.com/LostBitset/quiver_se/lib"
)

func (sp SeirPrgm) RunDSE(
	no_transition bool,
	max_iters int, // for no limit, use -1
) {
	assignment := sp.UninitializedAssignment()
	imm_spc, imm_failure := sp.PerformQuery(assignment)
	imm_pc := FlattenSpc(imm_spc)
	if imm_failure {
		panic("Something went wrong, program failed on uninitialized assignment.")
	}
	alt_stack := make([]uint, 0)
	desired_path := make([]qse.IdLiteral[string], 0)
	stack_setup_index := uint(0)
setupAltStackAndPathLoop:
	for _, item := range imm_pc.items {
		if strings.HasPrefix(item.Value.Value, "@__RAW__;;@") {
			continue setupAltStackAndPathLoop
		}
		alt_stack = append(alt_stack, stack_setup_index)
		desired_path = append(desired_path, item)
		stack_setup_index++
	}
	n_iters := 0
	// Main DSE loop
mainDSESearchAlternativesLoop:
	for (len(alt_stack) > 0) && (n_iters != max_iters) {
		n_iters++
		alt_stack_pop_index := len(alt_stack) - 1
		inv_index := alt_stack[alt_stack_pop_index]
		qse.SpliceOutReclaim(&alt_stack, alt_stack_pop_index)
		orig_expr := desired_path[inv_index]
		desired_path = desired_path[:inv_index]
		desired_path = append(desired_path, qse.IdLiteral[string]{
			Value: orig_expr.Value,
			Eq:    !orig_expr.Eq,
		})
		new_assignments_ptr := sp.SolveForAssignmentsZ3(desired_path)
		if new_assignments_ptr == nil {
			continue mainDSESearchAlternativesLoop
		}
		spc, fails := sp.PerformQuery(*new_assignments_ptr)
		if fails {
			fmt.Println("FOUND A BUG!!!!!!")
		}
		pc := FlattenSpc(spc)
		desired_path = make([]qse.IdLiteral[string], len(pc.items))
		for i, item := range pc.items {
			desired_path[i] = item
		}
	updateAltStackGivenDiscoveredPathLoop:
		for i := inv_index + 1; i < uint(len(desired_path)); i++ {
			curr_path_item := desired_path[i]
			if strings.HasPrefix(curr_path_item.Value.Value, "@__RAW__;;@") {
				continue updateAltStackGivenDiscoveredPathLoop
			}
			alt_stack = append(alt_stack, i)
		}
	}
}
