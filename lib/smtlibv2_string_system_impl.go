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

func (sctx SMTLibv2StringSolvedCtx) IsSat() (is bool) {
	is = sctx.sat
	return
}

func (sctx SMTLibv2StringSolvedCtx) GetModel() (model *string) {
	model = sctx.model
	return
}

func (sctx SMTLibv2StringSolvedCtx) ExtractMUS() (mus *[]int) {
	mus = sctx.mus
	return
}
