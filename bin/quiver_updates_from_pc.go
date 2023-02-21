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
		src := CallbackIdToQuiverIndex(
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
		dst := dmtq.ParameterizeIndex(dst_index)
		quiver_updates = append(
			quiver_updates,
			qse.Augmented[
				qse.QuiverUpdate[
					int,
					qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
					*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
				],
				[]qse.SMTFreeFun[string, string],
			]{
				Value: qse.QuiverUpdate[
					int,
					qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
					*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
				]{
					Src: src,
					Dst: dst,
					Edge: qse.StdlibMapToPHashMap(
						ConstraintsToLiteralMap(sound_prefix),
					),
				},
				Augment: free_funs,
			},
		)
	}
	return
}

func CallbackIdToQuiverIndex(
	cb eidin.CallbackId,
	known_callbacks map[int]qse.QuiverIndex,
	top_node qse.QuiverIndex,
	fail_node qse.QuiverIndex,
) (
	quiver_index qse.QuiverIndex,
) {
	start := cb.GetBytesStart()
	if start == cb.GetBytesEnd() {
		switch start {
		case 0:
			quiver_index = top_node
			return
		case 1:
			quiver_index = fail_node
			return
		default:
			panic("[ERR@bin/quiver_updates_from_pc.go:CallbackIdToQuiverIndex] Bad special start.")
		}
	}
	quiver_index = known_callbacks[int(cb.GetBytesStart())]
	return
}

func ConstraintsToLiteralMap(constraints []eidin.SMTConstraint) (
	m map[qse.Literal[qse.WithId_H[string]]]struct{},
) {
	// TODO
}
