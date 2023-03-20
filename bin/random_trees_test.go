package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomTreeGeneration(t *testing.T) {
	fmt.Printf("Generating %d random trees (with a predetermined number of leaves).", TEST_ITERATIONS)
	for i := 0; i < TEST_ITERATIONS; i++ {
		n_nonleaf := rand.Intn(40) + 2
		n_leaf := rand.Intn(30) + 2
		PruferEvenFinalRandomTree(n_nonleaf, n_leaf)
	}
}
