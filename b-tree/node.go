package b_tree

type Node struct {
	Index    int     `json:"index"`
	Keys     []int   `json:"keys"`
	Children []*Node `json:"children"`
	Parent   *Node   `json:"-"`
}

func (instance *Node) SetKeys(keys ...int) {
	instance.Keys = make([]int, 1)
	instance.Keys[0] = len(keys)
	instance.Keys = append(instance.Keys, keys...)
}

func (instance *Node) SetChildren(children []*Node) {
	instance.Children = children
	for i, child := range children {
		child.Index = i
	}
}

func (instance *Node) GetLeftestLeafFrom(index int) *Node {
	node := instance.Children[index]
	for node.Children != nil {
		node = node.Children[0]
	}
	return node
}

func (instance *Node) SearchKey(key int) (index, step int, ok bool) {
	end := len(instance.Keys)
	if end <= 1 {
		index = 1
		return
	}
	start := 1

	var previousMid = 0.0
	for true {
		step++

		mid := (float64(start) + float64(end)) / 2
		index = int(mid)
		if mid == previousMid {
			if mid == float64(index) {
				index--
			}
			return
		} else {
			previousMid = mid
		}

		if key == instance.Keys[index] {
			ok = true
			return
		} else if key < instance.Keys[index] {
			end = index
		} else { // if key > instance.Keys[intMid]
			start = index
		}
	}
	return
}

func (instance *Node) AddKey(tree *Tree, index, key int, child *Node) {
	instance.Keys = insertInt(instance.Keys, index, key)
	instance.Keys[0] ++

	if instance.Children != nil {
		for i := index; i < len(instance.Children); i++ {
			instance.Children[i].Index++
		}
		instance.Children = insertNode(instance.Children, index, child)
		child.Parent = instance
		child.Index = index
	}

	if instance.Keys[0] == tree.M {
		instance.adjustOverflow(tree)
	}
}

func (instance *Node) adjustOverflow(tree *Tree) {
	// 0. Prepare necessary data
	mid := tree.M/2 + 1
	newNode := &Node{}
	newNode.SetKeys(instance.Keys[mid+1:]...)
	if instance.Children != nil {
		children := make([]*Node, tree.M-mid+1)
		copy(children, instance.Children[mid:])
		newNode.SetChildren(children)
	}

	// 1. Handle parent instance
	if instance.Parent != nil {
		// Node instance.Keys[mid] is upgraded
		instance.Parent.AddKey(tree, instance.Index+1, instance.Keys[mid], newNode)
	} else {
		tree.Root = &Node{}
		tree.Root.SetKeys(instance.Keys[mid])
		tree.Root.SetChildren([]*Node{instance, newNode})
		instance.Parent = tree.Root
		newNode.Parent = tree.Root
	}

	// 2. Handle new instance
	for _, child := range newNode.Children {
		child.Parent = newNode
	}

	// 3. Handle instance instance
	instance.Keys = instance.Keys[0:mid]
	instance.Keys[0] = mid - 1
	if instance.Children != nil {
		instance.Children = instance.Children[0:mid]
	}
}

func (instance *Node) DeleteKey(tree *Tree, index int) {
	deleteNode := instance
	deleteIndex := index

	if instance.Children != nil {
		// Change the key with the right child's smallest key
		rightChildSmallestLeaf := instance.GetLeftestLeafFrom(index)
		instance.Keys[index] = rightChildSmallestLeaf.Keys[1]
		deleteNode = rightChildSmallestLeaf
		deleteIndex = 1
	}

	// Remove the key from the leaf node then adjust underflow if any
	deleteNode.Keys = append(deleteNode.Keys[:deleteIndex], deleteNode.Keys[deleteIndex+1:]...)
	deleteNode.Keys[0] --
	if deleteNode.Keys[0] < tree.Underflow {
		deleteNode.adjustUnderflow(tree)
	}
}

