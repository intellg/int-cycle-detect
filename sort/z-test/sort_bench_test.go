package z_test

import (
	"fmt"
	"int-fun/sort"
	"math/rand"
	"testing"
	"time"
)

const (
	base   = 100000000
	length = 100000
)

var (
	testItem  = make([]int, length)
	testItems = make([][]int, base/length)
	subTests  = []struct {
		name     string
		function func([]int) int
	}{
		{"InsertSort", sort.InsertionSort},
		{"MergeSort ", sort.MergeSort},
		{"HeapSort  ", sort.HeapSort},
		{"QuickSort ", sort.QuickSort},
		{"RadixSort ", sort.RadixSort},
		{"BucketSort", sort.BucketSort},
	}
)

func setupBenchmarkSorts() {
	fmt.Printf("length: %d\n", length)
	fmt.Println("Preparing test data ...")
	for i := range testItems {
		testItems[i] = make([]int, length)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for j := 0; j < length; j++ {
			testItems[i][j] = r.Intn(length << 6)
		}
	}
}

func BenchmarkSorts(b *testing.B) {
	setupBenchmarkSorts()

	for _, st := range subTests {
		b.Run(st.name, func(b *testing.B) {
			actualTestItems := prepareActualTestItems()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				c := st.function(actualTestItems[i])
				logCounter(b, i, c)
			}
		})
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
