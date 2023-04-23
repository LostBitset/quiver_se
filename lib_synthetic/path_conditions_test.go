package libsynthetic

import (
	"fmt"
	"testing"

	qse "github.com/LostBitset/quiver_se/lib"

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
		fails, pc := uprgm.ExecuteGetPathCondition(initial_model, false)
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

func TestSpecificPathCondition(t *testing.T) {
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
	test_model := `
	(define-fun x () Real    7.0)
	(define-fun y () Real    0.0)
	`
	_, pc := uprgm.ExecuteGetPathCondition(test_model, false)
	assert.Equal(t, 7, len(pc))
}

func TestSpecificFailureCheck(t *testing.T) {
	uprgm := Microprogram{
		StateTop:  1,
		StateFail: 2,
		Transitions: map[MicroprogramState][]MicroprogramTransition{
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
	test_model := `
	(define-fun x () Real    7.0)
	(define-fun y () Real    0.0)
	`
	fails, _ := uprgm.ExecuteGetPathCondition(test_model, false)
	assert.False(t, fails)
}

func TestModelSatisfiesConstraint(t *testing.T) {
	test_model := `
	(define-fun x () Real 0.0)
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
