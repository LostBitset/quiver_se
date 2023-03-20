package main

import (
	"math"
	"math/rand"
)

func BakePruferSequence(sequence []int) (ps PruferSequence) {
	ps.sequence = sequence
	n := len(sequence)
	degrees := make([]int, n+2)
	for i := 0; i < len(degrees); i++ {
		degrees[i] = 1
	}
	for _, item := range sequence {
		degrees[item-1] += 1
	}
	ps.degrees = degrees
	return
}

func (ps PruferSequence) ToTree() (tree SimpleTree) {
	degrees := make([]int, len(ps.degrees))
	copy(degrees, ps.degrees)
	observed_indegrees := make(map[int]int)
	for node := range degrees {
		observed_indegrees[node] = 0
	}
	al := make(map[int][]int) // Adjacency List
	for node := range degrees {
		al[node] = make([]int, 0)
	}
	// Actually build up the tree graph
	for _, item := range ps.sequence {
		src_node := item - 1
	findDestinationNodeLoop:
		for dst_node, degree := range degrees {
			if degree == 1 {
				al[src_node] = append(al[src_node], dst_node)
				observed_indegrees[dst_node] += 1
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
	observed_indegrees[final_dst] += 1
	// Find the root node
	// We know that the indegree is always one
	// Except for the root node with indegree zero
	root_node := -1
findRootNodeByIndegreeLoop:
	for node := range degrees {
		if observed_indegrees[node] == 0 {
			root_node = node
			break findRootNodeByIndegreeLoop
		}
	}
	if root_node == -1 {
		panic("Something went wrong. No root node found for what should have been a tree graph. ")
	}
	tree = SimpleTreeFromAdjList(al, root_node)
	return
}

func SimpleTreeFromAdjList(al map[int][]int, root int) (tree SimpleTree) {
	outneighbors := al[root]
	children := make([]*SimpleTree, len(outneighbors))
	for i, outneighbor := range outneighbors {
		backing_subtree := SimpleTreeFromAdjList(al, outneighbor)
		children[i] = &backing_subtree
	}
	tree = SimpleTree{
		id:       root,
		children: children,
	}
	return
}

func RandomPruferSequence(n int) (ps PruferSequence) {
	// A sequence of length n containing integers in the range [1, n+2]
	sequence := make([]int, n)
	for i := range sequence {
		sequence[i] = RandomPruferSequenceElement(n)
	}
	ps = BakePruferSequence(sequence)
	return
}

func RandomPruferSequenceElement(n int) (elem int) {
	in_unit_interval := rand.Float64()
	in_shifted_range := in_unit_interval * float64(n+2)
	float_value := in_shifted_range + 1
	elem = int(math.Floor(float_value))
	return
}

// Generate a random tree as defined by a random Prüfer Sequence
func PruferRandomTree(n int) (tree SimpleTree) {
	ps := RandomPruferSequence(n)
	tree = ps.ToTree()
	return
}
