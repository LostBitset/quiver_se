package qse

type LiteralSystemSolverless[ATOM any] struct {
	idsrc IdSource
}

type IdLiteral[T comparable] Literal[WithId_H[T]]

type NumericId = uint32

type WithId_H[T comparable] struct {
	value T
	id    NumericId
}

type IdSource struct {
	next_id NumericId
}
