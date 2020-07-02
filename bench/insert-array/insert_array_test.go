package insert_array

import "testing"

var InsertArrayItems = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func BenchmarkInsertWithPrepare(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertWithPrepare(InsertArrayItems, 5, 100)
	}
}

func BenchmarkInsertWith0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertWith0(InsertArrayItems, 5, 100)
	}
}

func BenchmarkInsertByIteration(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertByIteration(InsertArrayItems, 5, 100)
	}
}
