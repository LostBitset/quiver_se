package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
	"strings"

	qse "LostBitset/quiver_se/lib"

	"google.golang.org/protobuf/proto"
)

func PathConditionToAnalyzeMessages(msg eidin.PathCondition) (msgs [][]byte) {
	pc, free_funs := PathConditionMessageToConjunction(msg)
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{Idsrc: idsrc}
pcAlternativesLoop:
	for i := range pc {
		orig := pc[i]
		if strings.HasPrefix(orig.Value.Value, "@__RAW__") {
			continue pcAlternativesLoop
		}
		pc[i] = qse.IdLiteral[string]{
			Value: orig.Value,
			Eq:    !orig.Eq,
		}
		fmt.Println("[simple_dse] Querying solver...")
		sctx := sys.CheckSat(pc, free_funs)
		if *sctx.IsSat() {
			fmt.Println("[simple_dse] Obtained model. ")
			model := *sctx.GetModel()
			msgs = append(msgs, MakeAnalyzeMessage(model))
			fmt.Println("[simple_dse] Created new Analyze message (canidate).")
		}
		pc[i] = orig
	}
	return
}

func MakeAnalyzeMessage(model string) (msg_raw []byte) {
	msg := &eidin.Analyze{
		ForbidCaching: false,
		Model:         &model,
	}
	out, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	msg_raw = out
	return
}

func PathConditionMessageToConjunction(msg eidin.PathCondition) (
	conjunction []qse.IdLiteral[string],
	free_funs []qse.SMTFreeFun[string, string],
) {
	free_funs = make([]qse.SMTFreeFun[string, string], 0)
	for _, free_fun_ref := range msg.GetFreeFuns() {
		free_fun := *free_fun_ref
		free_funs = append(
			free_funs,
			qse.SMTFreeFun[string, string]{
				Name: free_fun.GetName(),
				Args: free_fun.GetArgSorts(),
				Ret:  free_fun.GetRetSort(),
			},
		)
	}
	conjunction = make([]qse.IdLiteral[string], 0)
	for _, segment_ref := range msg.GetSegmentedPc() {
		segment := *segment_ref
		for _, constraint_ref := range segment.GetPartialPc() {
			constraint := *constraint_ref
			hasher := fnv.New32a()
			hasher.Write([]byte(constraint.GetConstraint()))
			hash := hasher.Sum32()
			conjunction = append(
				conjunction,
				qse.IdLiteral[string](
					qse.Literal[qse.WithId_H[string]]{
						Value: qse.WithId_H[string]{
							Value: constraint.GetConstraint(),
							Id:    hash,
						},
						Eq: constraint.GetAssertionValue(),
					},
				),
			)
		}
	}
	return
}
