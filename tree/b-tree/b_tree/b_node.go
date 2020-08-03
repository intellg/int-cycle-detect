package b_tree

type Node struct {
	Keys     []int   `json:"keys"`
	Children []*Node `json:"children,omitempty"`
	IsLeaf   bool    `json:"-"`
}

func (node *Node) Search(key int) (*Node, int) {
	i := 0
	for i < len(node.Keys) && key > node.Keys[i] {
		i++
	}
	if i < len(node.Keys) && key == node.Keys[i] {
		return node, i
	} else if node.IsLeaf {
		return nil, -1
	} else {
		return node.Children[i].Search(key)
	}
}

func (node *Node) splitChild(index int) {
	child := node.Children[index]
	degree := (len(child.Keys) + 1) / 2
	newChild := &Node{}
	newChild.IsLeaf = child.IsLeaf

	node.Keys = append(node.Keys, 0)
	copy(node.Keys[index+1:], node.Keys[index:len(node.Keys)-1])
	node.Keys[index] = child.Keys[degree-1]

	node.Children = append(node.Children, nil)
	copy(node.Children[index+2:], node.Children[index+1:len(node.Children)-1])
	node.Children[index+1] = newChild

	newChild.Keys = make([]int, degree-1)
	for j := 0; j < degree-1; j++ {
		newChild.Keys[j] = child.Keys[j+degree]
	}
	child.Keys = child.Keys[:degree-1]

	if !child.IsLeaf {
		newChild.Children = make([]*Node, degree)
		for j := 0; j < degree; j++ {
			newChild.Children[j] = child.Children[j+degree]
		}
		child.Children = child.Children[:degree]
	}
}

func (node *Node) insert(degree, key int) {
	if node.IsLeaf {
		node.Keys = append(node.Keys, 0)
		i := len(node.Keys) - 1
		for i > 0 && key < node.Keys[i-1] {
			node.Keys[i] = node.Keys[i-1]
			i--
		}
		node.Keys[i] = key
	} else {
		i := len(node.Keys) - 1
		for i >= 0 && key < node.Keys[i] {
			i--
		}
		i++
		if len(node.Children[i].Keys) == 2*degree-1 {
			node.splitChild(i)
			if key > node.Keys[i] {
				i++
			}
		}
		node.Children[i].insert(degree, key)
	}
}

func (node *Node) mergeChild(index int) {
	leftChild := node.Children[index]
	rightChild := node.Children[index+1]
	leftChild.Keys = append(append(leftChild.Keys, node.Keys[index]), rightChild.Keys...)
	if !leftChild.IsLeaf {
		leftChild.Children = append(leftChild.Children, rightChild.Children...)
	}

	for i := index; i < len(node.Keys)-1; i++ {
		node.Keys[i] = node.Keys[i+1]
	}
	node.Keys = node.Keys[:len(node.Keys)-1]
	for i := index + 1; i < len(node.Children)-1; i++ {
		node.Children[i] = node.Children[i+1]
	}
	node.Children = node.Children[:len(node.Children)-1]
}

//        [6]            [4]
//      ┌──┴──┐        ┌──┴──┐
//    [2,4]  [8]  =>  [2]  [6,8]
//    ┌─┼─┐ ┌─┴─┐    ┌─┴─┐ ┌─┼─┐
//    1 3 5 7   9    1   3 5 7 9
func (node *Node) borrowLeft(index int) {
	currentChild := node.Children[index]
	leftChild := node.Children[index-1]

	// move parent's key to current child
	currentChild.Keys = append(currentChild.Keys, 0)
	copy(currentChild.Keys[1:], currentChild.Keys[0:])
	currentChild.Keys[0] = node.Keys[index-1]
	if !currentChild.IsLeaf {
		// move sibling's nearest child to current child
		currentChild.Children = append(currentChild.Children, nil)
		copy(currentChild.Children[1:], currentChild.Children[0:])
		currentChild.Children[0] = leftChild.Children[len(leftChild.Children)-1]
	}

	// move sibling's nearest key to node.key
	node.Keys[index-1] = leftChild.Keys[len(leftChild.Keys)-1]

	// shrink sibling
	leftChild.Keys = leftChild.Keys[:len(leftChild.Keys)-1]
	if !leftChild.IsLeaf {
		leftChild.Children = leftChild.Children[:len(leftChild.Children)-1]
	}
}

