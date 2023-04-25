package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

type SeirPrgm struct {
	source            string
	smt_free_funs     []qse.SMTFreeFun[string, string]
	names_source_symb func(string) string
}
