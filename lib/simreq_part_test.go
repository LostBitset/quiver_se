package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSiMReQPart(t *testing.T) {
	in_updates := make(chan Augmented[
		QuiverUpdate[int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex]],
		[]SMTFreeFun[string, string],
	])
	out_models := make(chan string)
	var idsrc IdSource
	sys := SMTLibv2StringSystem{idsrc}
	dmtq, top_node, fail_node := StartSiMReQ[int, string, string, string, string, SMTLibv2StringSolvedCtx](
		in_updates, out_models, sys,
	)
	update_dmt := NewDMT[WithId_H[string], QuiverIndex]()
	intended_node := dmtq.NewIntendedNode(47, &update_dmt)
	intended_node_cb_backing := func(index QuiverIndex) {
		go func() {
			in_updates <- Augmented[
				QuiverUpdate[int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex]],
				[]SMTFreeFun[string, string],
			]{
				QuiverUpdate[
					int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex],
				]{
					index,
					dmtq.ParameterizeIndex(fail_node),
					StdlibMapToPHashMap(
						map[Literal[WithId_H[string]]]struct{}{
							{
								WithId_H[string]{"(= a b)", idsrc.Gen()},
								true,
							}: {},
							{
								WithId_H[string]{"(> b 0)", idsrc.Gen()},
								true,
							}: {},
						},
					),
				},
				[]SMTFreeFun[string, string]{
					{"a", []string{}, "Int"},
					{"b", []string{}, "Int"},
				},
			}
			close(in_updates)
		}()
	}
	intended_node.cb = &intended_node_cb_backing
	in_updates <- Augmented[
		QuiverUpdate[int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex]],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex],
		]{
			top_node,
			intended_node,
			StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"(= a 1)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(= a a)", idsrc.Gen()},
						false,
					}: {},
				},
			),
		},
		[]SMTFreeFun[string, string]{
			{"a", []string{}, "Int"},
			{"b", []string{}, "Int"},
		},
	}
	models := make([]string, 0)
	for model := range out_models {
		models = append(models, model)
	}
	assert.Equal(t, 1, len(models))
	model := models[0]
	assert.Contains(t, model, "(define-fun a () Int\n    1)")
	assert.Contains(t, model, "(define-fun b () Int\n    1)")
}
