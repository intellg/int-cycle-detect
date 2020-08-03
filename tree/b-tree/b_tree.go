package b_tree

type Tree struct {
	Degree    int
	Root      *Node
	Underflow int
}

func (tree *Tree) Init(degree int) {
	tree.Degree = degree
	root := &Node{
		IsLeaf: true,
		Keys:   make([]int, 0),
	}
	tree.Root = root
}

func (tree *Tree) Insert(key int) {
	root := tree.Root
	if len(root.Keys) == 2*tree.Degree-1 {
		node := &Node{
			Keys:     make([]int, 0),
			Children: make([]*Node, 1),
		}
		tree.Root = node

		node.Children[0] = root
		node.splitChild(0)
		node.insert(tree.Degree, key)
	} else {
		root.insert(tree.Degree, key)
	}
}

func (tree *Tree) Delete(key int) {
	root := tree.Root
	if !root.IsLeaf && len(root.Keys) == 1 &&
		len(root.Children[0].Keys) == tree.Degree-1 &&
		len(root.Children[1].Keys) == tree.Degree-1 {
		root.mergeChild(0)
		tree.Root = root.Children[0]
		root = tree.Root
	}
	root.delete(tree.Degree, key)
}

func (tree *Tree) Validate() bool {
	root := tree.Root
	if len(root.Keys) == 0 {
		return true
	}
	ok, depthArr := root.Validate(tree.Degree, 0, make([]int, 0))
	if !ok {
		return false
	}

	depth := depthArr[0]
	for i := 1; i < len(depthArr); i++ {
		if depth != depthArr[i] {
			return false
		}

	}
	return true
}
