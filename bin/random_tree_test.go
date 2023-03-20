package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomTreeGeneration(t *testing.T) {
	fmt.Printf("Generating %d random trees.\n", TEST_ITERATIONS)
	for i := 0; i < TEST_ITERATIONS; i++ {
		seq := RandomPruferSequence(rand.Intn(15) + 3)
		seq.ToTree()
	}
}
