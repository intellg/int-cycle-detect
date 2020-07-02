package range_calculate

import "testing"

func BenchmarkByIndex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byIndex()
	}
}

func BenchmarkByRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byRange()
	}
}
