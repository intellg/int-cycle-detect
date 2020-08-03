package ref

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Node struct {
	Index    int     `json:"index"`
	Keys     []int   `json:"keys"`
	Children []*Node `json:"children"`
	Parent   *Node   `json:"-"`
}

func (node *Node) SetKeys(keys ...int) {
	node.Keys = make([]int, 1)
	node.Keys[0] = len(keys)
	node.Keys = append(node.Keys, keys...)
}

func (node *Node) SetChildren(children []*Node) {
	node.Children = children
	for i, child := range children {
		child.Index = i
	}
}

func (node *Node) GetLeftestLeafFrom(index int) *Node {
	n := node.Children[index]
	for n.Children != nil {
		n = n.Children[0]
	}
	return n
}

func (node *Node) SearchKey(key int) (index, step int, ok bool) {
	end := len(node.Keys)
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

		if key == node.Keys[index] {
			ok = true
			return
		} else if key < node.Keys[index] {
			end = index
		} else { // if key > node.Keys[intMid]
			start = index
		}
	}
	return
}

func (node *Node) AddKey(tree *Tree, index, key int, child *Node) {
	node.Keys = insertInt(node.Keys, index, key)
	node.Keys[0] ++

	if node.Children != nil {
		for i := index; i < len(node.Children); i++ {
			node.Children[i].Index++
		}
		node.Children = insertNode(node.Children, index, child)
		child.Parent = node
		child.Index = index
	}

	if node.Keys[0] == tree.M {
		node.adjustOverflow(tree)
	}
}

func (node *Node) adjustOverflow(tree *Tree) {
	// 0. Prepare necessary data
	mid := tree.M/2 + 1
	newNode := &Node{}
	newNode.SetKeys(node.Keys[mid+1:]...)
	if node.Children != nil {
		children := make([]*Node, tree.M-mid+1)
		copy(children, node.Children[mid:])
		newNode.SetChildren(children)
	}

	// 1. Handle parent node
	if node.Parent != nil {
		// Node node.Keys[mid] is upgraded
		node.Parent.AddKey(tree, node.Index+1, node.Keys[mid], newNode)
	} else {
		tree.Root = &Node{}
		tree.Root.SetKeys(node.Keys[mid])
		tree.Root.SetChildren([]*Node{node, newNode})
		node.Parent = tree.Root
		newNode.Parent = tree.Root
	}

	// 2. Handle new node
	for _, child := range newNode.Children {
		child.Parent = newNode
	}

	// 3. Handle node node
	node.Keys = node.Keys[0:mid]
	node.Keys[0] = mid - 1
	if node.Children != nil {
		node.Children = node.Children[0:mid]
	}
}

func (node *Node) DeleteKey(tree *Tree, index int) {
	deleteNode := node
	deleteIndex := index

	if node.Children != nil {
		// Change the key with the right child's smallest key
		rightChildSmallestLeaf := node.GetLeftestLeafFrom(index)
		node.Keys[index] = rightChildSmallestLeaf.Keys[1]
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

func (node *Node) adjustUnderflow(tree *Tree) {
	if node.Parent == nil {
		if node.Keys[0] == 0 {
			if node.Children != nil {
				tree.Root = node.Children[0]
				tree.Root.Index = -1
				tree.Root.Parent = nil
			}
		}
	}

	// 0. Borrow from brother logic
	leftNode := node
	rightNode := node
	if node.Index > 0 { // Try to borrow from left
		leftNode = node.Parent.Children[node.Index-1]
		if leftNode.Keys[0] > tree.Underflow {
			node.borrowFromBrother(leftNode, 1, node.Index, -1)
			return
		}
	} else { // Try to borrow from right
		rightNode = node.Parent.Children[node.Index+1]
		if rightNode.Keys[0] > tree.Underflow {
			node.borrowFromBrother(rightNode, node.Keys[0]+1, node.Index+1, 1)
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

func (node *Node) borrowFromBrother(brother *Node, currentIndex, parentIndex, brotherIndex int) {
	// 1. Move the key from parent node to self node
	node.Keys = insertInt(node.Keys, currentIndex, node.Parent.Keys[parentIndex])
	node.Keys[0]++

	// 2. Move the key from brother node to parent node
	node.Parent.Keys[parentIndex] = brother.Keys[parentIndex]
	brother.Keys = append(brother.Keys[:parentIndex], brother.Keys[parentIndex+1:]...)
	brother.Keys[0] --

	// 3. Move the child from brother to current node
	if node.Children != nil {
		if brotherIndex == -1 { // Borrow from left
			node.Children = insertNode(node.Children, 0, brother.Children[len(brother.Children)-1])
			brother.Children = brother.Children[:len(brother.Children)-1]
			node.Children[0].Parent = node
			for i := 0; i < node.Keys[0]+1; i++ {
				node.Children[i].Index = i
			}
		} else { // Borrow from right
			node.Children = append(node.Children, brother.Children[0])
			brother.Children = brother.Children[1:len(brother.Children)]
			node.Children[len(node.Children)-1].Parent = node
			node.Children[len(node.Children)-1].Index = node.Keys[0]
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

func (node *Node) OutputJson(fileName string) {
	jsonContent, err := json.MarshalIndent(node, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile(fileName, jsonContent, 0644)
}
