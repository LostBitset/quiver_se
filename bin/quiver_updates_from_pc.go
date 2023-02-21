package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	qse "LostBitset/quiver_se/lib"
)

func QuiverUpdatesFromPathCondition(
	pc eidin.PathCondition,
	known_callbacks map[int]qse.QuiverIndex,
	top_node qse.QuiverIndex,
	fail_node qse.QuiverIndex,
) (
	quiver_updates []qse.Augmented[
		qse.QuiverUpdate[
			int,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	],
) {
	// TODO
}
