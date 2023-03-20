package main

import "fmt"

func (tree *SimpleTree) ComputeLeafReferences() (leaf_refs []*SimpleTree) {
	if len(tree.children) == 0 {
		backing_current := tree
		leaf_refs = []*SimpleTree{backing_current}
	} else {
		leaf_refs = make([]*SimpleTree, 0)
		for _, child := range tree.children {
			leaf_refs = append(leaf_refs, child.ComputeLeafReferences()...)
		}
	}
	return
}

func (tree SimpleTree) ComputeLeafCount() (n_leaves int) {
	n_leaves = len(tree.ComputeLeafReferences())
	return
}

func PruferEvenFinalRandomTree(n_nonleaf int, n_leaves int) (tree SimpleTree) {
	tree = PruferRandomTree(n_nonleaf - 2)
	leaf_refs := tree.ComputeLeafReferences()
	period := len(leaf_refs)
	base_former_leaf_degree := n_leaves / period
	n_addl_leaf_degree := n_leaves % period
	for i, leaf_ref := range leaf_refs {
		leaf_degree := base_former_leaf_degree
		if i < n_addl_leaf_degree {
			leaf_degree += 1
		}
		actual_leaf_values := make([]int, leaf_degree)
		for j := 0; j < leaf_degree; j++ {
			actual_leaf_values[j] = -(j + 1)
		}
		actual_leaves := make([]*SimpleTree, leaf_degree)
		for j, value := range actual_leaf_values {
			backing_leaf := SimpleTree{
				id:       value,
				children: []*SimpleTree{},
			}
			actual_leaves[j] = &backing_leaf
		}
		(*leaf_ref).children = append((*leaf_ref).children, actual_leaves...)
	}
	if tree.ComputeSize() != (n_nonleaf + n_leaves) {
		fmt.Printf(
			"Incorrectly generated a tree of size %d. Should have been %d.\n",
			tree.ComputeSize(),
			n_leaves+n_nonleaf,
		)
		panic("Unreachable. Should have generated a tree of a different size.")
	}
	if tree.ComputeLeafCount() != n_leaves {
		fmt.Printf("Incorrrectly generated a tree with %d leaves.\n", tree.ComputeLeafCount())
		panic("Unreachable. Should have generated a tree with n_leaves leaves, but did not.")
	}
	return
}
