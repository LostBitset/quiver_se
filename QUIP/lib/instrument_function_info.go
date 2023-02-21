package lib

import (
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
	err := command.Start()
	if err != nil {
		panic(err)
	}
	err = command.Wait()
	if err != nil {
		panic(err)
	}
}
