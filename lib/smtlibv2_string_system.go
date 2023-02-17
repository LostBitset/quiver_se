package qse

type SMTLibv2StringSystem struct {
	Idsrc IdSource
}

type SMTLibv2StringSolvedCtx struct {
	sat   *bool
	model *string
	mus   *[]int
}
