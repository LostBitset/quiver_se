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
	quiver_nodes map[int]q.QuiverIndex,
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

func CallbackIdToQuiverIndex(
	cb eidin.CallbackId,
	top_node q.QuiverIndex,
	fail_node q.QuiverIndex,
	quiver_nodes map[int]q.QuiverIndex,
) (
	qindex q.QuiverIndex,
) {
	if cb.GetBytesStart() == cb.GetBytesEnd() {
		switch cb.GetBytesStart() {
		case 0:
			qindex = top_node
			return
		case 1:
			qindex = fail_node
			return
		default:
			panic("Unknown special-case EIDIN CallbackId.")
		}
	}
	qindex = quiver_nodes[int(cb.GetBytesStart())]
	return
}
