package qse

// A system which provides basic interfaces to an SMT solver
type SMTSystem[
	EXPR Hashable,
	IDENT any,
	SORT any,
	MODEL any,
	SCTX SMTSolvedContext[MODEL],
] interface {
	CheckSat([]EXPR, []SMTFreeFun[IDENT, SORT]) SCTX
}

// A context in which an SMT solver has been invoked and results are available
type SMTSolvedContext[MODEL any] interface {
	IsSat() *bool
	GetModel() *MODEL
	ExtractMUS() *[]int
}
