package main

import (
	"strings"

	qse "github.com/LostBitset/quiver_se/lib"
)

func StringConstraintToIdLiteral(
	constraint_raw string, idsrc *qse.IdSource,
) (
	id_literal qse.IdLiteral[string],
) {
	constraint := constraint_raw
	inverted := false
	for strings.HasPrefix(constraint, "@__INVERTED__") {
		constraint, _ = strings.CutPrefix(constraint, "@__INVERTED__")
		inverted = !inverted
	}
	var id_literal_raw qse.Literal[qse.WithId_H[string]]
	if inverted {
		id_literal_raw = qse.InvertingLiteral(
			qse.WithId_H[string]{
				Value: constraint,
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

func (uprgm SeirPrgm) UnitializedAssignment() (model string) {
	sort_values_uninit := map[string]string{
		"Real": "0.0",
		"Bool": "false",
	}
	model = uprgm.UniformAssignmentOfSMTFreeFuns(sort_values_uninit)
	return
}

func (uprgm SeirPrgm) UniformAssignmentOfSMTFreeFuns(
	sort_values map[string]string,
) (
	model string,
) {
	model = UniformAssignmentOfSMTFreeFuns(uprgm.smt_free_funs, sort_values)
	return
}

func UniformAssignmentOfSMTFreeFuns(
	smt_free_funs []qse.SMTFreeFun[string, string],
	sort_values map[string]string,
) (
	model string,
) {
	var sb strings.Builder
	for _, smt_free_fun := range smt_free_funs {
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
