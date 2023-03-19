package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMTLib2VAStringSystemUnsat(t *testing.T) {
	var sys SMTSystem[IdLiteral[string], string, string, string, SMTLib2VAStringSolvedCtx]
	var idsrc IdSource
	sys = SMTLib2VAStringSystem{idsrc}
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

func TestSMTLib2VAStringSystemSat(t *testing.T) {
	var sys SMTSystem[IdLiteral[string], string, string, string, SMTLib2VAStringSolvedCtx]
	var idsrc IdSource
	sys = SMTLib2VAStringSystem{idsrc}
	sctx := sys.CheckSat(
		[]IdLiteral[string]{
			{
				WithId_H[string]{"(<= a b)", idsrc.Gen()},
				true,
			},
			{
				WithId_H[string]{"(< a b)", idsrc.Gen()},
				false,
			},
			{
				WithId_H[string]{"(= a 4)", idsrc.Gen()},
				true,
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
	assert.True(t, *sctx.IsSat())
	assert.Nil(t, sctx.ExtractMUS())
	model := *sctx.GetModel()
	assert.Contains(t, model, "(define-fun a () Int\n    4)")
	assert.Contains(t, model, "(define-fun b () Int\n    4)")
}
