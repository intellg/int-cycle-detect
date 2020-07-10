package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const length = 10000

var testItem = make([]int, length)

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		testItem[i] = r.Intn(length << 6)
	}
	fmt.Println("Prepared")
}

func TestInsertionSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

	counter := insertionSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func TestMergeSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

	counter := mergeSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func TestHeapSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

	counter := heapSort(actualTestItem)
	t.Log(counter)

	validation(t, actualTestItem)
}

func TestQuickSort(t *testing.T) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

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
