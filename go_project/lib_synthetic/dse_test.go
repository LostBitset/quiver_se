package libsynthetic

import (
	"fmt"
	"testing"

	qse "github.com/LostBitset/quiver_se/lib"

	"github.com/stretchr/testify/assert"
)

func TestDSESimple(t *testing.T) {
	uprgm := Microprogram{
		StateTop:  1,
		StateFail: 2,
		Transitions: map[MicroprogramState][]MicroprogramTransition{
			1: {
				{5, []string{"false"}},
				{5, []string{"(and (= x 8) (= x 9))"}},
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
	assert.Equal(t, 1, n_bugs)
}
