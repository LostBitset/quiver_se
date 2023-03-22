package main

import (
	qse "LostBitset/quiver_se/lib"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ParseMicroprogramState(str string) (state MicroprogramState) {
	integer, err := strconv.Atoi(str)
	if err != nil {
		log.Warnf("Failed to parse microprogram state: \"%s\"\n", str)
		panic(err)
	}
	state = MicroprogramState(integer)
	return
}

func (uprgm Microprogram) SiMReQProcessPCs(
	in_pcs chan []string,
	bug_signal chan struct{},
) {
	// Setup everything necessary
	in_updates := make(chan qse.Augmented[
		qse.QuiverUpdate[
			MicroprogramState,
			qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
			*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
		],
		[]qse.SMTFreeFun[string, string],
	])
	out_models_unfiltered := make(chan string)
	var idsrc qse.IdSource
	sys := qse.SMTLibv2StringSystem{Idsrc: idsrc}
	// Start SiMReQ
	dmtq, top_node, fail_node, _ := qse.StartSiMReQ[
		MicroprogramState, string, string, string, string, qse.SMTLibv2StringSolvedCtx,
	](
		in_updates, out_models_unfiltered, sys, nil,
	)
	// Create all necessary nodes (for each state/callback)
	callback_nodes := make(map[MicroprogramState]qse.QuiverIndex)
addNodesForMicroprogramStatesLoop:
	for state := range uprgm.transitions {
		if state == uprgm.top_state {
			// Top state gets special treatment
			continue addNodesForMicroprogramStatesLoop
		}
		if state == uprgm.fail_state {
			// Fail state gets special treatment
			continue addNodesForMicroprogramStatesLoop
		}
		update_dmt := qse.NewDMT[qse.WithId_H[string], qse.QuiverIndex]()
		added_node_index := dmtq.InsertNode(state, &update_dmt)
		callback_nodes[state] = added_node_index
	}
	// Overwrite those for top and failure states since they have special indices on the quiver
	callback_nodes[uprgm.top_state] = top_node
	callback_nodes[uprgm.fail_state] = fail_node
	go func() {
		defer close(in_updates)
		defer close(bug_signal)
		for model_unfiltered := range out_models_unfiltered {
			canidate_model := FilterModelFromZ3(model_unfiltered)
			fails, pc := uprgm.ExecuteGetPathCondition(canidate_model)
			if fails {
				bug_signal <- struct{}{}
			} else {
				in_pcs <- pc
			}
		}
	}()
	for pc := range in_pcs {
		// Group the segmented path condition by segments (which represent transitions)
		grouped_by_transition := make(map[SimpleMicroprogramTransitionDesc][]string)
		current_transition_constraint := make([]string, 0)
	groupPcSegmentsLoop:
		for _, item := range pc {
			if strings.HasPrefix(item, "@__RAW__;;@RICHPC:") {
				if strings.HasPrefix(item, "@__RAW__;;@RICHPC:was-segment ") {
					fields := strings.Fields(item)
					src_state := ParseMicroprogramState(fields[1])
					dst_state := ParseMicroprogramState(fields[2])
					edge_desc := SimpleMicroprogramTransitionDesc{src_state, dst_state}
					new_constraint := make([]string, len(current_transition_constraint))
					copy(new_constraint, current_transition_constraint)
					grouped_by_transition[edge_desc] = new_constraint
					continue groupPcSegmentsLoop
				}
				log.Warn("Unknown rich path condition marker.")
				continue groupPcSegmentsLoop
			}
			current_transition_constraint = append(current_transition_constraint, item)
		}
		// Send the updates to SiMReQ
		for transition, constraints := range grouped_by_transition {
			constraints_in_qse_form := make([]qse.Literal[qse.WithId_H[string]], len(constraints))
			for i, constraint := range constraints {
				id_literal_constraint := MicroprogramConstraintToIdLiteral(constraint, &idsrc)
				constraints_in_qse_form[i] = qse.Literal[qse.WithId_H[string]](id_literal_constraint)
			}
			in_updates <- qse.Augmented[
				qse.QuiverUpdate[
					MicroprogramState,
					qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
					*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
				],
				[]qse.SMTFreeFun[string, string],
			]{
				Value: qse.QuiverUpdate[
					MicroprogramState,
					qse.PHashMap[qse.Literal[qse.WithId_H[string]], struct{}],
					*qse.DMT[qse.WithId_H[string], qse.QuiverIndex],
				]{
					Src: callback_nodes[transition.src],
					Dst: dmtq.ParameterizeIndex(callback_nodes[transition.dst]),
					Edge: pto(SliceToPHashMapSet(
						constraints_in_qse_form,
					)),
				},
				Augment: uprgm.smt_free_funs,
			}
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
