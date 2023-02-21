package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	qse "LostBitset/quiver_se/lib"
)

func QuiverUpdatesFromPathCondition(
	pc eidin.PathCondition,
	known_callbacks map[int]qse.QuiverIndex,
	top_node qse.QuiverIndex,
	fail_node qse.QuiverIndex,
	dmtq qse.Quiver[
		int,
		qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
		*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
	],
) (
	quiver_updates []qse.Augmented[
		qse.QuiverUpdate[
			int,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	],
) {
	quiver_updates = make(
		[]qse.Augmented[
			qse.QuiverUpdate[
				int,
				qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
				*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
			],
			[]qse.SMTFreeFun[string, string],
		],
		0,
	)
	eidin_free_funs := pc.GetFreeFuns()
	free_funs := make([]qse.SMTFreeFun[string, string], 0)
	for _, eidin_free_fun := range eidin_free_funs {
		eidin_free_fun := *eidin_free_fun
		free_funs = append(
			free_funs,
			qse.SMTFreeFun[string, string]{
				Name: eidin_free_fun.GetName(),
				Args: eidin_free_fun.GetArgSorts(),
				Ret:  eidin_free_fun.GetRetSort(),
			},
		)
	}
	sound_prefix := make([]eidin.SMTConstraint, 0)
	for _, segment := range pc.GetSegmentedPc() {
		segment := *segment
		for _, ppc := range segment.GetPartialPc() {
			ppc := *ppc
			sound_prefix = append(sound_prefix, ppc)
		}
		src_index := CallbackIdToQuiverIndex(
			segment.GetThisCallbackId(),
			known_callbacks,
			top_node,
			fail_node,
		)
		dst_index := CallbackIdToQuiverIndex(
			segment.GetNextCallbackId(),
			known_callbacks,
			top_node,
			fail_node,
		)
		src := dmtq.ParameterizeIndex(src_index)
		dst := dmtq.ParameterizeIndex(dst_index)
	}
}
