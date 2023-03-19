package qse

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SiMReQTestingQuiverIndexMap struct {
	m  map[int]QuiverIndex
	mu sync.Mutex
}

func NewSiMReQTestingQuiverIndexMap() (index_map SiMReQTestingQuiverIndexMap) {
	index_map = SiMReQTestingQuiverIndexMap{
		m: make(map[int]QuiverIndex),
	}
	return
}

func MakeNodeWithCallbackForSiMReQTesting(
	dmtq Quiver[
		int,
		PHashMap[Literal[WithId_H[string]], struct{}],
		*DMT[WithId_H[string], QuiverIndex],
	],
	top_node QuiverIndex,
	value int,
	save_index *SiMReQTestingQuiverIndexMap,
	wg *sync.WaitGroup,
) (
	update QuiverUpdate[
		int,
		PHashMap[Literal[WithId_H[string]], struct{}],
		*DMT[WithId_H[string], QuiverIndex],
	],
) {
	wg.Add(1)
	update_dmt := NewDMT[WithId_H[string], QuiverIndex]()
	intended_node := dmtq.NewIntendedNode(value, &update_dmt)
	intended_node_cb_backing := func(index QuiverIndex) {
		go func() {
			save_index.mu.Lock()
			save_index.m[value] = index
			save_index.mu.Unlock()
			wg.Done()
		}()
	}
	intended_node.cb = &intended_node_cb_backing
	update = QuiverUpdate[
		int,
		PHashMap[Literal[WithId_H[string]], struct{}],
		*DMT[WithId_H[string], QuiverIndex],
	]{
		top_node,
		intended_node,
		nil,
	}
	return
}

func TestSiMReQPartReal(t *testing.T) {
	in_updates := make(chan Augmented[
		QuiverUpdate[int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex]],
		[]SMTFreeFun[string, string],
	])
	out_models := make(chan string)
	var idsrc IdSource
	sys := SMTLib2VAStringSystem{idsrc}
	dmtq, top_node, fail_node, _ := StartSiMReQ[int, string, string, string, string, SMTLib2VAStringSolvedCtx](
		in_updates, out_models, sys, nil,
	)
	nodes := []int{11, 12, 13}
	index_map_container_backing := NewSiMReQTestingQuiverIndexMap()
	index_map_container := &index_map_container_backing
	var wg sync.WaitGroup
	smt_free_funs := []SMTFreeFun[string, string]{
		{"x", []string{}, "Int"},
		{"y", []string{}, "Int"},
	}
	for _, node := range nodes {
		in_updates <- Augmented[
			QuiverUpdate[int, PHashMap[Literal[WithId_H[string]], struct{}], *DMT[WithId_H[string], QuiverIndex]],
			[]SMTFreeFun[string, string],
		]{
			MakeNodeWithCallbackForSiMReQTesting(
				dmtq,
				top_node,
				node,
				index_map_container,
				&wg,
			),
			smt_free_funs,
		}
	}
	wg.Wait()
	index_map := index_map_container.m
	// begin edges
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			top_node,
			dmtq.ParameterizeIndex(index_map[11]),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(e1)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"@__RAW__(*/decl-var/* **jsvar_z)\n(*/write-var/* **jsvar_z *{{x}}*)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"@__RAW__(*/decl-var/* **jsvar_a)\n(*/write-var/* **jsvar_a *{{false}}*)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(< (*/read-var/* **jsvar_z) y)", idsrc.Gen()},
						true,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			index_map[11],
			dmtq.ParameterizeIndex(index_map[12]),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(e2)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(< (*/read-var/* **jsvar_z) y)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"@__RAW__(*/write-var/* **jsvar_z *{{(+ (*/read-var/* **jsvar_z) 2)}}*)", idsrc.Gen()},
						true,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			index_map[12],
			dmtq.ParameterizeIndex(index_map[13]),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(e3)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(= (*/read-var/* **jsvar_z) y)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(*/read-var/* **jsvar_a)", idsrc.Gen()},
						false,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			index_map[12],
			dmtq.ParameterizeIndex(index_map[11]),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(e4)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(not (and (= (*/read-var/* **jsvar_z) y) (not (*/read-var/* **jsvar_a))))", idsrc.Gen()},
						false,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			index_map[13],
			dmtq.ParameterizeIndex(index_map[11]),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(e5)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"@__RAW__(*/write-var/* **jsvar_z *{{(- (*/read-var/* **jsvar_z) 1)}}*)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(= (*/read-var/* **jsvar_z) 2)", idsrc.Gen()},
						true,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			index_map[13],
			dmtq.ParameterizeIndex(index_map[11]),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(e6)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"@__RAW__(*/write-var/* **jsvar_z *{{(- (*/read-var/* **jsvar_z) 1)}}*)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(= (*/read-var/* **jsvar_z) 2)", idsrc.Gen()},
						false,
					}: {},
					{
						WithId_H[string]{"@__RAW__(*/write-var/* **jsvar_a *{{true}}*)", idsrc.Gen()},
						true,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	in_updates <- Augmented[
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		],
		[]SMTFreeFun[string, string],
	]{
		QuiverUpdate[
			int,
			PHashMap[Literal[WithId_H[string]], struct{}],
			*DMT[WithId_H[string], QuiverIndex],
		]{
			index_map[11],
			dmtq.ParameterizeIndex(fail_node),
			pto(StdlibMapToPHashMap(
				map[Literal[WithId_H[string]]]struct{}{
					{
						WithId_H[string]{"@__RAW__;; !EDGE(F1)", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(= (*/read-var/* **jsvar_z) (+ y 1))", idsrc.Gen()},
						true,
					}: {},
					{
						WithId_H[string]{"(*/read-var/* **jsvar_a)", idsrc.Gen()},
						true,
					}: {},
				},
			)),
		},
		smt_free_funs,
	}
	// end edges
	close(in_updates)
	models := make([]string, 0)
	for model := range out_models {
		models = append(models, model)
	}
	assert.NotEqual(t, len(models), 0)
}
