package libsynthetic

import "math"

// Generates a G(n, p) Erdős-Rényi Graph.
func ErdosRenyiGraph(n int, p float64) (al SimpleQuiverAdjList) {
	adj_list := make([]SimpleEdgeDesc, 0)
	edge_exists_distr := SimpleDDistr[ConnectedOrNot]{map[ConnectedOrNot]float64{
		Connected:    p,
		NotConnected: 1.0 - p,
	}}
	edge_exists := BakeDDistr[ConnectedOrNot](edge_exists_distr)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Self-loops are allowed
			switch edge_exists.Sample() {
			case Connected:
				adj_list = append(adj_list, SimpleEdgeDesc{i, j})
			}
		}
	}
	al = SimpleQuiverAdjList{adj_list}
	return
}

// Generates r G(n, p) Erdős-Rényi Graphs, and merges them to form
// a quiver.
func ErdosRenyiQuiver(n int, p float64, r int) (al SimpleQuiverAdjList) {
	adj_list := make([]SimpleEdgeDesc, 0)
	for i := 0; i < r; i++ {
		adj_list = append(
			adj_list,
			ErdosRenyiGraph(n, p).adj_list...,
		)
	}
	al = SimpleQuiverAdjList{adj_list}
	return
}

// See ErdosRenyiQuiver.
// Takes in n, p, and the average number of edges
func ErdosRenyiQuiverGivenEdges(n int, p float64, avg_ne int) (al SimpleQuiverAdjList) {
	n_choose_2 := (n * (n + 1)) / 2
	r_float := float64(avg_ne) / (float64(n_choose_2) * p)
	r := int(math.Round(r_float))
	al = ErdosRenyiQuiver(n, p, r)
	return
}

func (sqal *SimpleQuiverAdjList) InsertEdge(src int, dst int) {
	sqal.adj_list = append(sqal.adj_list, SimpleEdgeDesc{src, dst})
	return
}

func (sqal SimpleQuiverAdjList) ExtractAdjListAsMap(n int) (al map[int][]int) {
	al = make(map[int][]int)
	for i := 0; i < n; i++ {
		al[i] = make([]int, 0)
	}
	for _, edge := range sqal.adj_list {
		al[edge.src] = append(al[edge.src], edge.dst)
	}
	return
}
