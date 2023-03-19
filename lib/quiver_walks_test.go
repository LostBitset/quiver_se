package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuiverUpdates(t *testing.T) {
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
		pto(99),
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

func TestQuiverUpdatesInduceCycle(t *testing.T) {
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
		pto(88),
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

func TestQuiverWalks(t *testing.T) {
	var q SimpleQuiver[int, int]
	n1 := q.InsertNodeSimple(7)
	n2 := q.InsertNodeSimple(8)
	n3 := q.InsertNodeSimple(9)
	fail := q.InsertNodeSimple(251)
	q.InsertEdge(n1, n2, 90)
	q.InsertEdge(n1, n2, 60)
	q.InsertEdge(n2, n1, 30)
	q.InsertEdge(n2, n3, 44)
	q.InsertEdge(n1, n3, 55)
	q.InsertEdge(n1, n3, 77)
	q.InsertEdge(n1, fail, 551)
	q.InsertEdge(n2, fail, 551)
	q.InsertEdge(n3, fail, 551)
	update := QuiverUpdate[int, int, *SimpleReversibleAssoc[int, QuiverIndex]]{
		n3,
		q.ParameterizeIndex(n1),
		pto(88),
	}
	walks_chan := make(chan QuiverWalk[int, int])
	q.ApplyUpdateAndEmitWalks(
		walks_chan,
		update,
		q.ParameterizeIndex(n1),
		fail,
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[int]{
			{90, n2},
			{60, n2},
			{55, n3},
			{77, n3},
			{551, fail},
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
			{551, fail},
		},
		q.AllOutneighbors(n3),
	)
	walks := make([][]int, 0)
	for walk_chunked := range walks_chan {
		new_walk := make([]int, 0)
		for _, chunk := range walk_chunked.edges_chunked {
			new_walk = append(new_walk, *chunk...)
		}
		walks = append(walks, new_walk)
	}
	assert.Greater(t, len(walks), 20)
	assert.Less(t, len(walks), 60000)
	assert.Contains(
		t,
		walks,
		[]int{
			60, 44, 88, 90, 44, 551,
		},
	)
	assert.Subset(
		t,
		walks,
		[][]int{
			{60, 44, 88, 90, 44, 551},
			{90, 44, 88, 77, 551},
			{60, 44, 88, 77, 551},
			{90, 44, 88, 551},
			{60, 44, 88, 551},
			{77, 88, 90, 551},
			{77, 88, 60, 551},
			{77, 88, 551},
		},
	)
}
