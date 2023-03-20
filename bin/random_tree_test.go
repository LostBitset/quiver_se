package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomTreeGeneration(t *testing.T) {
	for i := 0; i < 1000; i++ {
		seq := RandomPruferSequence(rand.Intn(15) + 3)
		fmt.Printf("seq: %#+v\n", seq.sequence)
		tree := seq.ToTree()
		fmt.Printf("size: %#+v\n", tree.ComputeSize())
	}
}
