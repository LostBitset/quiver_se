package libsynthetic

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

func (q *SimpleQuiverAdjList) CoerceAcyclic(
	src_whitelist map[int]struct{},
	dst_whitelist map[int]struct{},
) {
	offset := 0
	for i_orig := range q.adj_list {
		i := i_orig - offset
		curr := q.adj_list[i]
		if curr.dst > curr.src {
			if _, ok := src_whitelist[curr.src]; !ok {
				if _, ok := dst_whitelist[curr.dst]; !ok {
					q.adj_list[i] = SimpleEdgeDesc{
						src: curr.dst,
						dst: curr.src,
					}
				}
			}
		}
		if curr.dst == curr.src {
			qse.SpliceOutReclaim(&q.adj_list, i)
			offset += 1
		}
	}
}
