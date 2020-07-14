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
	tr := &Tree{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		node := &Node{Key: r.Intn(length << 6)}
		tr.Insert(node)
		tr.validate(t)
	}

	outputJson(tr.Root, "red-black-tree.json")
}

func outputJson(root *Node, fileName string) {
	jsonContent, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile(fileName, jsonContent, 0644)
}
