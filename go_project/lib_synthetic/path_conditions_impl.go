package libsynthetic

import (
	"fmt"
	"strconv"
	"strings"

	qse "github.com/LostBitset/quiver_se/lib"

	log "github.com/sirupsen/logrus"
)

const PC_REC_LIMIT = 30
const PC_CYCLE_LIMIT = 1

func (uprgm Microprogram) ExecuteGetPathCondition(
	model string, no_transition bool,
) (
	fails bool,
	pc []string,
) {
	fmt.Println("[REPORT] [EVAL-INFO] EXECUTION")
	fails, pc = uprgm.ExecuteGetPathConditionFrom(
		model,
		uprgm.StateTop,
		no_transition,
		PC_REC_LIMIT,
		make(map[MicroprogramState]int),
	)
	return
}

func (uprgm Microprogram) ExecuteGetPathConditionFrom(
	model string,
	state MicroprogramState,
	no_transition bool,
	rec_budget int,
	seen map[MicroprogramState]int,
) (
	fails bool,
	pc []string,
) {
	log.Info("[bin:path_conditions] Microprogram entered state: " + strconv.Itoa(int(state)))
	pc = make([]string, 0)
	fails = state == uprgm.StateFail
	if fails {
		fmt.Println("FAIL FAIL FAIL FOUND !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return
	}
	if rec_budget == 0 {
		fmt.Println("END OF PATH CONDITION (rec budget)")
		return
	}
	if v, ok := seen[state]; ok && (v > PC_CYCLE_LIMIT) {
		fmt.Println("END OF PATH CONDITION (cycle budget)")
		return
	}
	transitions := uprgm.Transitions[state]
	not_taken := make([]string, 0)
	for _, transition := range transitions {
		if uprgm.ModelSatisfiesConstraints(model, transition.Constraints) {
			pc = append(pc, not_taken...)
			pc = append(pc, transition.Constraints...)
			if !no_transition {
				pc = append(
					pc,
					"@__RAW__;;@RICHPC:was-segment "+
						strconv.Itoa(int(state))+
						" "+
						strconv.Itoa(int(transition.StateDst)),
				)
			}
			log.Infof(
				"[bin:path_conditions] PC: %#+v\n",
				pc,
			)
			if no_transition {
				// no rec_pc but we still need to check failure ourselves
				fails = transition.StateDst == uprgm.StateFail
				return
			}
			updatedSeen := make(map[MicroprogramState]int)
			for key := range seen {
				updatedSeen[key] = seen[key]
			}
			if _, ok := updatedSeen[state]; ok {
				updatedSeen[state] += 1
			} else {
				updatedSeen[state] = 1
			}
			fmt.Println("STATE = " + fmt.Sprintf("%d", state))
			fmt.Println("STATE_FAIL = " + fmt.Sprintf("%d", uprgm.StateFail))
			rec_fails, rec_pc := uprgm.ExecuteGetPathConditionFrom(
				model,
				transition.StateDst,
				no_transition,
				rec_budget-1,
				updatedSeen,
			)
			pc = append(pc, rec_pc...)
			log.Infof(
				"[bin:path_conditions] Bubbled up PC: %#+v\n",
				pc,
			)
			if rec_fails {
				fails = true
			}
			return
		}
		for _, not_taken_constraint := range transition.Constraints {
			not_taken = append(
				not_taken,
				InvertedConstraintForMicroprogram(not_taken_constraint),
			)
		}
	}
	pc = not_taken
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
		complete_query_with_ids[i] = MicroprogramConstraintToIdLiteral(part, &idsrc)
	}
	sys := qse.SMTLibv2StringSystem{Idsrc: idsrc}
	solver_result := sys.CheckSat(
		complete_query_with_ids,
		[]qse.SMTFreeFun[string, string]{}, // Query contains model, which gives definitions
	)
	does_ptr := solver_result.IsSat()
	does = false
	if does_ptr != nil && *does_ptr {
		does = true
	}
	return
}

func MicroprogramConstraintToIdLiteral(
	constraint_raw string, idsrc *qse.IdSource,
) (
	id_literal qse.IdLiteral[string],
) {
	constraint := constraint_raw
	inverted := false
	for strings.HasPrefix(constraint, "@__INVERTED__") {
		constraint, _ = strings.CutPrefix(constraint, "@__INVERTED__")
		inverted = !inverted
	}
	var id_literal_raw qse.Literal[qse.WithId_H[string]]
	if inverted {
		id_literal_raw = qse.InvertingLiteral(
			qse.WithId_H[string]{
				Value: constraint,
				Id:    idsrc.Gen(),
			},
		)
	} else {
		id_literal_raw = qse.BufferingLiteral(
			qse.WithId_H[string]{
				Value: constraint,
				Id:    idsrc.Gen(),
			},
		)
	}
	id_literal = qse.IdLiteral[string](id_literal_raw)
	return
}

func (uprgm Microprogram) UnitializedAssignment() (model string) {
	sort_values_uninit := map[string]string{
		"Real": "0.0",
		"Bool": "false",
	}
	model = uprgm.UniformAssignmentOfSMTFreeFuns(sort_values_uninit)
	return
}

func (uprgm Microprogram) UniformAssignmentOfSMTFreeFuns(
	sort_values map[string]string,
) (
	model string,
) {
	model = UniformAssignmentOfSMTFreeFuns(uprgm.smt_free_funs, sort_values)
	return
}

func UniformAssignmentOfSMTFreeFuns(
	smt_free_funs []qse.SMTFreeFun[string, string],
	sort_values map[string]string,
) (
	model string,
) {
	var sb strings.Builder
	for _, smt_free_fun := range smt_free_funs {
		if len(smt_free_fun.Args) != 0 {
			panic("Invalid. Cannot generate a uniform assignment of parametric SMT funs.")
		}
		sb.WriteString(
			StringSMTFreeFun{smt_free_fun}.DefinitionString(
				sort_values[smt_free_fun.Ret],
			),
		)
		sb.WriteRune('\n')
	}
	model = sb.String()
	return
}
