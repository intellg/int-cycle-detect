package b_tree

type Tree struct {
	M         int
	Root      *Node
	Underflow int
}

func (instance *Tree) SetM(m int) {
	if m < 3 {
		instance.M = 3
	} else {
		instance.M = m
	}
	instance.Underflow = (m - 1) / 2
}

func (instance *Tree) SearchKey(key int, node *Node, deep int) (*Node, int, int, bool) {
	deep += 1
	if node == nil {
		node = instance.Root
	}

	index, _, ok := node.SearchKey(key)
	if ok || node.Children == nil { // find key or node is leaf
		return node, index, deep, ok
	}
	return instance.SearchKey(key, node.Children[index], deep)
}

func (instance *Tree) AddKey(key int) bool {
	if instance.Root == nil {
		instance.Root = &Node{}
		instance.Root.SetKeys(key)
		return true
	}

	if node, index, _, ok := instance.SearchKey(key, nil, 0); ok {
		return false // The key already existed. No further action.
	} else {
		node.AddKey(instance, index+1, key, nil)
		return true
	}
}

func (instance *Tree) DeleteKey(key int) bool {
	if instance.Root == nil {
		return false
	}

	if node, index, _, ok := instance.SearchKey(key, nil, 0); ok {
		node.DeleteKey(instance, index)
		return true
	} else {
		return false // The key doesn't exist. No further action.
	}
}

func (instance *Tree) validate(node *Node, step, deep int) bool {
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
			if node.Keys[0] <= 0 || node.Keys[0] >= instance.M {
				return false
			}
		}
	} else { // non-root node
		if node.Keys[0] < instance.Underflow || node.Keys[0] >= instance.M {
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
			return instance.validate(child, step, deep)
		}
	}

	step--
	return true
}
