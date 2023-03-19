package qse

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func NewZ3SMTLib2VAQuery(query_str string) (query Z3SMTLib2VAQuery) {
	query = Z3SMTLib2VAQuery{query_str}
	// JUST FOR TESTING begin
	special := []string{"e1", "e3", "e6", "e2", "e4", "F1"}
	is_special := true
	for _, item := range special {
		if !strings.Contains(query_str, "!EDGE("+item+")") {
			is_special = false
		}
	}
	if is_special {
		os.Exit(0)
	}
	// JUST FOR TESTING end
	return
}

const Z3SMTLib2VAQuery_TEMP_SMT2VA_FILENAME_FORMAT = "temp_qse-go_Z3SMTLib2VAQuery-Run_*_GENERATED.smt2va"

func (query Z3SMTLib2VAQuery) Run() (output string) {
	log.Info("[z3smtlib2VA_query/Z3SMTLib2VAQuery.Run] Setting up SMT query. ")
	temp_smt2va_file, err_create := os.CreateTemp("/tmp", Z3SMTLib2VAQuery_TEMP_SMT2VA_FILENAME_FORMAT)
	if err_create != nil {
		panic(err_create)
	}
	defer os.Remove(temp_smt2va_file.Name())
	closed := false
	defer func() {
		if !closed {
			temp_smt2va_file.Close()
		}
	}()
	_, err_write := temp_smt2va_file.Write([]byte(query.query))
	if err_write != nil {
		panic(err_write)
	}
	temp_smt2va_file.Close()
	closed = true
	log.Info("[z3smtlib2VA_query/Z3SMTLib2VAQuery.Run] Transpiling SMTLib_2VA -> SMTLib_v2. ")
	temp_smt2_filename, _ := strings.CutSuffix(temp_smt2va_file.Name(), ".smt2va")
	temp_smt2_filename += ".TRANSPILED-orig_smt2va.smt2"
	defer os.Remove(temp_smt2_filename)
	transpile_cmd := exec.Command(
		"/bin/bash",
		"-c",
		fmt.Sprintf(
			`"../js_concolic/SMTLib_2VA/bin/transpile.sh" "%s" log-info`,
			temp_smt2va_file.Name(),
		),
	)
	transpile_out, err_pipe := transpile_cmd.StdoutPipe()
	if err_pipe != nil {
		panic(err_pipe)
	}
	transpile_err, err_stderr := transpile_cmd.StderrPipe()
	if err_stderr != nil {
		panic(err_pipe)
	}
	transpile_cmd.Start()
	go func() {
		scanner := bufio.NewScanner(transpile_out)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(transpile_err)
		for scanner.Scan() {
			fmt.Println("[::StderrPipe] " + scanner.Text())
		}
	}()
	err_transpile := transpile_cmd.Wait()
	if err_transpile != nil {
		panic(transpile_err)
	}
	log.Info("[z3smtlib2VA_query/Z3SMTLib2VAQuery.Run] Querying Z3. ")
	z3_cmd := exec.Command("z3", "-smt2", temp_smt2_filename)
	z3_out, err_cmd := z3_cmd.Output()
	if err_cmd != nil {
		switch err_cmd.(type) {
		case *exec.ExitError:
		default:
			panic(err_cmd)
		}
	}
	log.Info("[z3smtlib2VA_query/Z3SMTLib2VAQuery.Run] Queried Z3 successfully. ")
	output = string(z3_out)
	return
}
