package red_black_tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
)

const (
	length = 1000
)

func TestTree(t *testing.T) {
	tree := &Tree{}
	node := &Node{Key: 0}
	tree.Insert(node)

	nodes := make(map[int]*Node, length)
	nodes[0] = node
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < length; i++ {
		key := r.Intn(length << 6)
		if tree.Root.Search(key) == nil {
			node := &Node{Key: key}
			tree.Insert(node)
			nodes[i] = node
			tree.validate(t)
		} else {
			i--
		}
	}

	outputJson(tree.Root, "red-black-tree.json")

	for _, node := range nodes {
		tree.Delete(node)
		tree.validate(t)
	}
}

func TestSpecificTree(t *testing.T) {
	tree := &Tree{}
	n5 := &Node{Key: 5, Color: BLACK}
	tree.Insert(n5)

	n4 := &Node{Key: 4, Color: BLACK}
	n4.Parent = n5
	n5.Left = n4

	n7 := &Node{Key: 7, Color: RED}
	n7.Parent = n5
	n5.Right = n7

	n6 := &Node{Key: 6, Color: BLACK}
	n6.Parent = n7
	n7.Left = n6

	n8 := &Node{Key: 8, Color: BLACK}
	n8.Parent = n7
	n7.Right = n8

	tree.validate(t)
	tree.Delete(n4)
	tree.validate(t)
}

func outputJson(root *Node, fileName string) {
	jsonContent, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile(fileName, jsonContent, 0644)
}
