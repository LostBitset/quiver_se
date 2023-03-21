package main

func (uprgm Microprogram) RunDSE() (n_bugs int) {
	n_bugs = 0
	model := uprgm.UnitializedAssignment()
	imm_failure, imm_pc := uprgm.ExecuteGetPathCondition(model)
	if imm_failure {
		n_bugs = -1 // Don't use results when the program failed immediately
		return
	}
	// The two main variables for the concolic execution algorithm:
	// (both have changing lengths don't be fooled by their initial lengths)
	alt_stack := make([]uint, len(imm_pc))
	desired_path := make([]string, len(imm_pc))
	// Set them up using the first path condition
	for index, constraint := range imm_pc {
		alt_stack[index] = uint(index)
		desired_path[index] = constraint
	}
	// Main loop
	for len(alt_stack) > 0 {
		// TODO all of this
	}
}
