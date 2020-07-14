package red_black_tree

type Tree struct {
	Root *Node
}

func (tree *Tree) Insert(node *Node) {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		if node.Key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	node.Parent = parent
	if parent == nil {
		tree.Root = node
	} else if node.Key < parent.Key {
		parent.Left = node
	} else { // if node.Key >= parent.Key
		parent.Right = node
	}
	node.Color = RED
	tree.fixupInsert(node)
}

func (tree *Tree) fixupInsert(node *Node) {
	for {
		parent, grandparent, uncle, firstRotate, secondRotate, otherSide, ok := tree.getRelatedMaterial(node)
		if !ok || parent.Color != RED {
			break
		}

		if uncle != nil && uncle.Color == RED {
			parent.Color = BLACK
			uncle.Color = BLACK
			grandparent.Color = RED
			node = grandparent
		} else {
			if otherSide {
				node = parent
				firstRotate(node)
			}

			node.Parent.Color = BLACK
			if node.Parent.Parent != nil {
				node.Parent.Parent.Color = RED
				secondRotate(node.Parent.Parent)
			}
			break
		}
	}
	tree.Root.Color = BLACK
}

func (tree *Tree) Delete(node *Node) {
	var fixupNode *Node
	originalColor := node.Color
	if node.Left == nil {
		fixupNode = node.Right
		tree.transplant(node, node.Right)
	} else if node.Right == nil {
		fixupNode = node.Left
		tree.transplant(node, node.Left)
	} else { // node has both left and right children
		rgtMin := node.Right.Minimum()
		originalColor = rgtMin.Color
		fixupNode = rgtMin.Right
		if rgtMin.Parent == node {
			fixupNode.Parent = rgtMin
		} else {
			tree.transplant(rgtMin, rgtMin.Right)
			rgtMin.Right = node.Right
			rgtMin.Right.Parent = rgtMin
		}
		tree.transplant(node, rgtMin)
		rgtMin.Left = node.Left
		rgtMin.Left.Parent = rgtMin
		rgtMin.Color = node.Color
	}
	if originalColor == BLACK {
		tree.fixupDelete(fixupNode)
	}
}

func (tree *Tree) transplant(from, to *Node) {
	if from.Parent == nil {
		tree.Root = to
	} else if from == from.Parent.Left {
		from.Parent.Left = to
	} else { // if from == from.Parent.Right
		from.Parent.Right = to
	}
	to.Parent = from.Parent
}

func (tree *Tree) fixupDelete(node *Node) {

}

func (tree *Tree) leftRotate(center *Node) {
	moon := center.Right
	center.Right = moon.Left
	if moon.Left != nil {
		moon.Left.Parent = center
	}
	moon.Parent = center.Parent
	if center.Parent == nil {
		tree.Root = moon
	} else if center == center.Parent.Left {
		center.Parent.Left = moon
	} else { // if center == center.Parent.Right
		center.Parent.Right = moon
	}
	moon.Left = center
	center.Parent = moon
}

func (tree *Tree) rightRotate(center *Node) {
	moon := center.Left
	center.Left = moon.Right
	if moon.Right != nil {
		moon.Right.Parent = center
	}
	moon.Parent = center.Parent
	if center.Parent == nil {
		tree.Root = moon
	} else if center == center.Parent.Right {
		center.Parent.Right = moon
	} else { // if center == center.Parent.Right
		center.Parent.Left = moon
	}
	moon.Right = center
	center.Parent = moon
}

func (tree *Tree) getRelatedMaterial(node *Node) (parent, grandparent, uncle *Node, firstRotate, secondRotate func(*Node), otherSide, ok bool) {
	parent = node.Parent
	if parent == nil {
		return
	}
	grandparent = parent.Parent
	if grandparent == nil {
		return
	}

	if parent == grandparent.Left {
		uncle = grandparent.Right
		firstRotate = tree.leftRotate
		secondRotate = tree.rightRotate
		if node == parent.Right {
			otherSide = true
		}
	} else {
		uncle = grandparent.Left
		firstRotate = tree.rightRotate
		secondRotate = tree.leftRotate
		if node == parent.Left {
			otherSide = true
		}
	}

	ok = true
	return
}
