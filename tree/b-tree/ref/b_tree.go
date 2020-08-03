package ref

type Tree struct {
	M         int
	Root      *Node
	Underflow int
}

func (tree *Tree) SetM(m int) {
	if m < 3 {
		tree.M = 3
	} else {
		tree.M = m
	}
	tree.Underflow = (m - 1) / 2
}

func (tree *Tree) SearchKey(key int, node *Node, deep int) (*Node, int, int, bool) {
	deep += 1
	if node == nil {
		node = tree.Root
	}

	index, _, ok := node.SearchKey(key)
	if ok || node.Children == nil { // find key or node is leaf
		return node, index, deep, ok
	}
	return tree.SearchKey(key, node.Children[index], deep)
}

func (tree *Tree) AddKey(key int) bool {
	if tree.Root == nil {
		tree.Root = &Node{}
		tree.Root.SetKeys(key)
		return true
	}

	if node, index, _, ok := tree.SearchKey(key, nil, 0); ok {
		return false // The key already existed. No further action.
	} else {
		node.AddKey(tree, index+1, key, nil)
		return true
	}
}

func (tree *Tree) DeleteKey(key int) bool {
	if tree.Root == nil {
		return false
	}

	if node, index, _, ok := tree.SearchKey(key, nil, 0); ok {
		node.DeleteKey(tree, index)
		return true
	} else {
		return false // The key doesn't exist. No further action.
	}
}

func (tree *Tree) Validate(node *Node, step, deep int) bool {
	step++

	// The Keys[0] stores exact the count of the keys
	if node.Keys[0] != len(node.Keys)-1 {
		return false
	}

	// The count of keys is smaller than M and bigger than ceil(M/2)-1
	if node.Parent == nil { // root node
		if node.Children == nil { // root and leaf
			if node.Keys[0] < 0 {
				return false
			}
		} else {
			if node.Keys[0] <= 0 || node.Keys[0] >= tree.M {
				return false
			}
		}
	} else { // non-root node
		if node.Keys[0] < tree.Underflow || node.Keys[0] >= tree.M {
			return false
		}
	}

	// All leaf nodes have the same deep
	if node.Children == nil { // leaf node
		if deep == -1 {
			deep = step
		} else {
			if step != deep {
				return false
			}
		}
	}

	// The keys in Keys[] are ascending
	for i := 1; i < node.Keys[0]; i++ {
		if node.Keys[i] >= node.Keys[i+1] {
			return false
		}
	}

	if node.Children != nil {
		for i := 0; i < node.Keys[0]+1; i++ {
			// The index field stores the correct value
			if node.Children[i].Index != i {
				return false
			}
			// The parent refers to the correct object
			if node.Children[i].Parent != node {
				return false
			}
		}

		for i := 1; i < node.Keys[0]; i++ {
			if node.Keys[i] <= node.Children[i-1].Keys[len(node.Children[i-1].Keys)-1] ||
				node.Keys[i] >= node.Children[i].Keys[1] {
				return false
			}
		}

		// Traverse all nodes
		for _, child := range node.Children {
			return tree.Validate(child, step, deep)
		}
	}

	step--
	return true
}
