package qse

type SMTSystemSolverless[EXPR any] interface {
	Not(EXPR) EXPR
}

// A system which provides basic interfaces to an SMT solver
type SMTSystem[EXPR any, MODEL any, SCTX SMTSolvedContext[MODEL]] interface {
	SMTSystemSolverless[EXPR]
	CheckSat([]EXPR) SCTX
}

// A context in which an SMT solver has been invoked and results are available
type SMTSolvedContext[MODEL any] interface {
	IsSat() bool
	GetModel() *MODEL
	ExtractMUS() []int
}
