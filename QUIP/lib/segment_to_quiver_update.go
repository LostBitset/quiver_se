package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	q "LostBitset/quiver_se/lib"
)

func SegmentToQuiverUpdate(
	segment eidin.PathConditionSegment,
	free_funs []q.SMTFreeFun[string, string],
	top_node q.QuiverIndex,
	fail_node q.QuiverIndex,
) (
	update q.Augmented[
		q.QuiverUpdate[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
		],
		[]q.SMTFreeFun[string, string],
	],
) {
	// TODO
	return
}
