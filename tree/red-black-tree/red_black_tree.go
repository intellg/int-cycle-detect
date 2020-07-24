package red_black_tree

type Tree struct {
	Root *Node
}

var (
	nilNode = &Node{Key: -1, Color: BLACK}
)

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
	} else { // node.Key >= parent.Key
		parent.Right = node
	}
	node.Color = RED
	tree.fixupInsert(node)
}

func (tree *Tree) fixupInsert(node *Node) {
	for {
		parent, grandparent, uncle, parentRotate, grandparentRotate, ok := tree.getInsertMaterial(node)
		if !ok || parent.Color != RED {
			break
		}

		if uncle != nil && uncle.Color == RED {
			parent.Color = BLACK
			uncle.Color = BLACK
			grandparent.Color = RED
			node = grandparent
		} else {
			if parentRotate != nil { // node is different side as its parent
				node = parent
				parentRotate(node)
			}

			node.Parent.Color = BLACK
			grandparent := node.Parent.Parent
			if grandparent != nil {
				grandparent.Color = RED
				grandparentRotate(grandparent)
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
		if fixupNode == nil {
			fixupNode = nilNode
			fixupNode.Parent = node
			node.Right = fixupNode
		}
		tree.transplant(node, node.Right)
	} else if node.Right == nil {
		fixupNode = node.Left
		if fixupNode == nil {
			fixupNode = nilNode
			fixupNode.Parent = node
			node.Left = fixupNode
		}
		tree.transplant(node, node.Left)
	} else { // node has both left and right children
		replaceNode := node.Right.Minimum()
		fixupNode = replaceNode.Right
		originalColor = replaceNode.Color
		if fixupNode == nil {
			fixupNode = nilNode
			fixupNode.Parent = replaceNode
			replaceNode.Right = fixupNode
		}
		if replaceNode.Parent != node {
			tree.transplant(replaceNode, replaceNode.Right)
			replaceNode.Right = node.Right
			replaceNode.Right.Parent = replaceNode
		}
		tree.transplant(node, replaceNode)
		replaceNode.Left = node.Left
		replaceNode.Left.Parent = replaceNode
		replaceNode.Color = node.Color
	}
	if originalColor == BLACK {
		tree.fixupDelete(fixupNode)
	}
	if fixupNode == nilNode && fixupNode.Parent != nil {
		if fixupNode == fixupNode.Parent.Left {
			fixupNode.Parent.Left = nil
		} else { // fixupNode == fixupNode.Parent.Right
			fixupNode.Parent.Right = nil
		}
	}
}

func (tree *Tree) transplant(from, to *Node) {
	if from.Parent == nil {
		tree.Root = to
	} else if from == from.Parent.Left {
		from.Parent.Left = to
	} else { // from == from.Parent.Right
		from.Parent.Right = to
	}
	if to != nil {
		to.Parent = from.Parent
	}
}

func (tree *Tree) fixupDelete(node *Node) {
	for node != tree.Root && node.Color == BLACK {
		parent := node.Parent
		var isNodeLeft bool
		if node == parent.Left {
			isNodeLeft = true
		}
		sibling := tree.getSibling(parent, isNodeLeft)
		if sibling.Color == RED {
			sibling.Color = BLACK
			parent.Color = RED
			tree.towardRotate(isNodeLeft)(parent)
			sibling = tree.getSibling(parent, isNodeLeft) // original near nephew
		}
		nearNephew := tree.getNearNephew(sibling, isNodeLeft)
		farNephew := tree.getFarNephew(sibling, isNodeLeft)
		if (nearNephew == nil || nearNephew.Color == BLACK) && (farNephew == nil || farNephew.Color == BLACK) {
			sibling.Color = RED
			node = parent
		} else {
			if farNephew == nil || farNephew.Color == BLACK {
				nearNephew.Color = BLACK
				sibling.Color = RED
				tree.forwardRotate(isNodeLeft)(sibling)
				sibling = tree.getSibling(parent, isNodeLeft)
			}
			sibling.Color = parent.Color
			parent.Color = BLACK
			farNephew = tree.getFarNephew(sibling, isNodeLeft)
			if farNephew != nil {
				farNephew.Color = BLACK
			}
			tree.towardRotate(isNodeLeft)(parent)
			break
		}
	}
	node.Color = BLACK
}

func (tree *Tree) towardRotate(isNodeLeft bool) func(*Node) {
	if isNodeLeft {
		return tree.leftRotate
	} else {
		return tree.rightRotate
	}
}

func (tree *Tree) forwardRotate(isNodeLeft bool) func(*Node) {
	if isNodeLeft {
		return tree.rightRotate
	} else {
		return tree.leftRotate
	}
}

func (tree *Tree) getSibling(parent *Node, isNodeLeft bool) *Node {
	if isNodeLeft {
		return parent.Right
	} else {
		return parent.Left
	}
}

func (tree *Tree) getNearNephew(sibling *Node, isNodeLeft bool) *Node {
	if isNodeLeft {
		return sibling.Left
	} else {
		return sibling.Right
	}
}

func (tree *Tree) getFarNephew(sibling *Node, isNodeLeft bool) *Node {
	if isNodeLeft {
		return sibling.Right
	} else {
		return sibling.Left
	}
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
	} else { // center == center.Parent.Right
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
	} else { // center == center.Parent.Right
		center.Parent.Left = moon
	}
	moon.Right = center
	center.Parent = moon
}

func (tree *Tree) getInsertMaterial(node *Node) (parent, grandparent, uncle *Node, parentRotate, grandparentRotate func(*Node), ok bool) {
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
		if node == parent.Right {
			parentRotate = tree.leftRotate
		}
		grandparentRotate = tree.rightRotate
	} else { // parent == grandparent.Right
		uncle = grandparent.Left
		if node == parent.Left {
			parentRotate = tree.rightRotate
		}
		grandparentRotate = tree.leftRotate
	}

	ok = true
	return
}
