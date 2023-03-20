package main

import (
	qse "LostBitset/quiver_se/lib"
)

type SimpleQuiverAdjList struct {
	adj_list []QuiverEdgeDesc
}

type QuiverEdgeDesc struct {
	src qse.QuiverIndex
	dst qse.QuiverIndex
}

// Generates a G(n, p) Erdős-Rényi Graph.
func ErdosRenyiGraph(n int, p float64) (al SimpleQuiverAdjList)

// Generates r G(n, p) Erdős-Rényi Graphs, and merges them to form
// a quiver.
func ErdosRenyiQuiver(n int, p float64, r int) (al SimpleQuiverAdjList) {
	adj_list := make([]QuiverEdgeDesc, 0)
	for i := 0; i < r; i++ {
		adj_list = append(
			adj_list,
			ErdosRenyiGraph(n, p).adj_list...,
		)
	}
	al = SimpleQuiverAdjList{adj_list}
	return
}
