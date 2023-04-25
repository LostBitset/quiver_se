package main

import (
	"fmt"

	qse "github.com/LostBitset/quiver_se/lib"
)

func main() {
	test_sp := SeirPrgm{
		source: "(scope (+ symb_X 1))",
		smt_free_funs: []qse.SMTFreeFun[string, string]{
			{Name: "X", Args: []string{}, Ret: "Int"},
		},
		names_source_symb: func(smt_name string) string {
			return "symb_" + smt_name
		},
	}
	jsonstr := test_sp.MakeQueryJson(test_sp.UninitializedAssignment())
	fmt.Println(string(jsonstr))
}
