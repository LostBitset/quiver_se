package libsynthetic

func (tree *SimpleTree) CoerceToMaxDegree(n int) {
	if len(tree.children) > n {
		// Put all of the other children in their own tree
		children_set_aside := make([]*SimpleTree, len(tree.children)-1)
		copy(children_set_aside, tree.children)
		var zero int
		tree.children = []*SimpleTree{
			tree.children[len(tree.children)-1],
			// A new node to hold the other children
			{
				id:       zero,
				children: children_set_aside,
			},
		}
	}
	for _, child := range tree.children {
		child.CoerceToMaxDegree(n)
	}
}
