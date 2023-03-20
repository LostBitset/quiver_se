package main

import (
	qse "LostBitset/quiver_se/lib"
	"fmt"
)

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
	must_destroy_root := tree.CleanUpNonNegativeSubtrees()
	if must_destroy_root {
		if n_leaves == 0 {
			panic("Cannot generate a tree with no leaves. ")
		} else {
			panic("Unreachable. Was forced to generate a tree with no leaves. ")
		}
	}
	if tree.ComputeLeafCount() != n_leaves {
		fmt.Printf(
			"Incorrrectly generated a tree with %d leaves. Should have had %d leaves.\n",
			tree.ComputeLeafCount(),
			n_leaves,
		)
		panic("Unreachable. Should have generated a tree with n_leaves leaves, but did not.")
	}
	return
}

func (tree *SimpleTree) CleanUpNonNegativeSubtrees() (destroy bool) {
	destroy = false
	if len(tree.children) == 0 {
		if tree.id < 0 {
			// Negative leaf case
			return
		} else {
			destroy = true
			// Nonnegative leaf case
			return
		}
	}
	// Negative case
	children_to_destroy := make([]int, 0)
	destroy_all_children := true
	for i, child := range tree.children {
		destroy_child := child.CleanUpNonNegativeSubtrees()
		if destroy_child {
			children_to_destroy = append(children_to_destroy, i)
		} else {
			destroy_all_children = false
		}
	}
	if destroy_all_children {
		destroy = true
	} else {
		offset := 0
		for _, child_index := range children_to_destroy {
			child_index := child_index - offset
			qse.SpliceOutReclaim(&tree.children, child_index)
			offset++
		}
	}
	return
}
