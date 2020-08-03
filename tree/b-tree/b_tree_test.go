package b_tree

import (
	"math/rand"
	"testing"
	"time"
)

const (
	degree = 2
	length = 10000
)

func TestAll(t *testing.T) {
	tree := &Tree{}
	tree.Init(degree)
	m := make(map[int]int, length)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		key := r.Intn(length << 6)
		if _, ok := m[key]; ok {
			i--
			continue
		}
		m[key] = 0
		t.Logf("%8d tree.Insert(%d)", i, key)
		tree.Insert(key)
		if !tree.Validate() {
			t.Errorf("Wrong for insert %d", key)
			tree.Root.OutputJson("testInsert.json")
			return
		}
	}

	i := 0
	for key := range m {
		t.Logf("%8d tree.Delete(%d)", i, key)
		tree.Delete(key)
		if !tree.Validate() {
			t.Errorf("Wrong for delete %d", key)
			tree.Root.OutputJson("testDelete.json")
			return
		}
		i++
	}
}

func TestSpecific(t *testing.T) {
	tree := &Tree{}
	tree.Init(2)

	tree.Insert(4620)
	tree.Insert(5407)
	tree.Insert(3799)
	tree.Insert(2130)
	tree.Insert(1570)
	tree.Insert(1204)
	tree.Insert(1064)
	tree.Insert(2104)
	tree.Insert(1920)
	tree.Insert(209)
	tree.Insert(6215)
	tree.Insert(3377)
	tree.Insert(6092)
	tree.Insert(869)
	tree.Insert(4144)
	tree.Insert(5201)
	tree.Insert(5855)
	tree.Insert(512)
	tree.Insert(854)
	tree.Insert(1055)
	tree.Insert(3612)
	tree.Insert(1459)
	tree.Insert(4891)
	tree.Insert(4662)
	tree.Insert(88)
	tree.Insert(5161)
	tree.Insert(9)
	tree.Insert(87)
	tree.Insert(728)
	tree.Insert(4273)
	tree.Insert(2233)
	tree.Insert(2826)
	tree.Insert(3073)
	tree.Insert(1133)
	tree.Insert(1363)
	tree.Insert(3289)
	tree.Insert(681)
	tree.Insert(4443)
	tree.Insert(5623)
	tree.Insert(1231)
	tree.Insert(4620)
	tree.Insert(3359)
	tree.Insert(3229)
	tree.Insert(3071)
	tree.Insert(5739)
	tree.Insert(831)
	tree.Insert(4001)
	tree.Insert(374)
	tree.Insert(1980)
	tree.Insert(1215)
	tree.Insert(3125)
	tree.Insert(1327)
	tree.Insert(3006)
	tree.Insert(2292)
	tree.Insert(3011)
	tree.Insert(6062)
	tree.Insert(1959)
	tree.Insert(3181)
	tree.Insert(5967)
	tree.Insert(3198)
	tree.Insert(5639)
	tree.Insert(2536)
	tree.Insert(5095)
	tree.Insert(4744)
	tree.Insert(5534)
	tree.Insert(4966)
	tree.Insert(1354)
	tree.Insert(5179)
	tree.Insert(1183)
	tree.Insert(3603)
	tree.Insert(2792)
	tree.Insert(1396)
	tree.Insert(4902)
	tree.Insert(1262)
	tree.Insert(6218)
	tree.Insert(266)
	tree.Insert(1367)
	tree.Insert(1343)
	tree.Insert(1994)
	tree.Insert(508)
	tree.Insert(1505)
	tree.Insert(6217)
	tree.Insert(3621)
	tree.Insert(5149)
	tree.Insert(4488)
	tree.Insert(3328)
	tree.Insert(6211)
	tree.Insert(2587)
	tree.Insert(5949)
	tree.Insert(2213)
	tree.Insert(6091)
	tree.Insert(1773)
	tree.Insert(6218)
	tree.Insert(307)
	tree.Insert(785)
	tree.Insert(1626)
	tree.Insert(5645)
	tree.Insert(5002)
	tree.Insert(1875)
	tree.Insert(6310)
	tree.Delete(5179)
	tree.Delete(1204)
	tree.Delete(1064)
	tree.Delete(2104)
	tree.Delete(2826)
	tree.Delete(3125)
	validate(t, tree)
	tree.Delete(2292)
	validate(t, tree)
	tree.Delete(4966)
	validate(t, tree)
	tree.Delete(5949)
	validate(t, tree)

	tree.Root.OutputJson("TestSpecific0001.json")
	tree.Delete(854)
	tree.Root.OutputJson("TestSpecific0002.json")

	validate(t, tree)
}

func validate(t *testing.T, tree *Tree) {
	if !tree.Validate() {
		t.Errorf("Wrong")
	}
}