func (instance *Node) adjustUnderflow(tree *Tree) {
	if instance.Parent == nil {
		if instance.Keys[0] == 0 {
			if instance.Children != nil {
				tree.Root = instance.Children[0]
				tree.Root.Index = -1
				tree.Root.Parent = nil
			}
		}
	}

	// 0. Borrow from brother logic
	leftNode := instance
	rightNode := instance
	if instance.Index > 0 { // Try to borrow from left
		leftNode = instance.Parent.Children[instance.Index-1]
		if leftNode.Keys[0] > tree.Underflow {
			instance.borrowFromBrother(leftNode, 1, instance.Index, -1)
			return
		}
	} else { // Try to borrow from right
		rightNode = instance.Parent.Children[instance.Index+1]
		if rightNode.Keys[0] > tree.Underflow {
			instance.borrowFromBrother(rightNode, instance.Keys[0]+1, instance.Index+1, 1)
			return
		}
	}

	// 1. Combine right node into left node
	// 1.1 Handle left node's children
	if leftNode.Children != nil {
		for _, child := range rightNode.Children {
			leftNode.Children = append(leftNode.Children, child)
			child.Parent = leftNode
			child.Index += leftNode.Keys[0] + 1
		}
	}

	// 1.2 Handle left node's keys
	leftNode.Keys = append(leftNode.Keys, leftNode.Parent.Keys[rightNode.Index])
	leftNode.Keys = append(leftNode.Keys, rightNode.Keys[1:]...)
	leftNode.Keys[0] += rightNode.Keys[0] + 1

	// 2. Handle parent node
	// 2.1 Handle parent node's children
	leftNode.Parent.Children = append(leftNode.Parent.Children[:rightNode.Index], leftNode.Parent.Children[rightNode.Index+1:]...)
	if leftNode.Parent.Keys[0] > leftNode.Index {
		for _, child := range leftNode.Parent.Children[rightNode.Index:] {
			child.Index--
		}
	}

	// 2.2 Handle parent node's keys
	leftNode.Parent.Keys = append(leftNode.Parent.Keys[:rightNode.Index], leftNode.Parent.Keys[rightNode.Index+1:]...)
	leftNode.Parent.Keys[0] --
	if leftNode.Parent.Keys[0] < tree.Underflow {
		leftNode.Parent.adjustUnderflow(tree)
	}
}

func (instance *Node) borrowFromBrother(brother *Node, currentIndex, parentIndex, brotherIndex int) {
	// 1. Move the key from parent node to self node
	instance.Keys = insertInt(instance.Keys, currentIndex, instance.Parent.Keys[parentIndex])
	instance.Keys[0]++

	// 2. Move the key from brother node to parent node
	instance.Parent.Keys[parentIndex] = brother.Keys[parentIndex]
	brother.Keys = append(brother.Keys[:parentIndex], brother.Keys[parentIndex+1:]...)
	brother.Keys[0] --

	// 3. Move the child from brother to current node
	if instance.Children != nil {
		if brotherIndex == -1 { // Borrow from left
			instance.Children = insertNode(instance.Children, 0, brother.Children[len(brother.Children)-1])
			brother.Children = brother.Children[:len(brother.Children)-1]
			instance.Children[0].Parent = instance
			for i := 0; i < instance.Keys[0]+1; i++ {
				instance.Children[i].Index = i
			}
		} else { // Borrow from right
			instance.Children = append(instance.Children, brother.Children[0])
			brother.Children = brother.Children[1:len(brother.Children)]
			instance.Children[len(instance.Children)-1].Parent = instance
			instance.Children[len(instance.Children)-1].Index = instance.Keys[0]
			for i := 0; i < brother.Keys[0]+1; i++ {
				brother.Children[i].Index = i
			}
		}
	}
}

func insertInt(arr []int, index, item int) []int {
	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = item
	return arr
}

func insertNode(arr []*Node, index int, item *Node) []*Node {
	arr = append(arr, nil)
	copy(arr[index+1:], arr[index:])
	arr[index] = item
	return arr
}
