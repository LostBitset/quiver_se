package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMTQWarden(t *testing.T) {
	in_updates := make(
		chan Augmented[
			QuiverUpdate[
				uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex],
			],
			struct{},
		],
	)
	out_walks := make(
		chan Augmented[
			QuiverWalk[uint32_H, PHashMap[Literal[uint32_H], struct{}]],
			struct{},
		],
	)
	var dmtq Quiver[uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex]]
	top_node_dmt := NewDMT[uint32_H, QuiverIndex]()
	top_node := dmtq.InsertNode(uint32_H{0}, &top_node_dmt)
	fail_node_dmt := NewDMT[uint32_H, QuiverIndex]()
	fail_node := dmtq.InsertNode(uint32_H{251}, &fail_node_dmt)
	warden_config := DMTQWardenConfig[uint32_H, uint32_H, struct{}]{
		in_updates: in_updates,
		out_walks:  out_walks,
		walk_src:   top_node,
		walk_dst:   fail_node,
		dmtq:       dmtq,
	}
	warden_config.Start()
	update_node_dmt := NewDMT[uint32_H, QuiverIndex]()
	update1 := QuiverUpdate[
		uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex],
	]{
		top_node,
		NewQuiverIntendedNode[
			uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex], any,
		](
			uint32_H{1},
			&update_node_dmt,
		),
		StdlibMapToPHashMap(
			map[Literal[uint32_H]]struct{}{
				{uint32_H{46}, false}: {},
			},
		),
	}
	update2 := QuiverUpdate[
		uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex],
	]{
		top_node,
		dmtq.ParameterizeIndex(fail_node),
		StdlibMapToPHashMap(
			map[Literal[uint32_H]]struct{}{
				{uint32_H{47}, false}: {},
			},
		),
	}
	in_updates <- NewAugmentedSimple(update1)
	in_updates <- NewAugmentedSimple(update2)
	close(in_updates)
	walks := make([][]PHashMap[Literal[uint32_H], struct{}], 0)
	for walk_recv := range out_walks {
		walk_chunked := walk_recv.value
		new_walk := make([]PHashMap[Literal[uint32_H], struct{}], 0)
		for _, chunk := range walk_chunked.edges_chunked {
			new_walk = append(new_walk, *chunk...)
		}
		walks = append(walks, new_walk)
	}
	assert.Equal(t, 1, len(walks))
	assert.Equal(t, 1, len(walks[0]))
	assert.True(
		t,
		walks[0][0].HasKey(Literal[uint32_H]{uint32_H{47}, false}),
	)
	assert.Equal(
		t,
		0,
		len(dmtq.AllInneighbors(top_node)),
	)
	top_outneighbors := dmtq.AllOutneighbors(top_node)
	assert.Equal(
		t,
		2,
		len(top_outneighbors),
	)
}
