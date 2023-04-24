package libsynthetic

import (
	"fmt"
	"testing"
)

const TEST_ITERATIONS_MICROPROGRAMS = TEST_ITERATIONS * 20

func TestMicroprogramGeneration(t *testing.T) {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	fmt.Printf("Generating %d random microprograms.\n", TEST_ITERATIONS_MICROPROGRAMS)
	for i := 0; i < TEST_ITERATIONS_MICROPROGRAMS; i++ {
		uprgm_gen.RandomMicroprogram()
	}
}

func TestMicroprogramGenerationOnce(t *testing.T) {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	uprgm := uprgm_gen.RandomMicroprogram()
	for k, vList := range uprgm.Transitions {
		for _, v := range vList {
			fmt.Printf("| %v -> %v (trxn)\n", k, v.StateDst)
		}
	}
}
