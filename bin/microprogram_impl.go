package main

import "math/rand"

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
		gen.n_states,
		gen.p_transition,
		gen.avg_n_transitions,
	)
	node_allocation := gen.AllocateStateIds(gen.n_states + 2)
	// Add a failure node and connections to it with probability p_fallible
	failure_node := gen.n_states
	for i := 0; i < gen.n_states; i++ {
		if rand.Float64() < gen.p_fallible {
			base_quiver.InsertEdge(i, failure_node)
		}
	}
	// Add a top state connected to n_entry_samples random nodes
	top_node := gen.n_states + 1
	for i := 0; i < gen.n_entry_samples; i++ {
		base_quiver.InsertEdge(top_node, rand.Intn(gen.n_states))
	}
	// Replace edges with random trees of random constraints
	// TODO
}
