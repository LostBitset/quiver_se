package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func (uprgm Microprogram) SiMReQProcessPCs(
	in_pcs chan []string,
	bug_signal chan struct{},
) {
	for pc := range in_pcs {
		grouped_by_transition := make(map[SimpleEdgeDesc][]string)
		current_transition_constraint := make([]string, 0)
	groupPcSegmentsLoop:
		for _, item := range pc {
			if strings.HasPrefix(item, "@__RAW__;;@RICHPC:") {
				if strings.HasPrefix(item, "@__RAW__;;@RICHPC:was-segment ") {
					// TODO
					continue groupPcSegmentsLoop
				}
				log.Warn("Unknown rich path condition marker.")
				continue groupPcSegmentsLoop
			}
			// TODO
		}
	}
}
