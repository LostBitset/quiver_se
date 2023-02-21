package lib

import (
	"bufio"
	"fmt"
	"os/exec"
)

const SIMPLE_DSE_LOW_FREQUENCY_CYCLE_WAIT_TIME_MILLIS = 500

func RunSimpleDSELowFrequency(msg_prefix string) {
	RunSimpleDSE(
		msg_prefix,
		SIMPLE_DSE_LOW_FREQUENCY_CYCLE_WAIT_TIME_MILLIS,
		false,
		true,
	)
}

func RunSimpleDSE(
	msg_prefix string,
	cycle_wait_time int,
	single_callback_mode bool,
	perist_mode bool,
) {
	var cmd *exec.Cmd
	if perist_mode {
		if single_callback_mode {
			cmd = exec.Command(
				"../lib/run_simple_dse"+".sh",
				msg_prefix,
				fmt.Sprintf(
					"--cycle-wait-time=%d",
					cycle_wait_time,
				),
				"--single-callback",
				"--rename-persist-path-conditions",
			)
		} else {
			cmd = exec.Command(
				"../lib/run_simple_dse"+".sh",
				msg_prefix,
				fmt.Sprintf(
					"--cycle-wait-time=%d",
					cycle_wait_time,
				),
				"--rename-persist-path-conditions",
			)
		}
	} else {
		if single_callback_mode {
			cmd = exec.Command(
				"../lib/run_simple_dse"+".sh",
				msg_prefix,
				fmt.Sprintf(
					"--cycle-wait-time=%d",
					cycle_wait_time,
				),
				"--single-callback",
			)
		} else {
			cmd = exec.Command(
				"../lib/run_simple_dse"+".sh",
				msg_prefix,
				fmt.Sprintf(
					"--cycle-wait-time=%d",
					cycle_wait_time,
				),
			)
		}
	}
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
	KickstartDse(msg_prefix)
	go func() {
		defer fmt.Println("[QUIP:simple_dse.go/waiting_goroutine] Simple DSE process ended.")
		err := cmd.Wait()
		if err != nil {
			panic(err)
		}
	}()
}
