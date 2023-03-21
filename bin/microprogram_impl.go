package main

func (gen *MicroprogramGenerator) GetNextStateId() (state MicroprogramState) {
	state = gen.next_state_id
	gen.next_state_id += 1
	return
}

func (gen *MicroprogramGenerator) RandomMicroprogram() (uprgm Microprogram) {
	// Generate a random quiver to start
	base_quiver := ErdosRenyiQuiverGivenEdges(
		gen.n_states,
		gen.p_transition,
		gen.avg_n_transitions,
	)
	// Add a failure node and connections to it with probability p_fallible
	// Add a top state connected to n_entry_samples random nodes
	// Replace edges with random trees of random constraints
}
