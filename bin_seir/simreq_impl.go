package main

import (
	"fmt"

	qse "github.com/LostBitset/quiver_se/lib"
)

func (sp SeirPrgm) SiMReQProcessPCs(in_pcs chan FlatPc) {
	in_updates := make(chan qse.Augmented[
		qse.QuiverUpdate[
			string,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	])
	out_models := make(chan string)
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{idsrc}
	// Start SiMReQ
	dmtq, top_node, fail_node, _ := qse.StartSiMReQ[
		string, string, string, string, string, qse.SMTLib2VAStringSolvedCtx,
	](
		in_updates, out_models, sys, nil,
	)
	callback_nodes := make(map[string]qse.QuiverIndex)
	callback_nodes["__top__"] = top_node
	callback_nodes[SEIR_RESERVED_FAILURE_EVENT] = fail_node
	// Start goroutine to handle canidate models
	go func() {
		for model := range out_models {
			fmt.Println("::: GOT SIMREQ MODEL :::")
			fmt.Println(model)
			fmt.Println("::: END SIMREQ MODEL :::")
			_, fails := sp.PerformQuery(
				ParseZ3ModelString(model),
			)
			if fails {
				fmt.Println("FOUND A BUG!!!! (from simreq model)")
			}
		}
	}()
	panic("TODO TODO TODO")
}
