package qse

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func NewZ3SMTLibv2Query(query_str string) (query Z3SMTLibv2Query) {
	fmt.Println("[REPORT] [EVAL-INFO] SOLVERQUERY")
	query = Z3SMTLibv2Query{query_str}
	return
}

const Z3SMTLibv2Query_TEMP_SMT2_FILENAME_FORMAT = "temp_qse-go_Z3SMTLibv2Query-Run_*_GENERATED.smt2"

func (query Z3SMTLibv2Query) Run() (output string) {
	log.Info("[z3smtlibv2_query/Z3SMTLibv2Query.Run] Setting up SMT query. ")
	temp_smt2_file, err_create := os.CreateTemp("/tmp", Z3SMTLibv2Query_TEMP_SMT2_FILENAME_FORMAT)
	if err_create != nil {
		panic(err_create)
	}
	defer os.Remove(temp_smt2_file.Name())
	closed := false
	defer func() {
		if !closed {
			temp_smt2_file.Close()
		}
	}()
	_, err_write := temp_smt2_file.Write([]byte(query.query))
	if err_write != nil {
		panic(err_write)
	}
	temp_smt2_file.Close()
	closed = true
	log.Info("[z3smtlibv2_query/Z3SMTLibv2Query.Run] Querying Z3. ")
	z3_cmd := exec.Command("z3", "-smt2", temp_smt2_file.Name())
	z3_out, err_cmd := z3_cmd.Output()
	if err_cmd != nil {
		switch err_cmd.(type) {
		case *exec.ExitError:
		default:
			panic(err_cmd)
		}
	}
	log.Info("[z3smtlibv2_query/Z3SMTLibv2Query.Run] Queried Z3 successfully. ")
	output = string(z3_out)
	return
}
