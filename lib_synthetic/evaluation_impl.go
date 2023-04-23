package libsynthetic

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func RunEvaluationMain() {
	logrus.SetLevel(logrus.WarnLevel)
	uprgm_gen := BuildEvaluationMicroprogramGenerator()
	uprgm := uprgm_gen.RandomMicroprogram()
	n_samples := 1
	timeout := 12 * time.Second
	EvaluateAlgorithm(
		func(uprgm Microprogram, bug_signal chan struct{}) {
			bug_signal_values := make(chan uint32)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("[main recovered]")
					}
				}()
				defer close(bug_signal)
				for range bug_signal_values {
					bug_signal <- struct{}{}
				}
			}()
			uprgm.RunDSEContinuously(
				bug_signal_values, false, nil, false, -1, uprgm.StateTop,
			)
		},
		uprgm, n_samples, timeout, "dse",
	)
	EvaluateAlgorithm(
		func(uprgm Microprogram, bug_signal chan struct{}) {
			uprgm.RunSiMReQ(bug_signal, false)
		},
		uprgm, n_samples, timeout, "simreq:simple",
	)
	EvaluateAlgorithm(
		func(uprgm Microprogram, bug_signal chan struct{}) {
			uprgm.RunSiMReQ(bug_signal, true)
		},
		uprgm, n_samples, timeout, "simreq:jitdse",
	)
}

func EvaluateAlgorithm(
	algorithm func(Microprogram, chan struct{}),
	uprgm Microprogram,
	n_samples int,
	timeout time.Duration,
	name string,
) (
	count int,
) {
	fmt.Println("[REPORT] [EVALUATING] " + name)
	count = 0
	for i := 0; i < n_samples; i++ {
		bug_signal_orig := make(chan struct{})
		bug_signal := make(chan struct{})
		end_signal := make(chan struct{})
		go algorithm(uprgm, bug_signal_orig)
		go func() {
			for range bug_signal_orig {
				bug_signal <- struct{}{}
			}
			end_signal <- struct{}{}
		}()
		timeout_chan := time.After(timeout)
	tallyBugsLoop:
		for {
		tallyBugsSelect:
			select {
			case <-bug_signal:
				fmt.Println("[REPORT] __FOUND_A_BUG__" + name)
				count++
				break tallyBugsSelect
			case <-timeout_chan:
				fmt.Println("[bin:main] Timed out (this is normal).")
				break tallyBugsLoop
			case <-end_signal:
				break tallyBugsLoop
			}
		}
	}
	return
}
