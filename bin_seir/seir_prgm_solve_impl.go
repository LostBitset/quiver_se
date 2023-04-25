package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

func (sp SeirPrgm) SolveForAssignmentsZ3(
	constraints []qse.IdLiteral[string],
) (
	assignments_ptr *[]AssignedSMTValue,
) {
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{Idsrc: idsrc}
	sctx := sys.CheckSat(constraints, sp.smt_free_funs)
	is_sat_ptr := sctx.IsSat()
	if is_sat_ptr != nil && *&is_sat_ptr {
		model_ptr := sctx.GetModel()
		model := *model_ptr
		assignments := ParseZ3ModelString(model)
		assignments_ptr = &assignments
	}
}
