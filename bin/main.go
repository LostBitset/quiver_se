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

// Everything enclosed between "// bgn EXAMPLE SPECIFIC" and
// "// end EXAMPLE SPECIFIC" is, as you might have guessed,
// specific to the given example. This is all information that
// can easily be generated automatically, I just don't have the
// time to implement that right now. This is a research project,
// and I just want to see if the idea holds any merit right now.

func main() {
	SanityCheck()
	cwd, err_cwd := os.Getwd()
	if err_cwd != nil {
		panic(err_cwd)
	}
	// bgn EXAMPLE SPECIFIC
	target := cwd + "/example._fninf.js"
	// end EXAMPLE SPECIFIC
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
	go func() {
		for model := range out_models {
			SendAnalyzeRequest(msg_prefix, model)
		}
	}()
	known_callbacks := make(map[int]qse.QuiverIndex)
	// bgn EXAMPLE SPECIFIC
	known_callback_locations := []int{170, 355, 489}
	// end EXAMPLE SPECIFIC
	for _, loc := range known_callback_locations {
		backing_dmt := qse.NewDMT[qse.WithId_H[string], qse.QuiverIndex]()
		known_callbacks[loc] = dmtq.InsertNode(loc, &backing_dmt)
	}
	pc_chan := make(chan eidin.PathCondition)
	go func() {
		for pc := range pc_chan {
			quiver_updates := QuiverUpdatesFromPathCondition(
				pc,
				known_callbacks,
				top_node,
				fail_node,
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
			BytesStart: 170,
			BytesEnd:   335,
		},
		NextCallbackId: &eidin.CallbackId{
			BytesStart: 1,
			BytesEnd:   1,
		},
		PartialPc: []*eidin.SMTConstraint{
			&eidin.SMTConstraint{
				Constraint:     "(= (*/read-var/* **jsvar_z) Y)",
				AssertionValue: &yes,
			},
			&eidin.SMTConstraint{
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
}
