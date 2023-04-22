package main

import (
	"fmt"

	s "github.com/LostBitset/quiver_se/lib_synthetic"
	"github.com/emicklei/dot"
)

type GenEdge[A comparable] struct {
	Src A
	Dst A
}

const UPRGM_QUIVER_DOT_PENWIDTH_FACTOR = 2.0

func MicroprogramQuiverDot(uprgm s.Microprogram) (g *dot.Graph) {
	g = dot.NewGraph(dot.Directed)
	edges := make(map[GenEdge[s.MicroprogramState]]int)
	nodes := make(map[s.MicroprogramState]dot.Node)
	for k, v_arr := range uprgm.Transitions {
		for _, v := range v_arr {
			edge_key := GenEdge[s.MicroprogramState]{k, v.StateDst}
			if _, ok := edges[edge_key]; !ok {
				edges[edge_key] = 0
			}
			edges[edge_key] += 1
			if _, ok := nodes[edge_key.Src]; !ok {
				nodes[edge_key.Src] = g.Node(
					fmt.Sprintf("%#+v", edge_key.Src),
				)
			}
			if _, ok := nodes[edge_key.Dst]; !ok {
				nodes[edge_key.Dst] = g.Node(
					fmt.Sprintf("%#+v", edge_key.Dst),
				)
			}
		}
	}
	for edge, value := range edges {
		g.Edge(
			nodes[edge.Src],
			nodes[edge.Dst],
		).Attr("penwidth", 1+0*UPRGM_QUIVER_DOT_PENWIDTH_FACTOR*value)
	}
	return
}
