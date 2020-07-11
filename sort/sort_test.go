package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItems[0])

	counter := insertionSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func TestMergeSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItems[0])

	counter := mergeSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func TestHeapSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItems[0])

	counter := heapSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func TestQuickSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItems[0])

	counter := quickSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func validation(t *testing.T, testItem []int) {
	for i := 1; i < length; i++ {
		if testItem[i] < testItem[i-1] {
			t.Errorf("Unexpected value detected at %d", i)
			return
		}
	}
	t.Log("Correct")
}
