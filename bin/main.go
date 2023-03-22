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
		dse_n_bugs_chan := make(chan int)
		go func() {
			dse_n_bugs_chan <- uprgm.RunDSE()
		}()
		select {
		case dse_n_bugs := <-dse_n_bugs_chan:
			if dse_n_bugs != -1 {
				count_dse += dse_n_bugs
			}
		case <-time.After(timeout):
			fmt.Println("[bin:main] DSE timed out (this is normal).")
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
