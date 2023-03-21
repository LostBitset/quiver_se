package main

import (
	"sort"
	"strings"
)

func FilterModelFromZ3(model string) (filtered string) {
	lines := strings.Split(model, "\n")
	def_map := make(map[string]string) // def line -> all def content
	var active_sb strings.Builder
	active_def_line := "<parsefailed>"
	for _, line := range lines {
		if strings.HasPrefix(line, "    ") {
			// def content lines
			active_sb.WriteString(line)
		} else if strings.HasPrefix(line, "  ") {
			// def lines
			if active_def_line != "<parsefailed>" {
				def_map[active_def_line] = active_sb.String()
			}
			active_sb = *new(strings.Builder)
			active_def_line = line
		}
	}
	// Add final
	if active_def_line != "<parsefailed>" {
		def_map[active_def_line] = active_sb.String()
	}
	// Make a map of the keys for sorting
	def_map_keys := make([]string, 0)
buildSliceOfDefLinesLoop:
	for k := range def_map {
		k := k
		if strings.Contains(k, "(define-fun ga_") {
			continue buildSliceOfDefLinesLoop
		}
		def_map_keys = append(def_map_keys, k)
	}
	sort.Strings(def_map_keys)
	var final_sb strings.Builder
	for _, def_line := range def_map_keys {
		final_sb.WriteString(def_line)
		final_sb.WriteString(def_map[def_line])
		final_sb.WriteRune('\n')
	}
	filtered = final_sb.String()
	return
}
