package node

import "testing"

func TestCalculate(t *testing.T) {
	testData := [][]int{
		{10, 2, 4, 4},
		{10, 3, 4, 7},
		{100, 2, 14, 14},
		{1000, 7, 11, 499},
		{1000, 8, 10, 502},
	}
	for _, testItem := range testData {
		root := Calculate(testItem[0], testItem[1], testItem[2])
		if root.Value == testItem[3] {
			t.Logf("correct root value: %d\n", root.Value)
		} else {
			t.Errorf("invalid root value: expect %d bug get %d\n", testItem[3], root.Value)
		}
		sumCount := root.LeftCount + root.RightCount + 1
		if sumCount == testItem[0] {
			t.Logf("correct sum count: %d\n", sumCount)
		} else {
			t.Errorf("invalid sum count: expect %d bug get %d\n", testItem[0], sumCount)
		}
	}
}
