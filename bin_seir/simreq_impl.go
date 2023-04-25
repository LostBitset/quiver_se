package main

import (
	"fmt"
	"strings"

	qse "github.com/LostBitset/quiver_se/lib"
)

func (sp SeirPrgm) SiMReQProcessPCs(in_pcs chan FlatPc) {
	in_updates := make(chan qse.Augmented[
		qse.QuiverUpdate[
			string,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	])
	out_models := make(chan string)
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{idsrc}
	// Start SiMReQ
	dmtq, top_node, fail_node, _ := qse.StartSiMReQ[
		string, string, string, string, string, qse.SMTLib2VAStringSolvedCtx,
	](
		in_updates, out_models, sys, nil,
	)
	callback_nodes := make(map[string]qse.QuiverIndex)
	callback_nodes["__top__"] = top_node
	callback_nodes[SEIR_RESERVED_FAILURE_EVENT] = fail_node
	// Start goroutine to handle canidate models
	go func() {
		for model := range out_models {
			fmt.Println("::: GOT SIMREQ MODEL :::")
			fmt.Println(model)
			fmt.Println("::: END SIMREQ MODEL :::")
			_, fails := sp.PerformQuery(
				ParseZ3ModelString(model),
			)
			if fails {
				fmt.Println("FOUND A BUG!!!! (from simreq model)")
			}
		}
	}()
	seen_events := make(map[string]struct{})
	seen_events["__top__"] = struct{}{}
	seen_events[SEIR_RESERVED_FAILURE_EVENT] = struct{}{}
	// Process path conditions received
	for pc := range in_pcs {
		// Group the segmented path condition into segments
		grouped_by_trxn := make([]SiMReQConstrainedTrxn, 0)
		prev_event := "__top__"
		current_trxn_constraints := make([]qse.IdLiteral[string], 0)
	groupPcSegmentsLoop:
		for _, item := range pc.items {
			constraint := item.Value.Value
			if strings.HasPrefix(constraint, "@__RAW__;;@RICHPC:") {
				if strings.HasPrefix(constraint, PATHCOND_FLATTENING_BEGIN_SEGMENT) {
					event, _ := strings.CutPrefix(constraint, PATHCOND_FLATTENING_BEGIN_SEGMENT)
					if event == "__top__" {
						continue groupPcSegmentsLoop
					}
					new_constraints := make([]qse.IdLiteral[string], len(current_trxn_constraints))
					copy(new_constraints, current_trxn_constraints)
					grouped_by_trxn = append(
						grouped_by_trxn,
						SiMReQConstrainedTrxn{
							src_event:   prev_event,
							dst_event:   event,
							constraints: new_constraints,
						},
					)
				} else {
					panic("Unknown rich path condition marker.")
				}
			} else {
				current_trxn_constraints = append(current_trxn_constraints, item)
			}
		}
		// Send the updates to SiMReQ
		for _, ctrxn := range grouped_by_trxn {
			constraints := make([]qse.Literal[qse.WithId_H[string]], len(ctrxn.constraints))
			for i, constraint := range ctrxn.constraints {
				constraints[i] = qse.Literal[qse.WithId_H[string]]{
					Value: constraint.Value,
					Eq:    constraint.Eq,
				}
			}
			update := qse.Augmented[
				qse.QuiverUpdate[
					string,
					qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
					*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
				],
				[]qse.SMTFreeFun[string, string],
			]{
				Value: qse.QuiverUpdate[
					string,
					qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
					*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
				]{
					Src: callback_nodes[ctrxn.src_event],
					Dst: dmtq.ParameterizeIndex(callback_nodes[ctrxn.dst_event]),
					Edge: Pto(SliceToPHashMapSet(
						constraints,
					)),
				},
				Augment: sp.smt_free_funs,
			}
			in_updates <- update
		}
	}
}

func SliceToSet[T comparable](slice []T) (set map[T]struct{}) {
	set = make(map[T]struct{})
	for _, elem := range slice {
		set[elem] = struct{}{}
	}
	return
}

func SliceToPHashMapSet[T qse.Hashable](slice []T) (set qse.PHashMap[T, struct{}]) {
	set = qse.StdlibMapToPHashMap(SliceToSet(slice))
	return
}
