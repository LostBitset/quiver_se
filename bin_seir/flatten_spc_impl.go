package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

func FlattenSpc(spc []PathCondSegment) (pc FlatPc) {
	items := make([]qse.IdLiteral[string], 0)
	for _, segm := range spc {
		var trxn_pc_item PathCondItem
		if segm.Event == nil {
			trxn_pc_item = PathCondItem{
				"@__RAW__;;@RICHPC:was-segment __top__",
				true,
			}
		} else {
			trxn_pc_item = PathCondItem{
				"@__RAW__;;@RICHPC:was-segment " + *segm.Event,
				true,
			}
		}
		items = append(items, trxn_pc_item.AsIdLiteral())
		for _, pc_item := range segm.Segment {
			items = append(items, pc_item.AsIdLiteral())
		}
	}
	pc = FlatPc{
		items: items,
	}
	return
}
