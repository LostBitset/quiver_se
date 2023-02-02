package qse

type LiteralSystemSolverless[ATOM any] struct {
	idsrc IdSource
}

type NumericId = uint32

type IdH[T any] struct {
	value T
	id    NumericId
}

type IdSource struct {
	next_id NumericId
}
