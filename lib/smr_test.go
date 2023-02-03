package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMR(t *testing.T) {
	in_canidates := make(chan SMTQueryDNFClause[string, string, string])
	out_models := make(chan string)
	var idsrc IdSource
	smr_config := NewSMRConfig[
		string, string, string, string, SMTLibv2StringSolvedCtx,
	](
		in_canidates,
		out_models,
		SMTLibv2StringSystem{idsrc},
	)
	smr_config.Start()
	in_canidates <- SMTQueryDNFClause[string, string, string]{
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
	}
	close(in_canidates)
	models := make([]string, 0)
	for model := range out_models {
		models = append(models, model)
	}
	assert.Equal(t, 1, len(models))
}