//        [4]            [6]
//      ┌──┴──┐        ┌──┴──┐
//     [2]  [6,8] => [2,4]  [8]
//    ┌─┴─┐ ┌─┼─┐    ┌─┼─┐ ┌─┴─┐
//    1   3 5 7 9    1 3 5 7   9
func (node *Node) borrowRight(index int) {
	currentChild := node.Children[index]
	rightChild := node.Children[index+1]

	// move parent's key to current child
	currentChild.Keys = append(currentChild.Keys, node.Keys[index])
	if !currentChild.IsLeaf {
		// move sibling's nearest child to current child
		currentChild.Children = append(currentChild.Children, rightChild.Children[0])
	}

	// move left child's last key to node.key
	node.Keys[index] = rightChild.Keys[0]

	// shrink left child
	copy(rightChild.Keys, rightChild.Keys[1:])
	rightChild.Keys = rightChild.Keys[:len(rightChild.Keys)-1]
	if !rightChild.IsLeaf {
		copy(rightChild.Children, rightChild.Children[1:])
		rightChild.Children = rightChild.Children[:len(rightChild.Children)-1]
	}
}

func (node *Node) delete(degree, key int) bool {
	if node.IsLeaf {
		for i := 0; i < len(node.Keys); i++ {
			if key == node.Keys[i] {
				copy(node.Keys[i:], node.Keys[i+1:])
				node.Keys = node.Keys[:len(node.Keys)-1]
				return true
			} else if key < node.Keys[i] {
				break
			}
		}
		return false
	} else {
		var i int
		for i = 0; i < len(node.Keys); i++ {
			if key == node.Keys[i] {
				leftChild := node.Children[i]
				if len(leftChild.Keys) > degree-1 {
					// borrow from left
					predecessor := leftChild
					for !predecessor.IsLeaf {
						predecessor = predecessor.Children[len(predecessor.Children)-1]
					}
					node.Keys[i] = predecessor.Keys[len(predecessor.Keys)-1]
					leftChild.delete(degree, node.Keys[i])
				} else {
					rightChild := node.Children[i+1]
					if len(rightChild.Keys) > degree-1 {
						// borrow from right
						successor := rightChild
						for !successor.IsLeaf {
							successor = successor.Children[0]
						}
						node.Keys[i] = successor.Keys[0]
						rightChild.delete(degree, node.Keys[i])
					} else {
						// merge children
						node.mergeChild(i)
						leftChild.delete(degree, key)
					}
				}
				return true
			} else if key < node.Keys[i] {
				break
			}
		}

		if len(node.Children[i].Keys) == degree-1 {
			borrowed := false
			if i > 0 {
				if len(node.Children[i-1].Keys) > degree-1 {
					node.borrowLeft(i)
					borrowed = true
				}
			}
			if !borrowed && i < len(node.Keys) {
				if len(node.Children[i+1].Keys) > degree-1 {
					node.borrowRight(i)
					borrowed = true
				}
			}
			if !borrowed {
				if i != 0 && i == len(node.Children)-1 {
					i--
				}
				node.mergeChild(i)
			}
		}

		return node.Children[i].delete(degree, key)
	}
}

func (node *Node) Validate(degree, parentDepth int, depthIn []int) (ok bool, depthOut []int) {
	if len(node.Keys) == 0 ||
		(!node.IsLeaf && len(node.Keys) != len(node.Children)-1) ||
		len(node.Keys) < degree-1 || len(node.Keys) > degree*2-1 {
		return
	}
	for i := 1; i < len(node.Keys); i++ {
		if node.Keys[i-1] > node.Keys[i] {
			return
		}
	}
	if node.IsLeaf {
		ok = true
		depthOut = append(depthIn, parentDepth+1)
		return
	} else {
		depthOut = depthIn
		for i := 0; i < len(node.Children); i++ {
			ok = false
			child := node.Children[i]
			for j := 0; j < len(child.Keys); j++ {
				if i > 0 && child.Keys[j] < node.Keys[i-1] {
					return
				}
				if i < len(node.Children)-1 && child.Keys[j] > node.Keys[i] {
					return
				}
			}
			ok, depthOut = child.Validate(degree, parentDepth+1, depthOut)
			if !ok {
				return
			}
		}
		return
	}
}
