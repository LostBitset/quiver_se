package lib

import (
	q "LostBitset/quiver_se/lib"
)

func StartQUIP(
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
) {
}
