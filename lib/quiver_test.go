package qse

import (
	"testing"
)

func TestQuiverCreation(t *testing.T) {
	var q SimpleQuiver[int, int]
	n1 := q.insert_node(7)
	n2 := q.insert_node(8)
	n3 := q.insert_node(9)
	q.insert_edge(n1, n2, 90)
	q.insert_edge(n1, n2, 60)
	q.insert_edge(n2, n1, 30)
	q.insert_edge(n2, n3, 44)
}
