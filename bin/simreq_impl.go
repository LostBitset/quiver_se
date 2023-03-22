package main

import (
	"strconv"
	"strings"

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

func (uprgm Microprogram) SiMReQProcessPCs(
	in_pcs chan []string,
	bug_signal chan struct{},
) {
	// Setup everything necessary
	// TODO
	// Start SiMReQ
	// TODO
	// Listen for bugs
	// TODO
	for pc := range in_pcs {
		// Group the segmented path condition by segments (which represent transitions)
		grouped_by_transition := make(map[SimpleMicroprogramTransitionDesc][]string)
		current_transition_constraint := make([]string, 0)
	groupPcSegmentsLoop:
		for _, item := range pc {
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
		// TODO
	}
}
