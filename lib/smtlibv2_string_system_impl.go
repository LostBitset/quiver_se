package qse

func (sys SMTLibv2StringSystem) MakeAtom(expr string) (atom WithId_H[string]) {
	atom = WithId_H[string]{
		expr,
		sys.idsrc.Gen(),
	}
	return
}

func (sys SMTLibv2StringSystem) CheckSat(conjunction []string) (sctx SMTLibv2StringSolvedCtx) {
	// TODO
	return
}
