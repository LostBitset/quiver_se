package main

import (
	"fmt"
	"os"

	qse "github.com/LostBitset/quiver_se/lib"
)

func main() {
	seirBytes, errOpen := os.ReadFile("demo.seir")
	if errOpen != nil {
		panic(errOpen)
	}
	source := string(seirBytes)
	prgm := SeirPrgm{
		source: source,
		smt_free_funs: []qse.SMTFreeFun[string, string]{
			{Name: "A", Args: []string{}, Ret: "Bool"},
			{Name: "B", Args: []string{}, Ret: "Int"},
		},
		names_source_symb: func(smt_name string) string {
			return "symb_" + smt_name
		},
	}
	spc := prgm.PerformQuery(prgm.UninitializedAssignment())
	pc := FlattenSpc(spc)
	fmt.Printf("%#+v\n", pc)
}
