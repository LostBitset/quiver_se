package lib

import (
	"fmt"
	"os/exec"
)

const SIMPLE_DSE_LOW_FREQUENCY_CYCLE_WAIT_TIME_MILLIS = 500

func RunSimpleDSELowFrequency(msg_prefix string) {
	RunSimpleDSE(msg_prefix, SIMPLE_DSE_LOW_FREQUENCY_CYCLE_WAIT_TIME_MILLIS)
}

func RunSimpleDSE(msg_prefix string, cycle_wait_time int) {
	cmd := exec.Command(
		"run_simple_dse"+".sh",
		msg_prefix,
		fmt.Sprintf(
			"--cycle-wait-time=%d",
			cycle_wait_time,
		),
	)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	go func() {
		defer fmt.Println("[QUIP:simple_dse.go/waiting_goroutine] Simple DSE process ended.")
		cmd.Wait()
	}()
}
