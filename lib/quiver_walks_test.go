package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuiverUpdatesAcyclic(t *testing.T) {
	var q SimpleQuiver[int, int]
	n1 := q.InsertNodeSimple(7)
	n2 := q.InsertNodeSimple(8)
	n3 := q.InsertNodeSimple(9)
	q.InsertEdge(n1, n2, 90)
	q.InsertEdge(n1, n2, 60)
	q.InsertEdge(n2, n1, 30)
	q.InsertEdge(n2, n3, 44)
	q.InsertEdge(n1, n3, 55)
	q.InsertEdge(n1, n3, 77)
	update := QuiverUpdate[int, int, *SimpleReversibleAssoc[int, QuiverIndex]]{
		n1,
		q.ParameterizeIndex(n3),
		99,
	}
	q.ApplyUpdate(update)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{90, n2},
			{60, n2},
			{55, n3},
			{77, n3},
			{99, n3},
		},
		q.AllOutneighbors(n1),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{44, n2},
			{55, n1},
			{77, n1},
			{99, n1},
		},
		q.AllInneighbors(n3),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{},
		q.AllOutneighbors(n3),
	)
}

func TestQuiverUpdatesCyclic(t *testing.T) {
	var q SimpleQuiver[int, int]
	n1 := q.InsertNodeSimple(7)
	n2 := q.InsertNodeSimple(8)
	n3 := q.InsertNodeSimple(9)
	q.InsertEdge(n1, n2, 90)
	q.InsertEdge(n1, n2, 60)
	q.InsertEdge(n2, n1, 30)
	q.InsertEdge(n2, n3, 44)
	q.InsertEdge(n1, n3, 55)
	q.InsertEdge(n1, n3, 77)
	update := QuiverUpdate[int, int, *SimpleReversibleAssoc[int, QuiverIndex]]{
		n3,
		q.ParameterizeIndex(n1),
		88,
	}
	q.ApplyUpdate(update)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{90, n2},
			{60, n2},
			{55, n3},
			{77, n3},
		},
		q.AllOutneighbors(n1),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{44, n2},
			{55, n1},
			{77, n1},
		},
		q.AllInneighbors(n3),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{88, n1},
		},
		q.AllOutneighbors(n3),
	)
}

func TestQuiverWalksAcyclic(t *testing.T) {
	var q SimpleQuiver[int, int]
	n1 := q.InsertNodeSimple(7)
	n2 := q.InsertNodeSimple(8)
	n3 := q.InsertNodeSimple(9)
	q.InsertEdge(n1, n2, 90)
	q.InsertEdge(n1, n2, 60)
	q.InsertEdge(n2, n1, 30)
	q.InsertEdge(n2, n3, 44)
	q.InsertEdge(n1, n3, 55)
	q.InsertEdge(n1, n3, 77)
	update := QuiverUpdate[int, int, *SimpleReversibleAssoc[int, QuiverIndex]]{
		n1,
		q.ParameterizeIndex(n3),
		99,
	}
	walks_chan := make(chan QuiverWalk[int, int])
	q.ApplyUpdateAndEmitWalks(
		walks_chan,
		update,
		q.ParameterizeIndex(n1),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{90, n2},
			{60, n2},
			{55, n3},
			{77, n3},
			{99, n3},
		},
		q.AllOutneighbors(n1),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{44, n2},
			{55, n1},
			{77, n1},
			{99, n1},
		},
		q.AllInneighbors(n3),
	)
	walks := make([]QuiverWalk[int, int], 0)
	for walk := range walks_chan {
		walks = append(walks, walk)
	}
	assert.True(t, len(walks) == -1)
}
