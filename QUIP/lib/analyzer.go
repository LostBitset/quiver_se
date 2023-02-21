package lib

import (
	"bufio"
	"fmt"
	"os/exec"
)

func RunAnalyzer(target string, msg_prefix string) {
	cmd := exec.Command("../lib/run_analyzer"+".sh", target, msg_prefix)
	so, so_err := cmd.StdoutPipe()
	se, se_err := cmd.StderrPipe()
	if so_err != nil {
		panic(so_err)
	}
	go func() {
		sc := bufio.NewScanner(so)
		for sc.Scan() {
			fmt.Println("[QUIP:simple_dse.go::StdoutPipe] " + sc.Text())
		}
	}()
	if se_err != nil {
		panic(se_err)
	}
	go func() {
		sc := bufio.NewScanner(se)
		for sc.Scan() {
			fmt.Println("[QUIP:simple_dse.go::StderrPipe] " + sc.Text())
		}
	}()
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
