package sort

import "testing"

func TestLeft(t *testing.T) {
	testItems := [][]int{
		{4, 9},
		{9, 19},
	}
	for _, testItem := range testItems {
		result := left(testItem[0])
		if result == testItem[1] {
			t.Logf("left(%d) is correct", testItem[0])
		} else {
			t.Errorf("left(%d) is expect %d, but get %d", testItem[0], testItem[1], result)
		}
	}
}

func TestRight(t *testing.T) {
	testItems := [][]int{
		{4, 10},
		{9, 20},
	}
	for _, testItem := range testItems {
		result := right(testItem[0])
		if result == testItem[1] {
			t.Logf("right(%d) is correct", testItem[0])
		} else {
			t.Errorf("right(%d) is expect %d, but get %d", testItem[0], testItem[1], result)
		}
	}
}

func TestParent(t *testing.T) {
	testItems := [][]int{
		{19, 9},
		{20, 9},
	}
	for _, testItem := range testItems {
		result := parent(testItem[0])
		if result == testItem[1] {
			t.Logf("parent(%d) is correct", testItem[0])
		} else {
			t.Errorf("parent(%d) is expect %d, but get %d", testItem[0], testItem[1], result)
		}
	}
}
