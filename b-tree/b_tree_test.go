package b_tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
)

const (
	maxM      = 20
	testCount = 50
	testRange = 1000
)

func TestInsertInt(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	arr = insertInt(arr, 0, 0)
	if arr[0] != 0 || len(arr) != 5 || arr[4] != 4 {
		t.Error("Wrong")
		return
	}
	arr = insertInt(arr, 4, 5)
	if len(arr) != 6 || arr[4] != 5 || arr[5] != 4 {
		t.Error("Wrong")
		return
	}

	arr = insertInt(arr, 6, 6)
	if len(arr) != 7 || arr[6] != 6 {
		t.Error("Wrong")
		return
	}
	t.Log("Correct")
}

func TestTreeRandom(t *testing.T) {
	for m := 3; m <= maxM; m++ {
		tree := &Tree{}
		tree.SetM(m)

		// Construct b-tree
		nodes := make(map[int]int, testCount)
		nodeList := make([]int, 0)
		for len(nodes) < testCount {
			key := rand.Intn(testRange)

			if ok := tree.AddKey(key); !ok {
				continue
			}
			nodes[key] = 0
			nodeList = append(nodeList, key)

			// Test b-tree
			t.Logf("# ======== test M = %d and round = %d ========", tree.M, len(nodes))
			if tree.validate(tree.Root, 0, -1) {
				t.Log("Valid")
			} else {
				t.Error("Invalid")
			}

			isFailed := false
			for i := 1; i <= testRange; i++ {
				if _, _, _, ok := tree.SearchKey(i, nil, 0); ok {
					if _, ok = nodes[i]; !ok {
						t.Errorf("%d should be found", i)
						isFailed = true
					}
				} else {
					if _, ok = nodes[i]; ok {
						t.Errorf("%d should not be found", i)
						isFailed = true
					}
				}
			}
			if isFailed {
				t.Error(nodeList)
				outputJson(tree.Root, fmt.Sprintf("nodes_%d_%d.json", tree.M, len(nodes)))
			}
		}
	}
}

func outputJson(root *Node, fileName string) {
	jsonContent, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile(fileName, jsonContent, 0644)
}
