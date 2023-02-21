package lib

import "os/exec"

func InstrumentFunctionInfo(location string) {
	command := exec.Command("../lib/run_function_info_inplace"+".sh", location)
	err := command.Start()
	if err != nil {
		panic(err)
	}
	err = command.Wait()
	if err != nil {
		panic(err)
	}
}
