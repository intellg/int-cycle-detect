package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	base   = 200000000
	length = 1000000
)

var testItems = make([][]int, base/length)

func init() {
	fmt.Printf("length: %d\n", length)
	for i := range testItems {
		testItems[i] = make([]int, length)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for j := 0; j < length; j++ {
			testItems[i][j] = r.Intn(length << 6)
		}
	}
}

//func BenchmarkInsertionSort(b *testing.B) {
//	actualTestItems := prepareActualTestItems()
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		c := insertionSort(actualTestItems[i])
//		logCounter(b, i, c)
//	}
//}

func BenchmarkMergeSort(b *testing.B) {
	actualTestItems := prepareActualTestItems()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := mergeSort(actualTestItems[i])
		logCounter(b, i, c)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	actualTestItems := prepareActualTestItems()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := heapSort(actualTestItems[i])
		logCounter(b, i, c)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	actualTestItems := prepareActualTestItems()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := quickSort(actualTestItems[i])
		logCounter(b, i, c)
	}
}

func prepareActualTestItems() [][]int {
	actualTestItems := make([][]int, base/length)
	for i := range testItems {
		actualTestItems[i] = make([]int, length)
		copy(actualTestItems[i], testItems[i])
	}
	return actualTestItems
}

func logCounter(b *testing.B, i, counter int) {
	if i == -1 {
		b.Logf("counter is %d", counter)
	}
}
