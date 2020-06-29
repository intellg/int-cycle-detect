package bench

import "testing"

func BenchmarkTest1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test1()
	}
}

func BenchmarkTest2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test2()
	}
}

func BenchmarkTest3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test3()
	}
}

func BenchmarkTest4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test4()
	}
}

func BenchmarkTest5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test5()
	}
}
