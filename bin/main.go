package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.WarnLevel)
	uprgm_gen := BuildEvaluationMicroprogramGenerator()
	uprgm := uprgm_gen.RandomMicroprogram()
	n_samples := 1
	timeout := 5 * time.Second
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
				bug_signal_values, false, nil, false, -1, uprgm.top_state,
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

func BuildEvaluationMicroprogramGenerator() (uprgm_gen MicroprogramGenerator) {
	ops, vals := GetStandardItems()
	constraint_gen := ConstraintGenerator{
		n_depth_mean:   2.0,
		n_depth_stddev: 1.5,
		ops:            ops,
		vals:           vals,
		next_var_id:    pto(0),
	}
	var_sorts := SimpleDDistr[Sort]{
		map[Sort]float64{
			RealSort: 0.7,
			BoolSort: 0.3,
		},
	}
	var_sorts_distr := BakeDDistr[Sort](var_sorts)
	constraint_gen.AddVariables(4, var_sorts_distr, 0.75)
	uprgm_gen = MicroprogramGenerator{
		n_states:          30,
		p_transition:      0.8,
		avg_n_transitions: 8.0,
		p_fallible:        0.9,
		n_entry_samples:   7,
		n_tree_nonleaf:    3,
		constraintgen:     constraint_gen,
	}
	return
}
