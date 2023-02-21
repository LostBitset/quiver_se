package lib

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func InstrumentFunctionInfo(location string) {
	location_without_js, _ := strings.CutSuffix(location, ".js")
	command := exec.Command(
		"../lib/run_function_info_inplace"+".sh",
		location,
		location_without_js+"._fninf.js",
	)
	so, so_err := command.StdoutPipe()
	se, se_err := command.StderrPipe()
	if so_err != nil {
		panic(so_err)
	}
	go func() {
		sc := bufio.NewScanner(so)
		for sc.Scan() {
			fmt.Println("[QUIP:instrument_function_info.go::StdoutPipe] " + sc.Text())
		}
	}()
	if se_err != nil {
		panic(se_err)
	}
	go func() {
		sc := bufio.NewScanner(se)
		for sc.Scan() {
			fmt.Println("[QUIP:instrument_function_info.go::StderrPipe] " + sc.Text())
		}
	}()
	err := command.Start()
	if err != nil {
		panic(err)
	}
	err = command.Wait()
	if err != nil {
		panic(err)
	}
}
