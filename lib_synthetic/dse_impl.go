package libsynthetic

import (
	qse "github.com/LostBitset/quiver_se/lib"
	"fmt"
	"hash/fnv"
	"strings"
)

func (uprgm Microprogram) RunDSE() (n_bugs int) {
	n_bugs = 0
	bug_signal := make(chan uint32)
	go uprgm.RunDSEContinuously(
		bug_signal, false, nil, false, -1, uprgm.top_state,
	)
	for range bug_signal {
		n_bugs++
	}
	return
}

func (uprgm Microprogram) RunDSEContinuously(
	bug_signal chan uint32,
	emit_pcs bool,
	out_pcs *chan PathConditionResult,
	no_transition bool,
	max_iters int, // -1 for no limit
	top_state MicroprogramState,
) {
	var backing_idsrc qse.IdSource
	idsrc := &backing_idsrc
	model := uprgm.UnitializedAssignment()
	imm_failure, imm_pc := uprgm.ExecuteGetPathConditionFrom(
		model, top_state, no_transition, PC_REC_LIMIT,
	)
	if imm_failure {
		panic("[bad-input-panic] [bin:dse_impl] Immediate failure. ")
	}
	if emit_pcs {
		*out_pcs <- PathConditionResult{imm_pc, imm_failure}
	}
	// The two main variables for the concolic execution algorithm:
	alt_stack := make([]uint, 0)
	desired_path := make([]qse.IdLiteral[string], 0)
	// Also need to avoid duplicate detections
	detected_model_hashes := make(map[uint32]struct{})
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
	n_iters := 0
mainDSESearchAlternativesLoop:
	for (len(alt_stack) > 0) || (n_iters == max_iters) {
		n_iters++
		fmt.Printf("[STATUS-DSE] %d / %d iters\n", n_iters, max_iters)
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
		new_model := FilterModelFromZ3(*new_model_ptr)
		fails, pc := uprgm.ExecuteGetPathConditionFrom(
			new_model, top_state, no_transition, PC_REC_LIMIT,
		)
		if emit_pcs {
			saved_pc := make([]string, len(pc))
			copy(saved_pc, pc)
			*out_pcs <- PathConditionResult{saved_pc, fails}
		}
		if fails {
			hasher := fnv.New32a()
			hasher.Write([]byte(new_model))
			new_model_hash := hasher.Sum32()
			if _, ok := detected_model_hashes[new_model_hash]; !ok {
				bug_signal <- new_model_hash
				detected_model_hashes[new_model_hash] = struct{}{}
				if no_transition {
					fmt.Println("[result-note] notransition=true")
				}
				fmt.Println("[result] [bin:dse_impl/RunDSE] Found a failure-inducing input:")
				fmt.Println(new_model)
			}
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
	close(bug_signal)
	if emit_pcs {
		close(*out_pcs)
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
