package qse

import (
	"fmt"
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
	fmt.Print("start...")
	smr_config.Start()
	fmt.Println("ok")
	fmt.Print("send canidate...")
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
	fmt.Println("ok")
	fmt.Print("close canidates channel...")
	close(in_canidates)
	fmt.Println("ok")
	fmt.Print("recv models...")
	models := make([]string, 0)
	for model := range out_models {
		fmt.Print("(recv)")
		models = append(models, model)
	}
	fmt.Println("ok")
	assert.Equal(t, 1, len(models))
}
