package main

import (
	"fmt"
	"time"
)

func main() {
	uprgm_gen := BuildEvaluationMicroprogramGenerator()
	count_dse := 0
	n_samples := 2
	timeout := 3 * time.Second
	for i := 0; i < n_samples; i++ {
		uprgm := uprgm_gen.RandomMicroprogram()
		dse_bug_signal_orig := make(chan struct{})
		dse_bug_signal := make(chan struct{})
		dse_end_signal := make(chan struct{})
		go uprgm.RunDSEContinuously(dse_bug_signal_orig)
		go func() {
			for range dse_bug_signal_orig {
				dse_bug_signal <- struct{}{}
			}
			dse_end_signal <- struct{}{}
		}()
		dse_timeout := time.After(timeout)
	tallyDSEBugsLoop:
		for {
		tallyDSEBugsSelect:
			select {
			case <-dse_bug_signal:
				count_dse++
				break tallyDSEBugsSelect
			case <-dse_timeout:
				fmt.Println("[bin:main] DSE timed out (this is normal).")
				break tallyDSEBugsLoop
			case <-dse_bug_signal:
				break tallyDSEBugsLoop
			}
		}
	}
	fmt.Println("--- FINAL RESULTS ---")
	fmt.Printf("Generated a total of %d programs.\n", n_samples)
	fmt.Printf("DSE found %d bugs.\n", count_dse)
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
	constraint_gen.AddVariables(4, var_sorts_distr, 0.8)
	uprgm_gen = MicroprogramGenerator{
		n_states:          20,
		p_transition:      0.65,
		avg_n_transitions: 2.0,
		p_fallible:        0.4,
		n_entry_samples:   5,
		n_tree_nonleaf:    4,
		constraintgen:     constraint_gen,
	}
	return
}
