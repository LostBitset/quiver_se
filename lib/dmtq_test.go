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
				{uint32_H{77}, true}:  {},
				{uint32_H{88}, false}: {},
			},
		),
	)
	n1_outneighbors := q.all_outneighbors(n1)
	n1_outneighbors_stdlib_map := make([]Neighbor[map[Literal[uint32_H]]struct{}], 0)
	for _, neighbor := range n1_outneighbors {
		neighbor_stdlib_map := Neighbor[map[Literal[uint32_H]]struct{}]{
			neighbor.via_edge.ToStdlibMap(),
			neighbor.dst,
		}
		n1_outneighbors_stdlib_map = append(n1_outneighbors_stdlib_map, neighbor_stdlib_map)
	}
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
		n1_outneighbors_stdlib_map,
	)
}
