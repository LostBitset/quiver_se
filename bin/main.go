package main

import (
	"fmt"
	"time"
)

func main() {
	uprgm_gen := BuildEvaluationMicroprogramGenerator()
	uprgm := uprgm_gen.RandomMicroprogram()
	n_samples := 1
	timeout := 5 * time.Second
	count_dse := EvaluateAlgorithm(
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
			uprgm.RunDSEContinuously(bug_signal_values, false, nil)
		},
		uprgm, n_samples, timeout, "dse",
	)
	count_simreq := EvaluateAlgorithm(
		func(uprgm Microprogram, bug_signal chan struct{}) {
			uprgm.RunSiMReQ(bug_signal)
		},
		uprgm, n_samples, timeout, "simreq",
	)
	fmt.Println("--- FINAL RESULTS ---")
	fmt.Printf("Generated a total of %d programs.\n", n_samples)
	fmt.Printf("DSE    found %d bugs.\n", count_dse)
	fmt.Printf("SiMReQ found %d bugs.\n", count_simreq)
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
		p_fallible:        0.8,
		n_entry_samples:   7,
		n_tree_nonleaf:    5,
		constraintgen:     constraint_gen,
	}
	return
}
