package lib

import (
	"fmt"
	"os/exec"
)

func RunAnalyzer(target string, msg_prefix string) {
	cmd := exec.Command("../lib/run_analyzer"+".sh", target, msg_prefix)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	go func() {
		defer fmt.Println("[QUIP:analyzer.go/waiting_goroutine] Analyzer process ended.")
		err := cmd.Wait()
		if err != nil {
			panic(err)
		}
	}()
}
