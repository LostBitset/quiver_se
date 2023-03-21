package main

import (
	"fmt"
	"testing"
)

const TEST_ITERATIONS_PC = TEST_ITERATIONS / 16

func TestPathConditions(t *testing.T) {
	initial_model := `
	(define-fun var_0 () Real 0)
	(define-fun var_1 () Real 0)
	(define-fun var_2 () Real 0)
	(define-fun var_3 () Real 0)
	`
	uprgm_gen := BuildTestingMicroprogramGenerator()
	fmt.Printf("Computing path conditions for %d random microprograms.\n", TEST_ITERATIONS_PC)
	for i := 0; i < TEST_ITERATIONS_PC; i++ {
		uprgm := uprgm_gen.RandomMicroprogram()
		uprgm.ExecuteGetPathCondition(initial_model)
	}
}
