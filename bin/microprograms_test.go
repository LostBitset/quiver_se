package main

import (
	"fmt"
	"testing"
)

const TEST_ITERATIONS_MICROPROGRAMS = TEST_ITERATIONS * 2

func TestMicroprogramGeneration(t *testing.T) {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	fmt.Printf("Generating %d random microprograms.\n", TEST_ITERATIONS_MICROPROGRAMS)
	for i := 0; i < TEST_ITERATIONS_MICROPROGRAMS; i++ {
		uprgm_gen.RandomMicroprogram()
	}
}
