package search_tree

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
}

func (tree *Tree) Delete(node *Node) {
	if node.Left == nil {
		tree.Transplant(node, node.Right)
	} else if node.Right == nil {
		tree.Transplant(node, node.Left)
	} else { // node has both left and right children
		rgtMin := node.Right.Minimum()
		if rgtMin.Parent != node {
			tree.Transplant(rgtMin, rgtMin.Right)
			rgtMin.Right = node.Right
			rgtMin.Right.Parent = rgtMin
		}
		tree.Transplant(node, rgtMin)
		rgtMin.Left = node.Left
		rgtMin.Left.Parent = rgtMin
	}
}

func (tree *Tree) Transplant(from, to *Node) {
	if from.Parent == nil {
		tree.Root = to
	} else if from == from.Parent.Left {
		from.Parent.Left = to
	} else { // if from == from.Parent.Right
		from.Parent.Right = to
	}
	if to != nil {
		to.Parent = from.Parent
	}
}