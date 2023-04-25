package main

import (
	"fmt"

	qse "github.com/LostBitset/quiver_se/lib"
)

func (sp SeirPrgm) SolveForAssignmentsZ3(
	constraints []qse.IdLiteral[string],
) (
	assignments_ptr *[]AssignedSMTValue,
) {
	fmt.Println(constraints)
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{Idsrc: idsrc}
	sctx := sys.CheckSat(constraints, sp.smt_free_funs)
	is_sat_ptr := sctx.IsSat()
	if is_sat_ptr != nil && *is_sat_ptr {
		model_ptr := sctx.GetModel()
		model := *model_ptr
		fmt.Println("::: GOT MODEL :::")
		fmt.Println(model)
		fmt.Println("::: END MODEL :::")
		assignments := ParseZ3ModelString(model)
		assignments_ptr = &assignments
	}
	return
}

func ParseZ3ModelString(model string) (assignments []AssignedSMTValue) {
	panic("TODO TODO TODO")
}
