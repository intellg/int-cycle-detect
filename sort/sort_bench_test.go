package sort

import (
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertionSort(actualTestItem)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeSort(actualTestItem)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	actualTestItem := make([]int, length)
	copy(actualTestItem, testItem)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heapSort(actualTestItem)
	}
}
