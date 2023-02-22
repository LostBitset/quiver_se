package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	qse "LostBitset/quiver_se/lib"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func SanityCheck() {
	greeting := qse.Greet("World")
	fmt.Println(greeting)
}

// Everything enclosed between "// bgn EXAMPLE SPECIFIC" and
// "// end EXAMPLE SPECIFIC" is, as you might have guessed,
// specific to the given example. This is all information that
// can easily be generated automatically, I just don't have the
// time to implement that right now. This is a research project,
// and I just want to see if the idea holds any merit right now.

func main() {
	SanityCheck()
	log.SetLevel(log.InfoLevel)
	log.Info("[bin/main.go] Started.")
	cwd, err_cwd := os.Getwd()
	if err_cwd != nil {
		panic(err_cwd)
	}
	// bgn EXAMPLE SPECIFIC
	target := cwd + "/example._fninf.js"
	// end EXAMPLE SPECIFIC
	msg_prefix := GetMessagePrefix(target)
	fmt.Println("@parameter target=\"" + target + "\"")
	fmt.Println("@parameter msg_prefix=\"" + msg_prefix + "\"")
	in_updates := make(chan qse.Augmented[
		qse.QuiverUpdate[
			int,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	])
	out_models := make(chan string)
	known_callbacks := make(map[int]qse.QuiverIndex)
	// bgn EXAMPLE SPECIFIC
	aot_nodes := []int{170, 355, 489}
	// end EXAMPLE SPECIFIC
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{Idsrc: idsrc}
	dmtq, top_node, fail_node, aot_indices := qse.StartSiMReQ[
		int,
		string,
		string,
		string,
		string,
		qse.SMTLib2VAStringSolvedCtx,
	](
		in_updates, out_models, sys, aot_nodes,
	)
	for i, loc := range aot_nodes {
		known_callbacks[loc] = aot_indices[i]
	}
	go func() {
		for model := range out_models {
			SendAnalyzeRequest(msg_prefix, model)
		}
	}()

	pc_chan := make(chan eidin.PathCondition)
	go func() {
		for pc := range pc_chan {
			quiver_updates := QuiverUpdatesFromPathCondition(
				pc,
				known_callbacks,
				top_node,
				fail_node,
				dmtq,
			)
			for _, update := range quiver_updates {
				in_updates <- update
			}
		}
	}()
	go StreamPathConditions(msg_prefix, pc_chan)
	// bgn EXAMPLE SPECIFIC
	yes := true
	fail_segment := eidin.PathConditionSegment{
		ThisCallbackId: &eidin.CallbackId{
			BytesStart: 155,
			BytesEnd:   327,
		},
		NextCallbackId: &eidin.CallbackId{
			BytesStart: 1,
			BytesEnd:   1,
		},
		PartialPc: []*eidin.SMTConstraint{
			{
				Constraint:     "(= (*/read-var/* **jsvar_z) Y)",
				AssertionValue: &yes,
			},
			{
				Constraint:     "(*/read-var/* **jsvar_a)",
				AssertionValue: &yes,
			},
		},
	}
	pc_chan <- eidin.PathCondition{
		FreeFuns: []*eidin.SMTFreeFun{},
		SegmentedPc: []*eidin.PathConditionSegment{
			&fail_segment,
		},
	}
	// end EXAMPLE SPECIFIC
	for {
	}
}
