package lib

import (
	q "LostBitset/quiver_se/lib"
	"fmt"
)

// QUIP does the following:
// - Runs simple DSE at low frequency
// - Runs partial DSE on new callbacks
// - Listens for path conditions, updating the quiver accordingly
// - Sends Analyze requests in response to SiMReQ

func StartQUIP(
	target string,
	msg_prefix string,
) {
	out_updates := make(chan q.Augmented[
		q.QuiverUpdate[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
		],
		[]q.SMTFreeFun[string, string],
	])
	go RunAnalyzer(target, msg_prefix)
	go RunSimpleDSELowFrequency(msg_prefix)
	quiver_nodes := make(map[int]q.QuiverIndex)
	out_models := make(chan string)
	var idsrc q.IdSource
	sys := q.SMTLib2VAStringSystem{Idsrc: idsrc}
	_, top_node, fail_node := q.StartSiMReQ[int, string, string, string, string, q.SMTLib2VAStringSolvedCtx](
		out_updates,
		out_models,
		sys,
	)
	go ProcessPathConditions(out_updates, top_node, fail_node, target, "persist_"+msg_prefix, quiver_nodes)
	go func() {
		for model := range out_models {
			SendAnalyzeMessage(model, msg_prefix)
		}
	}()
	fmt.Println("[QUIP:quip_process.go] ENTERING INFINITE LOOP...")
	for {
	}
}
