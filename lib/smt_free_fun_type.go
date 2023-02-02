package qse

type SMTFreeFun[IDENT any, SORT any] struct {
	name IDENT
	args []SORT
	ret  SORT
}
