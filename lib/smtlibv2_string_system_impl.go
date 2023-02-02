package qse

import "strings"

func (sys SMTLibv2StringSystem) MakeAtom(expr string) (atom WithId_H[string]) {
	atom = WithId_H[string]{
		expr,
		sys.idsrc.Gen(),
	}
	return
}

func (sys SMTLibv2StringSystem) CheckSat(
	conjunction []IdLiteral[string],
	free_funs []SMTFreeFun[string, string],
) (sctx SMTLibv2StringSolvedCtx) {
	var sb strings.Builder
	sb.WriteString(
		sys.Prelude(),
	)
	sb.WriteString(
		sys.GenerateDecls(free_funs),
	)
	for i, lit := range conjunction {
		clause := SMTLibv2ExpandStringLiteral(lit)
		clause_marked := sys.MarkClauseIndex(clause, i)
		assertion := SMTLibv2WrapAssertion(clause_marked)
		sb.WriteString(assertion)
	}
	resp := QueryZ3SMTLibv2Complete(sb.String())
	sctx = ParseSMTLibv2StringSolvedCtx(resp)
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
