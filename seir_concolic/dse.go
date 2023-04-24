package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

type PathConditionResult struct {
	pc    []string
	fails bool
}

func SeirStateTop() (top SeirEventState) {
	top = SeirEventState{"#reserved__top__"}
	return
}

func SeirStateFail() (fail SeirEventState) {
	fail = SeirEventState{"#reserved__fail__"}
	return
}

type SeirEventState struct {
	Name string
}

type SeirPrgm struct {
	source        string
	smt_free_funs []qse.SMTFreeFun[string, string]
}
