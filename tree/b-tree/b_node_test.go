package b_tree

import (
	"testing"
)

func TestSplitChild1(t *testing.T) {
	root := &Node{}
	root.Keys = []int{10, 20}
	root.Children = make([]*Node, 3)
	root.IsLeaf = false

	node123 := &Node{}
	node123.Keys = []int{1, 2, 3}
	node123.IsLeaf = false
	c1 := &Node{Keys: []int{10}}
	c2 := &Node{Keys: []int{15}}
	c3 := &Node{Keys: []int{25}}
	c4 := &Node{Keys: []int{35}}
	node123.Children = []*Node{c1, c2, c3, c4}
	root.Children[0] = node123

	root.splitChild(0)
	root.OutputJson("testSplitChild1.json")
}

func TestSplitChild2(t *testing.T) {
	root := &Node{}
	root.Keys = []int{10, 20}
	root.Children = make([]*Node, 3)
	root.IsLeaf = false

	node111213 := &Node{}
	node111213.Keys = []int{11, 12, 13}
	node111213.IsLeaf = false
	c11 := &Node{Keys: []int{11}}
	c12 := &Node{Keys: []int{12}}
	c13 := &Node{Keys: []int{13}}
	c14 := &Node{Keys: []int{14}}
	node111213.Children = []*Node{c11, c12, c13, c14}
	root.Children[1] = node111213

	root.splitChild(1)
	root.OutputJson("testSplitChild2.json")
}

func TestBorrow(t *testing.T) {
	n1 := &Node{
		Keys:   []int{1},
		IsLeaf: true,
	}
	n3 := &Node{
		Keys:   []int{3},
		IsLeaf: true,
	}
	n5 := &Node{
		Keys:   []int{5},
		IsLeaf: true,
	}
	n7 := &Node{
		Keys:   []int{7},
		IsLeaf: true,
	}
	n9 := &Node{
		Keys:   []int{9},
		IsLeaf: true,
	}
	left := &Node{
		Keys:     []int{2, 4},
		Children: []*Node{n1, n3, n5},
	}
	right := &Node{
		Keys:     []int{8},
		Children: []*Node{n7, n9},
	}
	root := &Node{
		Keys:     []int{6},
		Children: []*Node{left, right},
	}

	root.OutputJson("beforeBorrow.json")
	root.borrowLeft(1)
	root.OutputJson("borrowLeft.json")
	root.borrowRight(0)
	root.OutputJson("borrowRight.json")
}
