package main

import (
	qse "LostBitset/quiver_se/lib"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSESimple(t *testing.T) {
	uprgm := Microprogram{
		top_state:  1,
		fail_state: 2,
		transitions: map[MicroprogramState][]MicroprogramTransition{
			1: {
				{3, []string{"false"}},
				{4, []string{"(= x 7)"}},
			},
			3: {
				{2, []string{"(= y 88)"}},
			},
			4: {
				{3, []string{"(> x 0)"}},
				{2, []string{"@__INVERTED__(> x 0)", "(= y 99)"}},
			},
		},
		smt_free_funs: []qse.SMTFreeFun[string, string]{
			{Name: "x", Args: []string{}, Ret: "Real"},
			{Name: "y", Args: []string{}, Ret: "Real"},
		},
	}
	fmt.Println("INITIAL ASSIGNMENT: ")
	fmt.Println(uprgm.UnitializedAssignment())
	n_bugs := uprgm.RunDSE()
	assert.Equal(t, 2, n_bugs)
}
