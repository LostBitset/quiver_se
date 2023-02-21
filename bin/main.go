package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	qse "LostBitset/quiver_se/lib"
	"fmt"
	"os"
)

func SanityCheck() {
	greeting := qse.Greet("World")
	fmt.Println(greeting)
}

func main() {
	SanityCheck()
	cwd, err_cwd := os.Getwd()
	if err_cwd != nil {
		panic(err_cwd)
	}
	target := cwd + "/example._fninf.js"
	msg_prefix := GetMessagePrefix(target)
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
	go func() {}()
}
