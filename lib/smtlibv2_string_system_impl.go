package qse

import (
	"fmt"
	"regexp"
	"strings"
)

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
		sys.Prologue(),
	)
	sb.WriteString(
		sys.GenDecls(free_funs),
	)
	sb.WriteString(`
	;; Clauses of the conjunction, as assertions @@ <qse.SMTLibv2StringSystem>.CheckSat
	;; All should be named, but only at the top level @@ ...
	;; This allows MUC generation to produce an MUS @@ ...
	`)
	for i, lit := range conjunction {
		clause := sys.ExpandStringLiteral(lit)
		clause_marked := sys.MarkClauseIndex(clause, uint(i))
		assertion := SMTLibv2WrapAssertion(clause_marked)
		sb.WriteString(assertion)
		sb.WriteRune('\n')
	}
	sb.WriteString(sys.Epilogue())
	resp := QueryZ3SMTLibv2Complete(sb.String())
	sctx = sys.ParseSolvedCtx(resp)
	return
}

func (sys SMTLibv2StringSystem) Prologue() (part string) {
	part = `
	;; GENERATED SMTLibv2 code, targeting z3 @@ <qse.SMTLibv2StringSystem>.Prologue

	;; Force MUC generation @@ <qse.SMTLibv2StringSystem>.Prologue
	(set-option :produce-unsat-cores true)
	(set-option :smt.core.minimize true) ;; *z3 specific* @@ <qse.SMTLibv2StringSystem>.Prologue
	`
	return
}

// Don't change the order of the `%!...%` substitutions!
// This gets turned into a regex, so also avoid regex special chars
// Only ".()|[]" are checked right now
// Other code depends on it because golang doesn't support named capture groups in regexes!
const SMTLIBV2_STRING_SYSTEM_OUTPUT_FORMAT = `
;; Solver done, response below @@ ::smtlibv2-invocation @@ <SMTLibv2StringSystem>.Epilogue
[|resp
	[|resp.sat %!sat%|]
	[|resp.mus %!mus%|]
	[|resp.mdl %!mdl%|]|]
`

func (sys SMTLibv2StringSystem) Epilogue() (part string) {
	output_format := SMTLIBV2_STRING_SYSTEM_OUTPUT_FORMAT
	var sb strings.Builder
	sb.WriteString(`
	;; Get the results from the SMT solver @@ <SMTLibv2StringSystem>.Epilogue
	;; This mostly s-exprs, but special lists are delimited with "[|" and "|]" @@ ...
	;; This allows it to be matched quickly with regexes @@ ...
	`)
	for _, token := range strings.Split(output_format, "%") {
		switch token {
		case "!sat":
			sb.WriteString("(check-sat)")
		case "!mus":
			sb.WriteString("(get-unsat-core)")
		case "!mdl":
			sb.WriteString("(get-model)")
		default:
			sb.WriteString(
				fmt.Sprintf("(echo \"%s\")", token),
			)
		}
		sb.WriteRune('\n')
	}
	sb.WriteString(`
	;; MARK EOF
	`)
	part = sb.String()
	return
}

func (sys SMTLibv2StringSystem) ParseSolvedCtx(str string) (sctx SMTLibv2StringSolvedCtx) {
	re_comments := regexp.MustCompile(`;;[^\n]\n`)
	str_resp := strings.TrimSpace(
		string(
			re_comments.ReplaceAllLiteral(
				[]byte(str),
				[]byte{},
			),
		),
	)
	str_resp = strings.ReplaceAll(str_resp, ".", "\\.")
	str_resp = strings.ReplaceAll(str_resp, "(", "\\(")
	str_resp = strings.ReplaceAll(str_resp, ")", "\\)")
	str_resp = strings.ReplaceAll(str_resp, "|", "\\|")
	str_resp = strings.ReplaceAll(str_resp, "[", "\\[")
	str_resp = strings.ReplaceAll(str_resp, "]", "\\]")
	re_substitutions := regexp.MustCompile(`%![^%]%`)
	resp_regex := strings.TrimSpace(
		string(
			re_substitutions.ReplaceAllLiteral(
				[]byte(SMTLIBV2_STRING_SYSTEM_OUTPUT_FORMAT),
				[]byte(`([^\\|\\]]+)`),
			),
		),
	)
	re_resp := regexp.MustCompile(resp_regex)
	re_resp_output := re_resp.FindStringSubmatch(str_resp)
	capture_groups := re_resp_output[1:]
	sat_trimmed := strings.TrimSpace(capture_groups[0])
	t, f := true, false
	switch sat_trimmed {
	case "sat":
		model := strings.TrimSpace(capture_groups[2])
		sctx = SMTLibv2StringSolvedCtx{
			&t,
			&model,
			nil,
		}
	case "unsat":
		mus_str := strings.TrimSpace(capture_groups[1])
		mus_str = strings.ReplaceAll(mus_str, "(", "")
		mus_str = strings.ReplaceAll(mus_str, ")", "")
		mus_str = strings.ReplaceAll(mus_str, "ga_", "")

	}
}

func (sys SMTLibv2StringSystem) GenDecls(free_funs []SMTFreeFun[string, string]) (part string) {
	var sb strings.Builder
	sb.WriteString(`
	;; Declarations, generated from free_funs @@ <qse.SMTLibv2StringSystem>.GenDecls
	`)
	for _, free_fun := range free_funs {
		sb.WriteString(sys.DeclSExpr(free_fun))
		sb.WriteRune('\n')
	}
	part = sb.String()
	return
}

func (sys SMTLibv2StringSystem) DeclSExpr(free_fun SMTFreeFun[string, string]) (s_expr string) {
	s_expr = fmt.Sprintf(
		"(declare-fun %s (%s) %s)",
		free_fun.name,
		strings.Join(free_fun.args, " "),
		free_fun.ret,
	)
	return
}

func (sys SMTLibv2StringSystem) ExpandStringLiteral(lit IdLiteral[string]) (s_expr string) {
	s_expr = lit.value.value
	if !lit.eq {
		s_expr = fmt.Sprintf(
			"(not %s)",
			s_expr,
		)
	}
	return
}

func (sys SMTLibv2StringSystem) MarkClauseIndex(clause string, index uint) (clause_marked string) {
	clause_marked = fmt.Sprintf(
		"(! %s :named ga_%d)",
		clause, index,
	)
	return
}

func SMTLibv2WrapAssertion(clause string) (s_expr string) {
	s_expr = fmt.Sprintf(
		"(assert %s)",
		clause,
	)
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
