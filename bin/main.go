package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	qse "LostBitset/quiver_se/lib"
	"fmt"
)

func SanityCheck() {
	greeting := qse.Greet("World")
	fmt.Println(greeting)
}

func main() {
	SanityCheck()
	in_updates := make(chan qse.Augmented[
		qse.QuiverUpdate[
			int,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	])
	out_models := make(chan string)
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{idsrc}
	dmtq, top_node, fail_node := qse.StartSiMReQ[
		int,
		string,
		string,
		string,
		string,
		qse.SMTLib2VAStringSolvedCtx,
	](
		in_updates, out_models, sys,
	)
	pc_chan := make(chan eidin.PathCondition)
}
