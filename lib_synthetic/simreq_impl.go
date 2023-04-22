package libsynthetic

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"

	qse "github.com/LostBitset/quiver_se/lib"

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

const SIMREQ_JIT_DSE_MAX_ITERS = 10

type SiMReQConstrainedTransition struct {
	transition SimpleMicroprogramTransitionDesc
	constraint []string
}

func (uprgm Microprogram) SiMReQProcessPCs(
	in_pcs chan PathConditionResult,
	bug_signal chan uint32,
	jit_dse bool,
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
	seen_states := make(map[MicroprogramState]struct{})
	seen_states[uprgm.top_state] = struct{}{}
	seen_states[uprgm.fail_state] = struct{}{}
	for pc := range in_pcs {
		log.Info("[bin:simreq] Received path condition in SiMReQProcessPCs.")
		// Group the segmented path condition by segments (which represent transitions)
		grouped_by_transition := make([]SiMReQConstrainedTransition, 0)
		current_transition_constraint := make([]string, 0)
		pc_states := make([]MicroprogramState, 0)
	groupPcSegmentsLoop:
		for _, item := range pc.pc {
			if strings.HasPrefix(item, "@__RAW__;;@RICHPC:") {
				if strings.HasPrefix(item, "@__RAW__;;@RICHPC:was-segment ") {
					fields := strings.Fields(item)
					src_state := ParseMicroprogramState(fields[1])
					dst_state := ParseMicroprogramState(fields[2])
					edge_desc := SimpleMicroprogramTransitionDesc{src_state, dst_state}
					pc_states = append(pc_states, dst_state)
					new_constraint := make([]string, len(current_transition_constraint))
					copy(new_constraint, current_transition_constraint)
					grouped_by_transition = append(
						grouped_by_transition,
						SiMReQConstrainedTransition{edge_desc, new_constraint},
					)
					continue groupPcSegmentsLoop
				}
				log.Warn("Unknown rich path condition marker.")
				continue groupPcSegmentsLoop
			}
			current_transition_constraint = append(current_transition_constraint, item)
		}
		// Loop through new nodes and find failures
		if jit_dse {
			for _, pc_state := range pc_states {
				if _, ok := seen_states[pc_state]; !ok {
					seen_states[pc_state] = struct{}{}
					edge_desc := SimpleMicroprogramTransitionDesc{pc_state, uprgm.fail_state}
					bug_signal_black_hole := make(chan uint32)
					go func() {
						for range bug_signal_black_hole {
						}
					}()
					out_local_pcs := make(chan PathConditionResult)
					go func() {
						defer fmt.Println("[STATUS-JITDSE] end")
						fmt.Println("[STATUS-JITDSE] begin")
						uprgm.RunDSEContinuously(
							bug_signal_black_hole,
							true,
							&out_local_pcs,
							true,
							SIMREQ_JIT_DSE_MAX_ITERS,
							pc_state,
						)
					}()
				updateJITDSEPathConditionsLoop:
					for local_pc := range out_local_pcs {
						if !local_pc.fails {
							continue updateJITDSEPathConditionsLoop
						}
						constraints := local_pc.pc
						grouped_by_transition = append(
							grouped_by_transition,
							SiMReQConstrainedTransition{
								edge_desc,
								constraints,
							},
						)
					}
				}
			}
		}
		// Send the updates to SiMReQ
		for _, constrained_transition := range grouped_by_transition {
			transition, constraints :=
				constrained_transition.transition, constrained_transition.constraint
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
					Edge: Pto(SliceToPHashMapSet(
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

func (uprgm Microprogram) RunSiMReQ(bug_signal chan struct{}, jit_dse bool) {
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
	go uprgm.RunDSEContinuously(bug_signal_values, true, &in_pcs, false, -1, uprgm.top_state)
	uprgm.SiMReQProcessPCs(in_pcs, bug_signal_values, jit_dse)
}
