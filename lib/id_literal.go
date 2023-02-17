package qse

type IdLiteral[ATOM comparable] Literal[WithId_H[ATOM]]

type NumericId = uint32

type WithId_H[T comparable] struct {
	Value T
	Id    NumericId
}

type IdSource struct {
	next_id NumericId
}
