package red_black_tree

import "testing"

type Leaf struct {
	node  *Node
	depth int
	black int
	red   int
}

type validator struct {
	leaves       []Leaf
	invalidColor *Node
	invalidKey   *Node
	depth        int
	black        int
	red          int
}

func (tree *Tree) validate(t *testing.T) {
	if tree.Root == nil {
		return
	}
	if tree.Root.Color == RED {
		t.Error("Invalid root color")
	}

	v := &validator{}
	tree.Root.traverse(v)

	if v.invalidKey != nil {
		t.Errorf("Invalid key for node %d", v.invalidKey.Key)
	}
	if v.invalidColor != nil {
		t.Errorf("Invalid color for node %d", v.invalidColor.Key)
	}

	blackCount := 0
	for _, leaf := range v.leaves {
		if blackCount == 0 {
			blackCount = leaf.black
		} else {
			if blackCount != leaf.black {
				t.Errorf("Conflict black count %d vs %d", blackCount, leaf.black)
			}
		}
	}
}

func (node *Node) traverse(v *validator) {
	v.enterNode(node)
	defer v.leaveNode(node)

	if node.Parent != nil {
		if node.Parent.Color == RED && node.Color == RED {
			v.invalidColor = node
			return
		}
		if node == node.Parent.Left {
			if node.Key > node.Parent.Key {
				v.invalidKey = node
			}
		} else {
			if node.Key < node.Parent.Key {
				v.invalidKey = node
			}
		}
	}

	if node.Left == nil && node.Right == nil {
		leaf := Leaf{node, v.depth, v.black, v.red}
		v.leaves = append(v.leaves, leaf)
		return
	} else {
		if node.Left != nil {
			node.Left.traverse(v)
		}
		if node.Right != nil {
			node.Right.traverse(v)
		}
	}
}

func (v *validator) enterNode(node *Node) {
	v.depth++
	if node.Color == RED {
		v.red++
	} else {
		v.black++
	}
}

func (v *validator) leaveNode(node *Node) {
	v.depth--
	if node.Color == RED {
		v.red--
	} else {
		v.black--
	}
}
