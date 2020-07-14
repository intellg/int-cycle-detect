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
	tree.fixup(node)
}

func (tree *Tree) fixup(node *Node) {
	for {
		parent, grandparent, ok := node.getParentGrandparent()
		if !ok || parent.Color != RED {
			break
		}
		if parent == grandparent.Left {
			uncle := grandparent.Right
			if uncle != nil && uncle.Color == RED {
				parent.Color = BLACK
				uncle.Color = BLACK
				grandparent.Color = RED
				node = grandparent
			} else {
				if node == parent.Right {
					node = parent
					tree.leftRotate(node)
				}

				parent, grandparent, ok := node.getParentGrandparent()
				parent.Color = BLACK // parent is never not nil
				if !ok {
					break
				}
				grandparent.Color = RED
				tree.rightRotate(grandparent)
			}
		} else { // if parent == grandparent.Right
			uncle := grandparent.Left
			if uncle != nil && uncle.Color == RED {
				parent.Color = BLACK
				uncle.Color = BLACK
				grandparent.Color = RED
				node = grandparent
			} else {
				if node == parent.Left {
					node = parent
					tree.rightRotate(node)
				}

				parent, grandparent, ok := node.getParentGrandparent()
				parent.Color = BLACK // parent is never not nil
				if !ok {
					break
				}
				grandparent.Color = RED
				tree.leftRotate(grandparent)
			}
		}
	}
	tree.Root.Color = BLACK
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
