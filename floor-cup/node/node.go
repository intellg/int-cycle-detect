package node

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

type node struct {
	Value      int   `json:"V"`
	Left       *node `json:"L"`
	Right      *node `json:"R"`
	Parent     *node `json:"-"`
	LeftCount  int   `json:"-"`
	RightCount int   `json:"-"`
	Remain     int   `json:"-"`
	IsLeft     bool  `json:"-"`
}

func Calculate(floor, cup, degree int) *node {
	// 1. Prepare the root node
	root := node{}
	root.Remain = cup - 1
	nodeList := make([][]*node, degree)
	nodeList[0] = make([]*node, 0)
	nodeList[0] = append(nodeList[0], &root)

	// 2.1 Construct the binary tree without bottom layer
	sum := 1
	for i := 1; i < degree-1; i++ {
		nodeList[i] = make([]*node, 0, int(math.Pow(2, float64(i))))
		for _, eachNode := range nodeList[i-1] {
			sum += addBothNode(eachNode, &nodeList[i])
		}
	}
	// 2.2 Construct the bottom layer
	restNodeNumber := floor - sum
	for _, eachNode := range nodeList[degree-2] {
		if restNodeNumber == 0 {
			break
		}
		if restNodeNumber == 1 {
			addSingleNode(eachNode, &nodeList[degree-1])
			break
		}
		restNodeNumber -= addBothNode(eachNode, &nodeList[degree-1])
	}

	// 3. Calculate the left count and right count for each node
	for i := degree - 2; i >= 0; i-- {
		adjustCount(&nodeList[i])
	}

	// 4. Fill in the value number fro each node
	for i := 0; i < degree; i++ {
		fillValue(&nodeList[i])
	}

	return &root
}

func addBothNode(parent *node, nodeList *[]*node) int {
	return addNode(parent, nodeList, false)
}

func addSingleNode(parent *node, nodeList *[]*node) int {
	return addNode(parent, nodeList, true)
}

func addNode(parent *node, nodeList *[]*node, single bool) (count int) {
	if parent.Remain > 0 {
		left := node{}
		left.Parent = parent
		left.Remain = parent.Remain - 1
		left.IsLeft = true
		parent.Left = &left
		*nodeList = append(*nodeList, &left)

		if single {
			return 1
		} else {
			count = 2
		}
	} else {
		count = 1
	}

	right := node{}
	right.Parent = parent
	right.Remain = parent.Remain
	parent.Right = &right
	*nodeList = append(*nodeList, &right)
	return
}

func adjustCount(nodeList *[]*node) {
	for _, eachNode := range *nodeList {
		if eachNode.Left != nil {
			eachNode.LeftCount = eachNode.Left.LeftCount + eachNode.Left.RightCount + 1
		}
		if eachNode.Right != nil {
			eachNode.RightCount = eachNode.Right.LeftCount + eachNode.Right.RightCount + 1
		}
	}
}

func fillValue(nodeList *[]*node) {
	for _, eachNode := range *nodeList {
		if eachNode.Parent == nil { // root
			eachNode.Value = eachNode.LeftCount + 1
		} else {
			if eachNode.IsLeft { // Left node
				eachNode.Value = eachNode.Parent.Value - eachNode.RightCount - 1
			} else { // Right node
				eachNode.Value = eachNode.Parent.Value + eachNode.LeftCount + 1
			}
		}
	}
}

func OutputJson(root *node) {
	jsonContent, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile("./nodes.json", jsonContent, 0644)
}
