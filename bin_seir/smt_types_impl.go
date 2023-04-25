package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

func UninitializedAssignedSFF(sff qse.SMTFreeFun[string, string]) (aver AssignedSMTValue) {
	if len(sff.Args) > 0 {
		panic("cannot create uninitialized value for smt function")
	}
	aver = AssignedSMTValue{
		smt_free_fun: sff,
		value_repr:   UninitializedValueForSMTSort(sff.Ret),
	}
	return
}

func UninitializedValueForSMTSort(sort string) (value_repr string) {
	switch sort {
	case "Int":
		value_repr = "0"
	case "Bool":
		value_repr = "false"
	default:
		panic("cannot create unitialized value for sort " + sort)
	}
	return
}

func (sp SeirPrgm) UninitializedAssignment() (assignment []AssignedSMTValue) {
	assignment = make([]AssignedSMTValue, len(sp.smt_free_funs))
	for i, sff := range sp.smt_free_funs {
		assignment[i] = UninitializedAssignedSFF(sff)
	}
	return
}
