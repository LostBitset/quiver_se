package qse

import "fmt"

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
	dmtq Quiver[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}], *DMT[WithId_H[ATOM], QuiverIndex]],
	top_node QuiverIndex,
	fail_node QuiverIndex,
) {
	walks := make(chan Augmented[
		QuiverWalk[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}]],
		[]SMTFreeFun[IDENT, SORT],
	])
	canidates := make(chan SMRDNFClause[ATOM, IDENT, SORT])
	smr_config := NewSMRConfig[ATOM, IDENT, SORT, MODEL, SCTX](
		canidates, out_models, sys,
	)
	top_node_dmt := NewDMT[WithId_H[ATOM], QuiverIndex]()
	var zero_node QNODE
	top_node = dmtq.InsertNode(zero_node, &top_node_dmt)
	warden_config := DMTQWardenConfig[QNODE, WithId_H[ATOM], []SMTFreeFun[IDENT, SORT]]{
		in_updates: in_updates,
		out_walks:  walks,
		walk_src:   top_node,
		walk_dst:   fail_node,
		dmtq:       dmtq,
	}
	smr_config.Start()
	go func() {
		defer close(canidates)
		processed_hashes := make(map[uint32]struct{})
		for walk_recv := range walks {
			sum := uint32(0xE4E4)
			walk_chunked := walk_recv.value
			chunks := walk_chunked.edges_chunked
			if len(*chunks[len(chunks)-1]) == 0 {
				continue
			}
			processed_ids := make(map[NumericId]struct{})
			walk := make([]IdLiteral[ATOM], 0)
			for _, chunk := range chunks {
				for _, set := range *chunk {
					stdlib_set := set.ToStdlibMap()
					for key := range stdlib_set {
						if _, ok := processed_ids[key.value.id]; !ok {
							walk = append(walk, IdLiteral[ATOM](key))
							processed_ids[key.value.id] = struct{}{}
							sum ^= FixDigest32(key.value.id, 0x4E)
						}
					}
				}
			}
			fmt.Println()
			if _, ok := processed_hashes[sum]; ok {
				continue
			}
			processed_hashes[sum] = struct{}{}
			canidates <- SMRDNFClause[ATOM, IDENT, SORT]{
				walk[:len(walk)],
				walk[:len(walk)], // this bit is very much a TODO
				walk_recv.augment,
			}
		}
	}()
	warden_config.Start()
	return
}
