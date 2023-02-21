package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

const SMTLIBV2_MODEL_DEFINE_FIRST_LINE_PREFIX = "  (define-fun "

func HandleAnalyze(msg eidin.Analyze, target string) {
	fmt.Println("[js_concolic:AnalyzerProcess] Processing Analyze message...")
	fmt.Println(msg)
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
					curr_sort = fields_after[len(fields_after)-1]
					looking_for_value = true
				} else {
					looking_for_value = false
				}
			} else {
				if looking_for_value && strings.HasPrefix(line, "    ") {
					curr_value += parseModelValueLine(line, curr_sort)
				}
			}
		}
		if curr_value != "" {
			mapping_raw[curr_key] = curr_value
		}
		fmt.Println("[js_concolic:AnalyzerProcess] Generated assignments, serializing...")
		var json strings.Builder
		first := true
		json.WriteRune('{')
		for k, v := range mapping_raw {
			if !first {
				json.WriteRune(',')
			}
			first = false
			json.WriteString(
				fmt.Sprintf(
					`"%s":%s`,
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
	if msg.GetSingleCallback() {
		addl_args = append(addl_args, "--single-callback")
	}
	std_args := []string{
		"../run_analysis_code" + ".sh",
		target,
	}
	std_args = append(std_args, addl_args...)
	fmt.Println("[js_concolic:AnalyzerProcess] Starting (jalangi2-based) PC analysis...")
	cmd := exec.Command("/bin/sh", std_args...)
	stdout, err_stdout := cmd.StdoutPipe()
	stderr, err_stderr := cmd.StderrPipe()
	cmd.Start()
	if err_stdout == nil {
		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()
	} else {
		panic(err_stdout)
	}
	if err_stderr == nil {
		go func() {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				fmt.Println("[::StderrPipe] " + scanner.Text())
			}
		}()
	} else {
		panic(err_stderr)
	}
	cmd.Wait()
}

func parseModelValueLine(line string, sort string) (repr string) {
	fmt.Printf(
		"parseModelValueLine(\"%s\", \"%s\")\n",
		line, sort,
	)
	switch sort {
	case "Real":
		repr = line[4 : len(line)-1]
	case "Int":
		repr = line[4 : len(line)-1]
	default:
		repr = "undefined"
	}
	return
}
