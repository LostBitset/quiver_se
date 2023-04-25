package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

type AssignedSMTValue struct {
	smt_free_fun qse.SMTFreeFun[string, string]
	value_repr   string
}
