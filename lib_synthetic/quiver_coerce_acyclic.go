package libsynthetic

func (q *SimpleQuiverAdjList) CoerceAcyclic() {
	for i := range q.adj_list {
		curr := q.adj_list[i]
		if curr.dst > curr.src {
			q.adj_list[i] = SimpleEdgeDesc{
				src: curr.dst,
				dst: curr.src,
			}
		}
	}
}
