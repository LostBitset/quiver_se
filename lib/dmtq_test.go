package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMTQ(t *testing.T) {
	var q Quiver[uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex]]
	n1_container := NewDMT[uint32_H, QuiverIndex]()
	n2_container := NewDMT[uint32_H, QuiverIndex]()
	n1 := q.insert_node(uint32_H{1}, &n1_container)
	n2 := q.insert_node(uint32_H{2}, &n2_container)
	q.insert_edge(
		n1,
		n2,
		StdlibMapToPHashMap(
			map[Literal[uint32_H]]struct{}{
				{uint32_H{77}, true}: {},
				{uint32_H{88}, true}: {},
			},
		),
	)
	q.insert_edge(
		n1,
		n2,
		StdlibMapToPHashMap(
			map[Literal[uint32_H]]struct{}{
				{uint32_H{77}, true}: {},
				{uint32_H{88}, false}: {},
			},
		),
	)
	assert.ElementsMatch(
		t,
		[]Neighbor[map[Literal[uint32_H]]struct{}]{
			{
				map[Literal[uint32_H]]struct{}{
					{uint32_H{77}, true}: {},
				},
				n2,
			},
		},
		q.all_outneighbors(n1),
	)
}

