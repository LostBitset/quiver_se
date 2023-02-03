package qse

func StartSiMReQ[
	QNODE any,
	ATOM comparable,
	IDENT any,
	SORT any,
	MODEL any,
	SCTX SMTSolvedContext[MODEL],
	SYS SMTSystem[
		IdLiteral[ATOM],
		IDENT,
		SORT,
		MODEL,
		SCTX,
	],
](
	in_updates chan Augmented[
		QuiverUpdate[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}], *DMT[WithId_H[ATOM], QuiverIndex]],
		[]SMTFreeFun[IDENT, SORT],
	],
	out_models chan MODEL,
	sys SYS,
) (
	top_node QuiverIndex,
) {
	walks := make(chan Augmented[
		QuiverWalk[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}]],
		[]SMTFreeFun[IDENT, SORT],
	])
	canidates := make(chan SMTQueryDNFClause[ATOM, IDENT, SORT])
	smr_config := NewSMRConfig[ATOM, IDENT, SORT, MODEL, SCTX](
		canidates, out_models, sys,
	)
	var dmtq Quiver[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}], *DMT[WithId_H[ATOM], QuiverIndex]]
	top_node_dmt := NewDMT[WithId_H[ATOM], QuiverIndex]()
	var zero_node QNODE
	top_node = dmtq.InsertNode(zero_node, &top_node_dmt)
	warden_config := DMTQWardenConfig[QNODE, WithId_H[ATOM], []SMTFreeFun[IDENT, SORT]]{
		in_updates: in_updates,
		out_walks:  walks,
		walk_src:   top_node,
		dmtq:       dmtq,
	}
	smr_config.Start()
	go func() {
		defer close(canidates)
		for walk_recv := range walks {
			walk_chunked := walk_recv.value
			walk := make([]IdLiteral[ATOM], 0)
			for _, chunk := range walk_chunked.edges_chunked {
				for _, set := range *chunk {
					stdlib_set := set.ToStdlibMap()
					for key := range stdlib_set {
						walk = append(walk, IdLiteral[ATOM](key))
					}
				}
			}
			canidates <- SMTQueryDNFClause[ATOM, IDENT, SORT]{
				walk,
				walk_recv.augment,
			}
		}
	}()
	warden_config.Start()
	return
}
