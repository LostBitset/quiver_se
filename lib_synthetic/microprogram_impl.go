package libsynthetic

import (
	"fmt"
	"math/rand"
)

func (gen *MicroprogramGenerator) GetNextStateId() (state MicroprogramState) {
	state = gen.next_state_id
	gen.next_state_id += 1
	return
}

func (gen *MicroprogramGenerator) AllocateStateIds(n int) (start_of_allocation MicroprogramState) {
	gen.next_state_id = gen.next_state_id.ShiftBy(n)
	return
}

func (state MicroprogramState) ShiftBy(n int) (shifted_state MicroprogramState) {
	shifted_state = state + MicroprogramState(n)
	return
}

func (gen *MicroprogramGenerator) RandomMicroprogram() (uprgm Microprogram) {
	// Generate a random quiver to start
	base_quiver := ErdosRenyiQuiverGivenEdges(
		gen.P_n_states,
		gen.P_p_transition,
		gen.P_n_merged_graphs,
	)
	fmt.Println("adj list length:")
	fmt.Println(len(base_quiver.adj_list))
	//fmt.Println("ADJLIST ABOVE!!!!")
	n_nodes := gen.P_n_states + 2
	node_allocation := gen.AllocateStateIds(n_nodes)
	// Add a failure node and connections to it with probability p_fallible
	failure_node := gen.P_n_states
	failure_state := node_allocation.ShiftBy(failure_node)
	for i := 0; i < gen.P_n_states; i++ {
		if rand.Float64() < gen.P_p_fallible {
			base_quiver.InsertEdge(i, failure_node)
		}
	}
	// Add a top state connected to n_entry_samples random nodes
	top_node := gen.P_n_states + 1
	top_state := node_allocation.ShiftBy(top_node)
	for i := 0; i < gen.P_n_entry_samples; i++ {
		base_quiver.InsertEdge(top_node, rand.Intn(gen.P_n_states))
	}
	for _, edge := range base_quiver.adj_list {
		fmt.Printf("& %v -> %v\n", edge.src, edge.dst)
	} // DEBUG
	// Replace edges with random trees of random constraints
	adj_list_map := base_quiver.ExtractAdjListAsMap(n_nodes)
	uprgm_transitions := make(map[MicroprogramState][]MicroprogramTransition)
buildUpUprgmTransitionsLoop:
	for src, dst_list := range adj_list_map {
		if len(dst_list) == 0 {
			continue buildUpUprgmTransitionsLoop
		}
		var new_transitions []MicroprogramTransition
		src_state := node_allocation.ShiftBy(src)
		if len(dst_list) == 1 {
			new_transitions = []MicroprogramTransition{
				{
					StateDst: node_allocation.ShiftBy(dst_list[0]),
					Constraints: []string{
						gen.P_constraintgen.Generate(BoolSort),
					},
				},
			}
		} else {
			n_branches := len(dst_list)
			tree := PruferEvenFinalRandomTree(gen.P_n_tree_nonleaf, n_branches)
			tree.CoerceToMaxDegree(2)
			dst_states := make([]MicroprogramState, n_branches)
			for i, dst := range dst_list {
				dst_states[i] = node_allocation.ShiftBy(dst)
			}
			new_transitions = tree.AsMicroprogramTransitions(dst_states, gen.P_constraintgen)
		}
		uprgm_transitions[src_state] = append(uprgm_transitions[src_state], new_transitions...)
	}
	uprgm = Microprogram{
		StateTop:      top_state,
		StateFail:     failure_state,
		Transitions:   uprgm_transitions,
		smt_free_funs: gen.P_constraintgen.SMTFreeFuns(),
	}
	return
}
