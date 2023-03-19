package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleReversibleAssoc(t *testing.T) {
	ra := NewSimpleRA[string, int]()
	ra.Insert("hi", 42)
	assert.Equal(
		t,
		42,
		*ra.FwdLookup("hi"),
	)
	assert.Equal(
		t,
		[]string{
			"hi",
		},
		ra.RevLookup(42),
	)
}

func TestQuiver(t *testing.T) {
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
}

func TestCyclicQuiver(t *testing.T) {
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
	q.InsertEdge(n3, n1, 88)
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
