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

const UPRGM_QUIVER_DOT_PENWIDTH_FACTOR = 0.1

func MicroprogramQuiverDot(uprgm s.Microprogram) (g *dot.Graph) {
	fmt.Printf(".StateTop  = %#+v\n", uprgm.StateTop)
	fmt.Printf(".StateFail = %#+v\n", uprgm.StateFail)
	g = dot.NewGraph(dot.Directed)
	edges := make(map[GenEdge[s.MicroprogramState]]int)
	nodes := make(map[s.MicroprogramState]dot.Node)
	for k, v_arr := range uprgm.Transitions {
		for _, v := range v_arr {
			if v.Constraints == nil {
				panic("GOT NIL CONSTRAINT!!!!")
			}
			edge_key := GenEdge[s.MicroprogramState]{k, v.StateDst}
			if _, ok := edges[edge_key]; !ok {
				edges[edge_key] = 0
			}
			edges[edge_key] += 1
			var state s.MicroprogramState
			if _, ok := nodes[edge_key.Src]; !ok {
				state = edge_key.Src
				node_text := fmt.Sprintf("%#+v", state)
				if state == uprgm.StateTop {
					node_text = "TOP"
				}
				nodes[edge_key.Src] = g.Node(node_text)
			}
			if _, ok := nodes[edge_key.Dst]; !ok {
				state = edge_key.Dst
				node_text := fmt.Sprintf("%#+v", state)
				if state == uprgm.StateFail {
					node_text = "FAIL"
				}
				nodes[edge_key.Dst] = g.Node(node_text)
			}
		}
	}
	for edge, value := range edges {
		edge := g.Edge(
			nodes[edge.Src],
			nodes[edge.Dst],
		)
		edge.Attr("penwidth", 1+(UPRGM_QUIVER_DOT_PENWIDTH_FACTOR*float64(value)))
	}
	return
}
