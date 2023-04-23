package libsynthetic

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomTreeGeneration(t *testing.T) {
	fmt.Printf(
		"Generating %d random trees (with a predetermined number of leaves).\n",
		TEST_ITERATIONS,
	)
	for i := 0; i < TEST_ITERATIONS; i++ {
		n_nonleaf := rand.Intn(40) + 2
		n_leaf := rand.Intn(30) + 2
		fmt.Printf(
			"Generating random tree, starting with %d non-leaf nodes and %d leaf nodes...\n",
			n_nonleaf,
			n_leaf,
		)
		tree := PruferEvenFinalRandomTree(n_nonleaf, n_leaf)
		fmt.Printf("size of final random tree: %d.\n", tree.ComputeSize())

	}
}

func TestRandomTreeCoerceMaxDegreeOnce(t *testing.T) {
	tree := PruferEvenFinalRandomTree(22, 17)
	tree.CoerceToMaxDegree(2)
	actual_max_degree := tree.MaxDegree()
	assert.LessOrEqual(t, actual_max_degree, 2)
}
