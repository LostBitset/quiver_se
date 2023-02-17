package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"strings"
)

const SMTLIBV2_MODEL_DEFINE_FIRST_LINE_PREFIX = "  (define-fun "

func HandleAnalyze(msg eidin.Analyze, target string) {
	fmt.Println("[js_concolic:AnalyzerProcess] Processing Analyze message...")
	var addl_args []string
	if len(msg.GetModel()) > 0 {
		lines := strings.Split(msg.GetModel(), "\n")
		mapping_raw := make(map[string]string)
		var curr_key string
		var curr_sort string
		var curr_value string
		looking_for_value := false
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if strings.HasPrefix(line, SMTLIBV2_MODEL_DEFINE_FIRST_LINE_PREFIX) {
				if curr_value != "" {
					mapping_raw[curr_key] = curr_value
					curr_key, curr_sort, curr_value = "", "", ""
				}
				line_after := line[len(SMTLIBV2_MODEL_DEFINE_FIRST_LINE_PREFIX):]
				name, _, _ := strings.Cut(
					line_after,
					" ",
				)
				if !(strings.HasPrefix(name, "ga_") || strings.HasPrefix(name, "mhide_")) {
					curr_key = name
					fields_after := strings.Fields(line_after)
					curr_sort = fields_after[len(fields_after)]
					looking_for_value = true
				} else {
					looking_for_value = false
				}
			} else {
				if looking_for_value {
					curr_value += parseModelValueLine(line, curr_sort)
				}
			}
		}
		if curr_value != "" {
			mapping_raw[curr_key] = curr_value
		}
		var json strings.Builder
		first := true
		json.WriteRune('{')
		for k, v := range mapping_raw {
			if !first {
				json.WriteRune(',')
			}
			first = true
			json.WriteString(
				fmt.Sprintf(
					`"%s": %s`,
					k, v,
				),
			)
		}
		json.WriteRune('}')
		addl_args = []string{
			"json-model",
			json.String(),
		}
	} else {
		addl_args = make([]string, 0)
	}
}

func parseModelValueLine(line string, sort string) (repr string) {
	repr = "424242"
	return
}
