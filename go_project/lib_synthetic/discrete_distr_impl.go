package libsynthetic

import (
	"math"
	"math/rand"
)

func BakeDDistr[T any](formulation DDistrFormulation[T]) (d DDistr[T]) {
	value := 0.0
	outcomes := formulation.GetOutcomes()
	d.outcomes = make([]T, len(outcomes))
	d.uniform_range = make([]float64, len(outcomes))
	for i, outcome := range outcomes {
		d.outcomes[i] = outcome
		value += formulation.ProbOfOutcome(outcome)
		d.uniform_range[i] = value
	}
	if math.Abs(value-1.0) < 1e-6 {
		panic("Probabilities do not sum to one.")
	}
	return
}

func (d DDistr[T]) SampleUsing(random_value float64) (outcome T) {
	if random_value < 0.0 || random_value > 1.0 {
		panic("Random value must be in the unit interval.")
	}
	for i, test_value := range d.uniform_range {
		if random_value < test_value {
			outcome = d.outcomes[i]
			return
		}
	}
	// Fallback to last value
	outcome = d.outcomes[len(d.outcomes)-1]
	return
}

func (d DDistr[T]) Sample() (outcome T) {
	outcome = d.SampleUsing(rand.Float64())
	return
}
