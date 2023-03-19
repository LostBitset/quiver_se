package main

type DDistrFormulation[T any] interface {
	GetOutcomes() (outcomes []T)
	ProbOfOutcome(outcome T) (p float64)
}

type DDistr[T any] struct {
	outcomes      []T
	uniform_range []float64
}
