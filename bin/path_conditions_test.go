package main

import (
	"fmt"
	"testing"
)

const TEST_ITERATIONS_PC = TEST_ITERATIONS / 16

func TestPathConditions(t *testing.T) {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	initial_model := uprgm_gen.UnitializedAssignment()
	fmt.Printf("Computing path conditions for %d random microprograms.\n", TEST_ITERATIONS_PC)
	for i := 0; i < TEST_ITERATIONS_PC; i++ {
		uprgm := uprgm_gen.RandomMicroprogram()
		uprgm.ExecuteGetPathCondition(initial_model)
	}
}
