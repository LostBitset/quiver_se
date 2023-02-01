package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMTQWarden(t *testing.T) {
	in_updates := make(
		chan QuiverUpdate[
			uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex],
		],
	)
	out_walks := make(
		chan QuiverWalk[uint32_H, PHashMap[Literal[uint32_H], struct{}]],
	)
	var dmtq Quiver[uint32_H, PHashMap[Literal[uint32_H], struct{}], *DMT[uint32_H, QuiverIndex]]
	top_node_dmt := NewDMT[uint32_H, QuiverIndex]()
	top_node := dmtq.InsertNode(uint32_H{0}, &top_node_dmt)
	warden_config := DMTQWardenConfig[uint32_H, uint32_H]{
		in_updates: in_updates,
		out_walks:  out_walks,
		walk_src:   top_node,
		dmtq:       dmtq,
	}
	warden_config.Start()
	update_node_dmt := NewDMT[uint32_H, QuiverIndex]()
	update := QuiverUpdate[
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
				{uint32_H{47}, false}: {},
			},
		),
	}
	in_updates <- update
	close(in_updates)
	// TODO
	assert.True(t, false)
}
