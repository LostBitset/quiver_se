package libsynthetic

import "fmt"

func (tree *SimpleTree) CoerceToMaxDegree(n int) {
	if len(tree.children) == 0 {
		fmt.Printf("(in CoerceMaxDegree) <subtree which is leaf>.id = %#+v\n", tree.id)
		return
	}
	if len(tree.children) > n {
		// Put all of the other children in their own tree
		// set aside = moved into new subtree
		set_aside_len := len(tree.children) - (n - 1)
		children_set_aside := make([]*SimpleTree, set_aside_len)
		copy(children_set_aside, tree.children)
		new_children := make([]*SimpleTree, 0)
		for i := set_aside_len; i < len(tree.children); i++ {
			new_children = append(new_children, tree.children[i])
		}
		new_children = append(
			new_children,
			// A new node to hold the other children
			Pto(SimpleTree{
				id:       0,
				children: children_set_aside,
			}),
		)
		tree.children = new_children
	}
	for _, child := range tree.children {
		child.CoerceToMaxDegree(n)
	}
}

func (tree *SimpleTree) CoerceForbidDegreeOne() {
	for len(tree.children) == 1 {
		children_source := tree.children[0].children
		new_children := make([]*SimpleTree, len(children_source))
		for i, child_ref := range children_source {
			new_children[i] = child_ref
			// (*child_ref).ccount += 1
		}
		tree.children = new_children
	}
	for _, child := range tree.children {
		child.CoerceForbidDegreeOne()
	}
}
