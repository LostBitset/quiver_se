package qse

import (
	log "github.com/sirupsen/logrus"
)

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
	aot_nodes []QNODE,
) (
	dmtq *Quiver[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}], *DMT[WithId_H[ATOM], QuiverIndex]],
	top_node QuiverIndex,
	fail_node QuiverIndex,
	aot_indices []QuiverIndex,
) {
	var dmtq_backing Quiver[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}], *DMT[WithId_H[ATOM], QuiverIndex]]
	dmtq = &dmtq_backing
	walks := make(chan Augmented[
		QuiverWalk[QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}]],
		[]SMTFreeFun[IDENT, SORT],
	])
	canidates := make(chan SMRDNFClause[ATOM, IDENT, SORT])
	smr_config := NewSMRConfig[ATOM, IDENT, SORT, MODEL, SCTX](
		canidates, out_models, sys,
	)
	top_node_dmt := NewDMT[WithId_H[ATOM], QuiverIndex]()
	fail_node_dmt := NewDMT[WithId_H[ATOM], QuiverIndex]()
	var zero_node QNODE
	top_node = dmtq.InsertNode(zero_node, &top_node_dmt)
	fail_node = dmtq.InsertNode(zero_node, &fail_node_dmt)
	if aot_nodes != nil {
		aot_indices = make([]QuiverIndex, len(aot_nodes))
		for i, node := range aot_nodes {
			aot_backing_dmt := NewDMT[WithId_H[ATOM], QuiverIndex]()
			aot_indices[i] = dmtq.InsertNode(node, &aot_backing_dmt)
		}
	}
	warden_config := DMTQWardenConfig[QNODE, WithId_H[ATOM], []SMTFreeFun[IDENT, SORT]]{
		in_updates: in_updates,
		out_walks:  walks,
		walk_src:   top_node,
		walk_dst:   fail_node,
		dmtq:       dmtq,
	}
	warden_config.Start()
	go func(canidates chan SMRDNFClause[ATOM, IDENT, SORT]) {
		defer close(canidates)
		processed_hashes := make(map[uint32]struct{})
		for walk_recv := range walks {
			log.Info("[simreq_part/go1(junction)] Received (augmented) quiver walk. ")
			sum := uint32(0xE4E4)
			walk_chunked := walk_recv.Value
			chunks := walk_chunked.edges_chunked
			/*
				// There's a chance this might be important, but right now it just breaks stuff.
				if len(*chunks[len(chunks)-1]) == 0 {
					log.Info("[simreq_part/go1(junction)] Empty last chunk, ignoring. ")
					continue
				}
			*/
			processed_ids := make(map[NumericId]struct{})
			walk := make([]IdLiteral[ATOM], 0)
			boundaries := make([]int, 0)
			for _, chunk := range chunks {
				for _, set := range *chunk {
					stdlib_set := set.ToStdlibMap()
					for key := range stdlib_set {
						if _, ok := processed_ids[key.Value.Id]; !ok {
							walk = append(walk, IdLiteral[ATOM](key))
							processed_ids[key.Value.Id] = struct{}{}
							sum ^= FixDigest32(key.Value.Id, 0x4E)
						}
					}
				}
				var old_boundary int
				if len(boundaries) > 0 {
					old_boundary = boundaries[len(boundaries)-1]
				} else {
					old_boundary = 0
				}
				new_boundary := len(walk)
				if old_boundary != new_boundary {
					boundaries = append(boundaries, new_boundary)
				}
			}
			if len(boundaries) > 0 && boundaries[len(boundaries)-1] == len(walk) {
				SpliceOutReclaim(&boundaries, len(boundaries)-1)
			}
			var failure_boundary int
			if len(boundaries) > 0 {
				failure_boundary = boundaries[len(boundaries)-1]
			} else {
				failure_boundary = 0
			}
			if _, ok := processed_hashes[sum]; ok {
				log.Info("[simreq_part/go1(junction)] Seen hash already, skipping. ")
				continue
			}
			processed_hashes[sum] = struct{}{}
			canidates <- SMRDNFClause[ATOM, IDENT, SORT]{
				walk[:failure_boundary],
				walk[failure_boundary:],
				walk_recv.Augment,
			}
			log.Info("[simreq_part/go1(junction)] Sent (augmented) quiver walk in canidate form. ")
		}
	}(canidates)
	smr_config.Start()
	return
}
