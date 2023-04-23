package libsynthetic

func (tree *SimpleTree) CoerceToDegree(n int) {
	if len(tree.children) > n {
		// Put all of the other children in their own tree
		set_aside_len := n - 1
		children_set_aside := make([]*SimpleTree, len(tree.children)-set_aside_len)
		copy(children_set_aside, tree.children)
		var zero int
		new_children := make([]*SimpleTree, 0)
		for _, child_ref := range children_set_aside {
			new_children = append(new_children, child_ref)
		}
		new_children = append(
			new_children,
			// A new node to hold the other children
			Pto(SimpleTree{
				id:       zero,
				children: children_set_aside,
			}),
		)
	} else if len(tree.children) < n {
		// TODO TODO TODO
		// somehow? coerce to smaller degree
		// by like absorbing children
		// this isnt a horror story sorry about the last two words
	}
	for _, child := range tree.children {
		child.CoerceToDegree(n)
	}
}
