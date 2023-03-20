package main

func BakePruferSequence(sequence []int) (ps PruferSequence) {
	ps.sequence = sequence
	n := len(sequence)
	degrees := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		degrees[i] = 1
	}
	for _, item := range sequence {
		degrees[item] += 1
	}
	ps.degrees = degrees
	return
}

func (ps PruferSequence) ToTree() (tree SimpleTree) {
	degrees := make([]int, len(ps.degrees))
	copy(degrees, ps.degrees)
	al := make(map[int][]int) // Adjacency List
	for node := range degrees {
		al[node] = make([]int, 0)
	}
	// Actually build up the tree graph
	for _, item := range ps.sequence {
		src_node := item + 1
	findDestinationNodeLoop:
		for dst_node, degree := range degrees {
			if degree == 1 {
				al[src_node] = append(al[src_node], dst_node)
				degrees[src_node] -= 1
				degrees[dst_node] -= 1
				break findDestinationNodeLoop
			}
		}
	}
	// Add the final edge
	final_src, final_dst := -1, -1
addFinalNodeLoop:
	for node := range degrees {
		if degrees[node] == 0 {
			final_src = node
		} else {
			final_dst = node
			break addFinalNodeLoop
		}
	}
	al[final_src] = append(al[final_src], final_dst)
	// Convert to a tree
	// We know that the indegree is always one
	// Except for node 0
	tree = SimpleTreeFromAdjList(al, len(degrees))
	return
}

func SimpleTreeFromAdjList(al map[int][]int, n int) (tree SimpleTree) {
	return
}
