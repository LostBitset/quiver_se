package main

import (
	"hash/fnv"

	qse "github.com/LostBitset/quiver_se/lib"
)

const PATHCOND_FLATTENING_BEGIN_SEGMENT = "@__RAW__;;@RICHPC:bgn-segment "

func FlattenSpc(spc []PathCondSegment) (pc FlatPc) {
	items := make([]qse.IdLiteral[string], 0)
	for _, segm := range spc {
		var trxn_pc_item PathCondItem
		if segm.Event == nil {
			trxn_pc_item = PathCondItem{
				PATHCOND_FLATTENING_BEGIN_SEGMENT + "__top__",
				true,
			}
		} else {
			trxn_pc_item = PathCondItem{
				PATHCOND_FLATTENING_BEGIN_SEGMENT + *segm.Event,
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

func (item PathCondItem) AsIdLiteral() (idlit qse.IdLiteral[string]) {
	hasher := fnv.New32a()
	hasher.Write([]byte(item.Constraint))
	hash32_constraint := hasher.Sum32()
	var idlit_raw qse.Literal[qse.WithId_H[string]]
	if item.Followed {
		idlit_raw = qse.BufferingLiteral(
			qse.WithId_H[string]{
				Value: item.Constraint,
				Id:    qse.NumericId(hash32_constraint),
			},
		)
	} else {
		idlit_raw = qse.InvertingLiteral(
			qse.WithId_H[string]{
				Value: item.Constraint,
				Id:    qse.NumericId(hash32_constraint),
			},
		)
	}
	idlit = qse.IdLiteral[string](idlit_raw)
	return
}
