package search_tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTree(t *testing.T) {
	tr := &Tree{}
	n1 := &Node{Key: 1}
	tr.Insert(n1)
	n2 := &Node{Key: 2}
	(*Tree).Insert(tr, n2)
	n3 := &Node{Key: 3}
	tr.Insert(n3)
	n4 := &Node{Key: 4}
	tr.Insert(n4)

	result := tr.Root.Search(3)
	if result != n3 {
		t.Error("Search() is wrong")
	}
	result = tr.Root.Search(5)
	if result != nil {
		t.Error("Search() is wrong")
	}
	min := tr.Root.Minimum()
	if min.Key != 1 {
		t.Error("Minimum() is wrong")
	}
	max := tr.Root.Maximum()
	if max.Key != 4 {
		t.Error("Maximum() is wrong")
	}
	successor := n2.Successor()
	if successor != n3 {
		t.Error("Successor() is wrong")
	}
	predecessor := n2.Predecessor()
	if predecessor != n1 {
		t.Error("Predecessor() is wrong")
	}
	tr.Transplant(n1, n3)
	if tr.Root != n3 {
		t.Error("Transplant() is wrong")
	}
	tr.Delete(n3)
	if tr.Root != n4 {
		t.Error("Delete() is wrong")
	}

	outputJson(tr.Root, "binary-tree.json")
}

func outputJson(root *Node, fileName string) {
	jsonContent, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile(fileName, jsonContent, 0644)
}
