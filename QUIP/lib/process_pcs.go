package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	q "LostBitset/quiver_se/lib"
)

// This function does the following:
// - Runs partial DSE on new callbacks
// - Listens for path conditions, updating the quiver accordingly
func ProcessPathConditions(
	out_updates chan q.Augmented[
		q.QuiverUpdate[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
		],
		[]q.SMTFreeFun[string, string],
	],
	top_node q.QuiverIndex,
	fail_node q.QuiverIndex,
	target string,
	msg_prefix string,
) {
	segment_chan := make(chan eidin.PathConditionSegment)
	seen_callbacks := make(map[uint64]struct{})
	go InterpretPathConditionSegments(segment_chan, msg_prefix)
	for segment := range segment_chan {
		cb_this := segment.GetThisCallbackId()
		cb_next := segment.GetNextCallbackId()
		if cb_this.GetBytesStart() != cb_this.GetBytesEnd() {
			if _, ok := seen_callbacks[cb_this.GetBytesStart()]; !ok {
				PerformPartialDse(*cb_this, target, msg_prefix)
			}
		}
		if cb_next.GetBytesStart() != cb_next.GetBytesEnd() {
			if _, ok := seen_callbacks[cb_next.GetBytesStart()]; !ok {
				PerformPartialDse(*cb_next, target, msg_prefix)
			}
		}
		out_updates <- SegmentToQuiverUpdate(segment, top_node, fail_node)
	}
	close(out_updates)
}

func InterpretPathConditionSegments(
	segment_chan chan eidin.PathConditionSegment,
	msg_prefix string,
) {
	// TODO
}
