package main

import (
	qse "LostBitset/quiver_se/lib"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (uprgm Microprogram) RunDSE() (n_bugs int) {
	var backing_idsrc qse.IdSource
	idsrc := &backing_idsrc
	n_bugs = 0
	model := uprgm.UnitializedAssignment()
	imm_failure, imm_pc := uprgm.ExecuteGetPathCondition(model)
	if imm_failure {
		log.Info("[bin:dse_impl] Immediate failure. ")
		n_bugs = -1 // Don't use results when the program failed immediately
		return
	}
	// The two main variables for the concolic execution algorithm:
	alt_stack := make([]uint, 0)
	desired_path := make([]qse.IdLiteral[string], 0)
	// Set them up using the first path condition
	stack_setup_index := 0
setupAltStackAndPathLoop:
	for _, constraint := range imm_pc {
		if strings.HasPrefix(constraint, "@__RAW__;;@RICHPC:") {
			continue setupAltStackAndPathLoop
		}
		alt_stack = append(alt_stack, uint(stack_setup_index))
		desired_path = append(
			desired_path,
			MicroprogramConstraintToIdLiteral(
				constraint,
				idsrc,
			))
		stack_setup_index++
	}
	// Main loop
mainDSESearchAlternativesLoop:
	for len(alt_stack) > 0 {
		alt_stack_pop_index := len(alt_stack) - 1
		inv_index := alt_stack[alt_stack_pop_index]
		qse.SpliceOutReclaim(&alt_stack, alt_stack_pop_index)
		orig_expr := desired_path[inv_index]
		desired_path = desired_path[:inv_index]
		desired_path = append(desired_path, qse.IdLiteral[string]{
			Value: orig_expr.Value,
			Eq:    !orig_expr.Eq,
		})
		new_model_ptr := uprgm.SolveForInputZ3(desired_path, idsrc)
		if new_model_ptr == nil {
			continue mainDSESearchAlternativesLoop
		}
		new_model := *new_model_ptr
		fails, pc := uprgm.ExecuteGetPathCondition(new_model)
		if fails {
			n_bugs++
		}
		desired_path = make([]qse.IdLiteral[string], len(pc))
		for i, path_condition_elem := range pc {
			desired_path[i] = MicroprogramConstraintToIdLiteral(
				path_condition_elem,
				idsrc,
			)
		}
	updateAltStackGivenDiscoveredPathLoop:
		for index := inv_index + 1; index < uint(len(desired_path)); index++ {
			curr_path_item := desired_path[index]
			if strings.HasPrefix(curr_path_item.Value.Value, "@__RAW__;;@RICHPC:") {
				continue updateAltStackGivenDiscoveredPathLoop
			}
			alt_stack = append(alt_stack, index)
		}
	}
	return
}

func (uprgm Microprogram) SolveForInputZ3(
	constraints []qse.IdLiteral[string],
	idsrc *qse.IdSource,
) (
	model_ptr *string,
) {
	sys := qse.SMTLibv2StringSystem{Idsrc: *idsrc}
	sctx := sys.CheckSat(constraints, uprgm.smt_free_funs)
	is_sat_ptr := sctx.IsSat()
	if is_sat_ptr != nil && *is_sat_ptr {
		model_ptr = sctx.GetModel()
	} else {
		model_ptr = nil
	}
	return
}
