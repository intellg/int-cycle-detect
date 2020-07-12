package sort

import "testing"

func TestCounting(t *testing.T) {
	nums := []int{12, 25, 33, 40, 52, 63, 70, 83}
	result := countingSort(nums, 10, 10)
	t.Log(result)
}

func TestGetDigitAt(t *testing.T) {
	r := getDigitAt(123, 1)
	if r == 3 {
		t.Log("Correct")
	} else {
		t.Error("Wrong")
	}
	r = getDigitAt(123, 10)
	if r == 2 {
		t.Log("Correct")
	} else {
		t.Error("Wrong")
	}
	r = getDigitAt(123, 100)
	if r == 1 {
		t.Log("Correct")
	} else {
		t.Error("Wrong")
	}
}