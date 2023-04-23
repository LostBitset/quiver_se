package libsynthetic

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
		avg_ne := rand.Intn(4) + 1
		ErdosRenyiQuiverGivenEdges(n, p, avg_ne)
		fmt.Printf("Generated quiver with n=%d p=%.3f n_quivers=%d.\n", n, p, avg_ne)
	}
}
