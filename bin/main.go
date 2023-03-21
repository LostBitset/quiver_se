package main

import (
	qse "LostBitset/quiver_se/lib"
	"fmt"
)

func main() {
	uprgm := Microprogram{
		top_state:  1,
		fail_state: 2,
		transitions: map[MicroprogramState][]MicroprogramTransition{
			1: {
				{3, []string{"false"}},
				{4, []string{"(= x 7)"}},
			},
			3: {
				{2, []string{"(< y 4)"}},
			},
			4: {
				{3, []string{"(> x 0)", "(> y 0)"}},
				{2, []string{"(= y 99)"}},
			},
		},
		smt_free_funs: []qse.SMTFreeFun[string, string]{
			{Name: "x", Args: []string{}, Ret: "Real"},
			{Name: "y", Args: []string{}, Ret: "Real"},
		},
	}
	n_bugs := uprgm.RunDSE()
	if n_bugs == -1 {
		fmt.Println("RESULT: Program fails immediately. Should not use.")
	} else {
		fmt.Printf("RESULT: Concolic execution found %d bugs.\n", n_bugs)
	}
}
