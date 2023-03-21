package main

import (
	qse "LostBitset/quiver_se/lib"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_ITERATIONS_PC = TEST_ITERATIONS / 16

func TestPathConditions(t *testing.T) {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	fail_count := 0
	fmt.Printf("Computing path conditions for %d random microprograms.\n", TEST_ITERATIONS_PC)
	for i := 0; i < TEST_ITERATIONS_PC; i++ {
		uprgm := uprgm_gen.RandomMicroprogram()
		initial_model := uprgm.UnitializedAssignment()
		fails, pc := uprgm.ExecuteGetPathCondition(initial_model)
		fmt.Printf("Computed path condition of length %d.\n", len(pc))
		if fails {
			fmt.Println("Path condition led to failure.")
			fail_count++
		} else {
			fmt.Println("Path condition did not lead to failure. ")
		}
		fmt.Printf("So far: %d / %d PCs led to failure.", fail_count, i+1)
	}
}

func TestModelStatisfiesConstraint(t *testing.T) {
	test_model := `
	(define-fun x () Real 0)
	`
	test_constraints := []string{"(= x 7)"}
	uprgm := Microprogram{
		smt_free_funs: []qse.SMTFreeFun[string, string]{
			{Name: "x", Args: []string{}, Ret: "Real"},
		},
	}
	assert.False(
		t,
		uprgm.ModelSatisfiesConstraints(test_model, test_constraints),
	)
}
