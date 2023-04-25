package main

import (
	"testing"

	qse "github.com/LostBitset/quiver_se/lib"
	"github.com/stretchr/testify/assert"
)

func TestMakeQueryJson(t *testing.T) {
	test_sp := SeirPrgm{
		source: "(scope (+ symb_X {int 1}))",
		smt_free_funs: []qse.SMTFreeFun[string, string]{
			{Name: "X", Args: []string{}, Ret: "Int"},
		},
		names_source_symb: func(smt_name string) string {
			return "symb_" + smt_name
		},
	}
	jsonstr := test_sp.MakeQueryJson(test_sp.UninitializedAssignment())
	str := string(jsonstr)
	assert.Equal(
		t,
		str,
		`{"languages":{"smt":"smtlib_2va","source":"seir"},"source":"(scope (+ symb_X {int 1}))","vars":[{"assigned_value":"0","smt_name":"X","sort":"Int","source_name":"symb_X"}]}`,
	)
}
