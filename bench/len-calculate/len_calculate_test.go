package len_calculate

import "testing"

func BenchmarkWithPreparation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withPreparation()
	}
}

func BenchmarkWithoutPreparation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withoutPreparation()
	}
}
