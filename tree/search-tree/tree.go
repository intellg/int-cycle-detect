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

func (tree *Tree) Delete(n *Node) {
	if n.Left == nil {
		tree.Transplant(n, n.Right)
	} else if n.Right == nil {
		tree.Transplant(n, n.Left)
	} else { // n has both left and right children
		t := n.Right.Minimum()
		if t.Parent != n {
			tree.Transplant(t, t.Right)
			t.Right = n.Right
			t.Right.Parent = t
		}
		tree.Transplant(n, t)
		t.Left = n.Left
		t.Left.Parent = t
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