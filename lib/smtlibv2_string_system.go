package qse

type SMTLibv2StringSystem struct {
	idsrc IdSource
}

type SMTLibv2StringSolvedCtx struct {
	sat   bool
	model *string
	mus   *[]int
}
