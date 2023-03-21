package main

import (
	qse "LostBitset/quiver_se/lib"
	"strings"
)

func (uprgm Microprogram) ExecuteGetPathCondition(model string) (pc []string) {
	pc = uprgm.ExecuteGetPathConditionFrom(model, uprgm.top_state)
	return
}

func (uprgm Microprogram) ExecuteGetPathConditionFrom(
	model string, state MicroprogramState,
) (
	pc []string,
) {
	transitions := uprgm.transitions[state]
	pc = make([]string, 0)
selectTransitionLoop:
	for _, transition := range transitions {
		if uprgm.ModelSatisfiesConstraints(model, transition.constraints) {
			pc = append(pc, transition.constraints...)
			pc = append(pc, uprgm.ExecuteGetPathConditionFrom(
				model, transition.dst_state,
			)...)
			break selectTransitionLoop
		}
	}
	return
}

func (uprgm Microprogram) ModelSatisfiesConstraints(model string, constraints []string) (does bool) {
	complete_query := make([]string, len(constraints)+1)
	for i := range complete_query {
		if i == 0 {
			complete_query[i] = "@__RAW__" + model
		} else {
			complete_query[i] = constraints[i-1]
		}
	}
	var idsrc qse.IdSource
	complete_query_with_ids := make([]qse.IdLiteral[string], len(complete_query))
	for i, part := range complete_query {
		complete_query_with_ids[i] = MicroprogramConstraintToIdLiteral(part, &idsrc)
	}
	sys := qse.SMTLibv2StringSystem{Idsrc: idsrc}
	solver_result := sys.CheckSat(complete_query_with_ids, uprgm.smt_free_funs)
	does_ptr := solver_result.IsSat()
	does = false
	if does_ptr != nil && *does_ptr {
		does = true
	}
	return
}

func MicroprogramConstraintToIdLiteral(
	constraint string, idsrc *qse.IdSource,
) (
	id_literal qse.IdLiteral[string],
) {
	raw_constraint := constraint
	constraint_prefixes_reversed := make([]string, 0)
	for strings.HasPrefix(raw_constraint, "@__") {
		raw_constraint, new_prefix := CutConstraintMarkerPrefix(raw_constraint)
		constraint_prefixes_reversed = append(constraint_prefixes_reversed, new_prefix)
	}
	n_constraint_prefixes := len(constraint_prefixes_reversed)
	var constraint_prefixes_sb strings.Builder
	for i := range constraint_prefixes_reversed {
		new_prefix := constraint_prefixes_reversed[n_constraint_prefixes-i-1]
		constraint_prefixes_sb.WriteString(new_prefix)
	}
	constraint_prefixes := constraint_prefixes_sb.String()

}
