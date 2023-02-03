package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMTLibv2StringSystemUnsat(t *testing.T) {
	var sys SMTSystem[IdLiteral[string], string, string, string, SMTLibv2StringSolvedCtx]
	var idsrc IdSource
	sys = SMTLibv2StringSystem{idsrc}
	sctx := sys.CheckSat(
		[]IdLiteral[string]{
			{
				WithId_H[string]{"(< a b)", idsrc.Gen()},
				true,
			},
			{
				WithId_H[string]{"(> a b)", idsrc.Gen()},
				true,
			},
			{
				WithId_H[string]{"(= a b)", idsrc.Gen()},
				false,
			},
		},
		[]SMTFreeFun[string, string]{
			{
				"a",
				[]string{},
				"Int",
			},
			{
				"b",
				[]string{},
				"Int",
			},
		},
	)
	assert.False(t, *sctx.IsSat())
	assert.Nil(t, sctx.GetModel())
	assert.ElementsMatch(
		t,
		[]int{0, 1},
		*sctx.ExtractMUS(),
	)
}
