package libsynthetic

func (tree *SimpleTree) MaxDegree() (md int) {
	md = len(tree.children)
	for _, child := range tree.children {
		child_md := child.MaxDegree()
		if child_md > md {
			md = child_md
		}
	}
	return
}
