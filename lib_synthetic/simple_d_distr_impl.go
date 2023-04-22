package libsynthetic

func (d SimpleDDistr[T]) GetOutcomes() (outcomes []T) {
	outcomes = make([]T, len(d.outcomes))
	for outcome := range d.outcomes {
		outcomes = append(outcomes, outcome)
	}
	return
}

func (d SimpleDDistr[T]) ProbOfOutcome(outcome T) (prob float64) {
	prob = d.outcomes[outcome]
	return
}
