package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"hash/fnv"

	qse "LostBitset/quiver_se/lib"
)

func PathConditionToAnalyzeMessages(msg eidin.PathCondition) (msgs [][]byte) {
	pc, free_funs := PathConditionMessageToConjunction(msg)
	var idsrc qse.IdSource
	sys := qse.SMTLibv2StringSystem{Idsrc: idsrc}
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
