package qse

type SMTFreeFun[IDENT any, SORT any] struct {
	Name IDENT
	Args []SORT
	Ret  SORT
}
