package main

func (uprgm Microprogram) ExecuteGetPathCondition(model string) (pc []string) {
	pc = uprgm.ExecuteGetPathConditionFrom(model, uprgm.top_state)
	return
}

func (uprgm Microprogram) ExecuteGetPathConditionFrom(
	model string, state MicroprogramState,
) (
	pc []string,
) {
	transitions := uprgm.transitions[state]
	pc = make([]string, 0)
selectTransitionLoop:
	for _, transition := range transitions {
		if uprgm.ModelSatisfiesConstraints(model, transition.constraints) {
			pc = append(pc, transition.constraints...)
			pc = append(pc, uprgm.ExecuteGetPathConditionFrom(
				model, transition.dst_state,
			)...)
			break selectTransitionLoop
		}
	}
	return
}
