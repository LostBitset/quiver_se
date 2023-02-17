package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"hash/fnv"

	qse "LostBitset/quiver_se/lib"
)

func PathConditionToAnalyzeMessages(msg eidin.PathCondition) (msgs [][]byte) {
	pc := make([]qse.IdLiteral[string], 0)
	free_funs := make([]qse.SMTFreeFun[string, string], 0)
	for _, segment_ref := range msg.GetSegmentedPc() {
		segment := *segment_ref
		for _, constraint_ref := range segment.GetPartialPc() {
			constraint := *constraint_ref
			hasher := fnv.New32a()
			hasher.Write([]byte(constraint.GetConstraint()))
			hash := hasher.Sum32()
			pc = append(
				pc,
				qse.IdLiteral[string](
					qse.Literal[qse.WithId_H[string]]{
						qse.WithId_H[string]{
							constraint.GetConstraint(),
							hash,
						},
						constraint.GetAssertionValue(),
					},
				),
			)
		}
	}
	return
}
