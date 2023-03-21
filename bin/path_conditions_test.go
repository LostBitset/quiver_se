package main

import (
	"fmt"
	"testing"
)

const TEST_ITERATIONS_PC = TEST_ITERATIONS / 16

func TestPathConditions(t *testing.T) {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	initial_model := uprgm_gen.UnitializedAssignment()
	fail_count := 0
	fmt.Printf("Computing path conditions for %d random microprograms.\n", TEST_ITERATIONS_PC)
	for i := 0; i < TEST_ITERATIONS_PC; i++ {
		uprgm := uprgm_gen.RandomMicroprogram()
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
