package lib

import (
	q "LostBitset/quiver_se/lib"
)

// QUIP does the following:
// - Runs simple DSE at low frequency
// - Runs partial DSE on new callbacks
// - Listens for path conditions, updating the quiver accordingly
// - Sends Analyze requests in response to SiMReQ

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
	target string,
	msg_prefix string,
) {
	go RunAnalyzer(target, msg_prefix)
	go RunSimpleDSELowFrequency(msg_prefix)
	go ProcessPathConditions(out_updates, top_node, fail_node, target, msg_prefix)
	out_models := make(chan string)
	var idsrc q.IdSource
	sys := q.SMTLib2VAStringSystem{Idsrc: idsrc}
	q.StartSiMReQ[int, string, string, string, string, q.SMTLib2VAStringSolvedCtx](
		out_updates,
		out_models,
		sys,
	)
	go func() {
		for model := range out_models {
			SendAnalyzeMessage(model, msg_prefix)
		}
	}()
}
