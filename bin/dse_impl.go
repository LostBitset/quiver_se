package main

import "strings"

func (uprgm Microprogram) RunDSE() (n_bugs int) {
	n_bugs = 0
	model := uprgm.UnitializedAssignment()
	imm_failure, imm_pc := uprgm.ExecuteGetPathCondition(model)
	if imm_failure {
		n_bugs = -1 // Don't use results when the program failed immediately
		return
	}
	// The two main variables for the concolic execution algorithm:
	alt_stack := make([]uint, 0)
	desired_path := make([]string, 0)
	// Set them up using the first path condition
	stack_setup_index := 0
setupAltStackAndPathLoop:
	for _, constraint := range imm_pc {
		if strings.HasPrefix(constraint, "@__RAW__;;@RICHPC:") {
			continue setupAltStackAndPathLoop
		}
		alt_stack = append(alt_stack, uint(stack_setup_index))
		desired_path = append(desired_path, constraint)
		stack_setup_index++
	}
	// Main loop
	for len(alt_stack) > 0 {
		// TODO all of this
	}
	return
}
