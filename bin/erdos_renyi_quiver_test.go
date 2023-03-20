package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestErdosRenyiQuiverBasic(t *testing.T) {
	fmt.Printf("Generating %d random quivers using the Erdős-Rényi model.\n", TEST_ITERATIONS)
	for i := 0; i < TEST_ITERATIONS; i++ {
		n := rand.Intn(20) + 3
		p := (rand.Float64() * 0.8) + 0.1
		avg_ne := rand.Intn(20) + 3
		ErdosRenyiQuiverGivenEdges(n, p, avg_ne)
	}
}
