package main

/*
import qse "LostBitset/quiver_se/lib"

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

func (uprgm Microprogram) ModelSatisfiesConstraints(model string, constraints []string) (does bool) {
	complete_query := make([]string, len(constraints)+1)
	for i := range complete_query {
		if i == 0 {
			complete_query[i] = "@__RAW__" + model
		} else {
			complete_query[i] = constraints[i-1]
		}
	}
	var idsrc qse.IdSource
	complete_query_with_ids := make([]qse.IdLiteral[string], len(complete_query))
	for i, part := range complete_query {
		complete_query_with_ids[i] = Wi
	}
	sys := qse.SMTLibv2StringSystem{idsrc}
	sys.CheckSat(complete_query)
}
*/
