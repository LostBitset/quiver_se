package main

func (tree SimpleTree) CoerceToMaxDegree(n int) {
	if len(tree.children) > n {
		// Put all of the other children in their own tree
		// TODO
	}
	for _, child := range tree.children {
		child.CoerceToMaxDegree(n)
	}
}
