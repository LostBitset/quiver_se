package qse

import (
	"sync"
	"testing"
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
	wg sync.WaitGroup,
) (
	update QuiverUpdate[
		int,
		PHashMap[Literal[WithId_H[string]], struct{}],
		*DMT[WithId_H[string], QuiverIndex],
	],
) {
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
	nodes := []int{1, 2, 3}
}
