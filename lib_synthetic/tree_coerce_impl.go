package libsynthetic

func (tree *SimpleTree) CoerceToDegree(n int) {
	if len(tree.children) == 0 {
		return
	}
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
	}
	for _, child := range tree.children {
		child.CoerceToDegree(n)
	}
}

func (tree *SimpleTree) CoerceForbidDegreeOne() {
	if len(tree.children) == 1 {
		children_source := tree.children[0].children
		new_children := make([]*SimpleTree, len(children_source))
		for i, child_ref := range children_source {
			new_children[i] = child_ref
			(*child_ref).ccount += 1
		}
		tree.children = new_children
	}
	for _, child := range tree.children {
		child.CoerceForbidDegreeOne()
	}
}
