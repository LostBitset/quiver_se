package libsynthetic

func (d SimpleDDistr[T]) GetOutcomes() (outcomes []T) {
	outcomes = make([]T, len(d.Outcomes))
	for outcome := range d.Outcomes {
		outcomes = append(outcomes, outcome)
	}
	return
}

func (d SimpleDDistr[T]) ProbOfOutcome(outcome T) (prob float64) {
	prob = d.Outcomes[outcome]
	return
}
