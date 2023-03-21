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
	solver_result := sys.CheckSat(
		complete_query_with_ids,
		[]qse.SMTFreeFun[string, string]{}, // Query contains model, which gives definitions
	)
	does_ptr := solver_result.IsSat()
	does = false
	if does_ptr != nil && *does_ptr {
		does = true
	}
	return
}

func MicroprogramConstraintToIdLiteral(
	constraint_raw string, idsrc *qse.IdSource,
) (
	id_literal qse.IdLiteral[string],
) {
	constraint := constraint_raw
	inverted := strings.HasPrefix(constraint, "@__INVERTED__")
	var id_literal_raw qse.Literal[qse.WithId_H[string]]
	if inverted {
		underlying_constraint, _ := strings.CutPrefix(constraint, "@__INVERTED__")
		id_literal_raw = qse.InvertingLiteral(
			qse.WithId_H[string]{
				Value: underlying_constraint,
				Id:    idsrc.Gen(),
			},
		)
	} else {
		id_literal_raw = qse.BufferingLiteral(
			qse.WithId_H[string]{
				Value: constraint,
				Id:    idsrc.Gen(),
			},
		)
	}
	id_literal = qse.IdLiteral[string](id_literal_raw)
	return
}

func (uprgm_gen MicroprogramGenerator) UnitializedAssignment() (model string) {
	sort_values_uninit := map[string]string{
		"Real": "0.0",
		"Bool": "false",
	}
	model = uprgm_gen.UniformAssignmentOfSMTFreeFuns(sort_values_uninit)
	return
}

func (uprgm_gen MicroprogramGenerator) UniformAssignmentOfSMTFreeFuns(
	sort_values map[string]string,
) (
	model string,
) {
	model = uprgm_gen.UniformAssignmentOfSMTFreeFuns(sort_values)
	return
}

func (cgen ConstraintGenerator) UniformAssignmentOfSMTFreeFuns(
	sort_values map[string]string,
) (
	model string,
) {
	var sb strings.Builder
	for _, smt_free_fun := range cgen.SMTFreeFuns() {
		if len(smt_free_fun.Args) != 0 {
			panic("Invalid. Cannot generate a uniform assignment of parametric SMT funs.")
		}
		sb.WriteString(
			StringSMTFreeFun{smt_free_fun}.DefinitionString(
				sort_values[smt_free_fun.Ret],
			),
		)
		sb.WriteRune('\n')
	}
	model = sb.String()
	return
}
