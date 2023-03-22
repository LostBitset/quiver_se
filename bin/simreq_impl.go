package main

import (
	qse "LostBitset/quiver_se/lib"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"

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

const STAGGER_TIME_INTERVAL = 10 // milliseconds between aot node explorations

const SIMREQ_AOT_DSE_MAX_ITERS = 10

func (uprgm Microprogram) SiMReQProcessPCs(
	in_pcs chan PathConditionResult,
	bug_signal chan uint32,
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
	// Start goroutines for performing DSE AOT on nodes
	stagger_time := 0
runDSEAheadOfTimeForFailureSetLoop:
	for state := range callback_nodes {
		if state == uprgm.fail_state {
			continue runDSEAheadOfTimeForFailureSetLoop
		}
		stagger_time += STAGGER_TIME_INTERVAL
		state := state
		go func() {
			<-time.After(time.Duration(stagger_time) * time.Millisecond)
			bug_signal_black_hole := make(chan uint32)
			go func() {
				for range bug_signal_black_hole {
				}
			}()
			out_pcs_node_failure := make(chan PathConditionResult)
			go func() {
			readNodeFailurePathConditionsLoop:
				for pc := range out_pcs_node_failure {
					if !pc.fails {
						continue readNodeFailurePathConditionsLoop
					}
					constraints := pc.pc
					constraints_in_qse_form := make([]qse.Literal[qse.WithId_H[string]], len(constraints))
					for i, constraint := range constraints {
						id_literal_constraint := MicroprogramConstraintToIdLiteral(constraint, &idsrc)
						constraints_in_qse_form[i] = qse.Literal[qse.WithId_H[string]](id_literal_constraint)
					}
					transition := SimpleMicroprogramTransitionDesc{
						src: state,
						dst: uprgm.fail_state,
					}
					update := qse.Augmented[
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
					in_updates <- update
				}
			}()
			uprgm.RunDSEContinuously(
				bug_signal_black_hole,
				true,
				&out_pcs_node_failure,
				true,
				SIMREQ_AOT_DSE_MAX_ITERS,
			)
		}()
	}
	// Start goroutine to handle canidate models
	go func() {
		defer close(bug_signal)
		for model_unfiltered := range out_models_unfiltered {
			canidate_model := FilterModelFromZ3(model_unfiltered)
			fmt.Println("[bin:simreq] SMR FOUND CANIDATE MODEL:")
			fmt.Println(canidate_model)
			fails, pc := uprgm.ExecuteGetPathCondition(canidate_model, false)
			if fails {
				hasher := fnv.New32a()
				hasher.Write([]byte(canidate_model))
				bug_signal <- hasher.Sum32()
			} else {
				in_pcs <- PathConditionResult{pc, fails}
			}
		}
	}()
	for pc := range in_pcs {
		log.Info("[bin:simreq] Received path condition in SiMReQProcessPCs.")
		// Group the segmented path condition by segments (which represent transitions)
		grouped_by_transition := make(map[SimpleMicroprogramTransitionDesc][]string)
		current_transition_constraint := make([]string, 0)
	groupPcSegmentsLoop:
		for _, item := range pc.pc {
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
			update := qse.Augmented[
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
			/*fmt.Printf("src: %#+v\n", update.Value.Src)
			fmt.Printf("dst: %#+v\n", update.Value.Dst)
			fmt.Printf("edge: %#+v\n", update.Value.Edge.ToStdlibMap())*/
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

func (uprgm Microprogram) RunSiMReQ(bug_signal chan struct{}) {
	bug_signal_values := make(chan uint32)
	go func() {
		seen_model_hashes := make(map[uint32]struct{})
		for model_hash := range bug_signal_values {
			if _, ok := seen_model_hashes[model_hash]; !ok {
				seen_model_hashes[model_hash] = struct{}{}
				bug_signal <- struct{}{}
			}
		}
	}()
	in_pcs := make(chan PathConditionResult)
	go uprgm.RunDSEContinuously(bug_signal_values, true, &in_pcs, false, -1)
	uprgm.SiMReQProcessPCs(in_pcs, bug_signal_values)
}
