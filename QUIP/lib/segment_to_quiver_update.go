package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	q "LostBitset/quiver_se/lib"
	"fmt"
	"hash/fnv"
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
	src, src_ok := CallbackIdToQuiverIndex(
		*segment.GetThisCallbackId(),
		top_node,
		fail_node,
		quiver_nodes,
	)
	dst_index, dst_index_ok := CallbackIdToQuiverIndex(
		*segment.GetNextCallbackId(),
		top_node,
		fail_node,
		quiver_nodes,
	)
	if !src_ok {
		fmt.Println("[QUIP:segment_to_quiver_update.go] Unable to find source in quiver.")
	}
	var dst q.QuiverUpdateDst[
		int,
		q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
		*q.DMT[q.WithId_H[string], q.QuiverIndex],
	]
	if dst_index_ok {
		dst = q.TrustingParameterizeQuiverIndex[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
		](dst_index)
	} else {
		backing_edge_container := q.NewDMT[q.WithId_H[string], q.QuiverIndex]()
		dst = q.NewQuiverIntendedNode[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
			any,
		](
			int(
				(*segment.GetNextCallbackId()).GetBytesStart(),
			),
			&backing_edge_container,
		)
	}
	quiver_edge_stdlib_map := make(map[q.Literal[q.WithId_H[string]]]struct{})
	for _, item := range segment.GetPartialPc() {
		item := item
		constraint := *item
		hasher := fnv.New32a()
		hasher.Write(
			[]byte(constraint.GetConstraint()),
		)
		id := q.NumericId(hasher.Sum32())
		key := q.Literal[q.WithId_H[string]]{
			Value: q.WithId_H[string]{
				Value: constraint.GetConstraint(),
				Id:    id,
			},
			Eq: constraint.GetAssertionValue(),
		}
		quiver_edge_stdlib_map[key] = struct{}{}
	}
	quiver_edge := q.StdlibMapToPHashMap(quiver_edge_stdlib_map)
	quiver_update := q.QuiverUpdate[
		int,
		q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
		*q.DMT[q.WithId_H[string], q.QuiverIndex],
	]{
		Src:  src,
		Dst:  dst,
		Edge: quiver_edge,
	}
	update = q.Augmented[
		q.QuiverUpdate[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
		],
		[]q.SMTFreeFun[string, string],
	]{
		Value:   quiver_update,
		Augment: free_funs,
	}
	return
}

func CallbackIdToQuiverIndex(
	cb eidin.CallbackId,
	top_node q.QuiverIndex,
	fail_node q.QuiverIndex,
	quiver_nodes map[int]q.QuiverIndex,
) (
	qindex q.QuiverIndex,
	ok bool,
) {
	if cb.GetBytesStart() == cb.GetBytesEnd() {
		switch cb.GetBytesStart() {
		case 0:
			qindex, ok = top_node, true
			return
		case 1:
			qindex, ok = fail_node, true
			return
		default:
			panic("Unknown special-case EIDIN CallbackId.")
		}
	}
	qindex, ok = quiver_nodes[int(cb.GetBytesStart())]
	return
}
