package qse

type SMTLib2VAStringSystem struct {
	Idsrc IdSource
}

type SMTLib2VAStringSolvedCtx struct {
	sat   *bool
	model *string
	mus   *[]int
}
