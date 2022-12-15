package qse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuiverCreation(t *testing.T) {
	var q SimpleQuiver[int, int]
	n1 := q.insert_node_simple(7)
	n2 := q.insert_node_simple(8)
	n3 := q.insert_node_simple(9)
	fmt.Printf("q.arena: %v\n", q.arena)
	q.insert_edge(n1, n2, 90)
	q.insert_edge(n1, n2, 60)
	q.insert_edge(n2, n1, 30)
	q.insert_edge(n2, n3, 44)
	q.insert_edge(n1, n3, 55)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{90, n2},
			{60, n2},
			{55, n3},
		},
		q.all_outneighbors(n1),
	)
}
