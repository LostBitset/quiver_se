package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPruferRandomTreeGeneration(t *testing.T) {
	fmt.Printf("Generating %d random trees.\n", TEST_ITERATIONS)
	for i := 0; i < TEST_ITERATIONS; i++ {
		seq := RandomPruferSequence(rand.Intn(15) + 3)
		fmt.Println("Generated random Prüfer Sequence:")
		fmt.Printf("prüfer seq #%d: %#+v\n", i, seq)
		tree := seq.ToTree()
		fmt.Println("Converted Prüfer Sequence to tree.")
		fmt.Printf("size of Prüfer tree #%d: %#+v\n", i, tree.ComputeSize())
	}
}
