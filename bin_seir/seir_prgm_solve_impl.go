package main

import (
	"fmt"
	"strings"

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
	if is_sat_ptr != nil && *is_sat_ptr {
		model_ptr := sctx.GetModel()
		model := *model_ptr
		fmt.Println("::: GOT MODEL :::")
		fmt.Println(model)
		fmt.Println("::: END MODEL :::")
		assignments := ParseZ3ModelString(model)
		fmt.Println("::: GOT ASSIGNMENTS :::")
		for _, assignment := range assignments {
			fmt.Println(assignment)
		}
		fmt.Println("::: END ASSIGNMENTS :::")
		assignments_ptr = &assignments
	}
	return
}

// NOT GENERAL PURPOSE AT ALL
// ONLY WORKS WHEN DEFINE LINE AND VALUE ARE ONE SEPERATE ADJ LINE EACH
// VERY DEPENDENT ON FORMATTING, NOT AN ACTUAL SEXPR PARSER!!!!
func ParseZ3ModelString(model string) (assignments []AssignedSMTValue) {
	lines := strings.Split(model, "\n")
	defs := make(map[string]qse.SMTFreeFun[string, string])
	values := make(map[string]string)
	last_smt_name := "<name:unreachable>"
buildDefsAndValuesLoop:
	for _, line_spaces := range lines {
		if !strings.HasPrefix(line_spaces, "  ") {
			continue buildDefsAndValuesLoop
		}
		line := strings.TrimSpace(line_spaces)
		if strings.HasPrefix(line, "(define-fun ") {
			after, _ := strings.CutPrefix(line, "(define-fun ")
			fields := strings.Fields(after)
			if len(fields) != 3 {
				panic("There should be three fields in SMT model (name, args, ret).")
			}
			if fields[1] != "()" {
				panic("Cannot parse SMT model containing functions. ")
			}
			smt_name := fields[0]
			defs[smt_name] = qse.SMTFreeFun[string, string]{
				Name: smt_name,
				Args: []string{},
				Ret:  fields[2],
			}
			last_smt_name = smt_name
		} else if strings.HasSuffix(line, ")") {
			before, _ := strings.CutSuffix(line, ")")
			values[last_smt_name] = before
		}
	}
	assignments = make([]AssignedSMTValue, 0)
buildAssignmentsLoop:
	for key := range values {
		if _, ok := defs[key]; !ok {
			continue buildAssignmentsLoop
		}
		if strings.HasPrefix(key, "ga_") {
			continue buildAssignmentsLoop
		}
		assignments = append(assignments, AssignedSMTValue{
			smt_free_fun: defs[key],
			value_repr:   values[key],
		})
	}
	return
}
