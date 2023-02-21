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
	// TODO
}
