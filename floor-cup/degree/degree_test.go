package degree

import "testing"

func TestSumCompose(t *testing.T) {
	testData := [][]int{
		{5, 1, 1},
		{5, 2, 6},
		{5, 3, 16},
		{5, 4, 26},
		{5, 5, 31},
		{5, 6, 32},
		{6, 3, 22},
		{6, 4, 42},
		{6, 5, 57},
		{6, 6, 63},
	}
	for _, testItem := range testData {
		sum := sumCompose(testItem[0], testItem[1])
		if sum == testItem[2] {
			t.Logf("correct sum: %d\n", sum)
		} else {
			t.Errorf("invalid sum: expect %d but get %d\n", testItem[2], sum)
		}
	}
}

func TestInnerCalculate(t *testing.T) {
	testFuncs := []func(int, int) int{
		InnerCalculateA,
		InnerCalculateB,
		InnerCalculateC,
	}
	testData := [][]int{
		{7, 3, 3},
		{8, 3, 4},
		{7, 30, 3},
		{8, 30, 4},
		{14, 3, 4},
		{15, 3, 5},
		{25, 3, 5},
		{26, 3, 6},
		{126, 6, 7},
		{127, 6, 8},
		{246, 6, 8},
		{247, 6, 9},
		{511, 35, 9},
		{512, 35, 10},
		{7098, 8, 13},
		{7099, 8, 14},
	}
	for i, testFunc := range testFuncs {
		t.Logf("======== Testing for function %d ========\n", i)
		for _, testItem := range testData {
			degree := Calculate(testItem[0], testItem[1], testFunc)
			if degree == testItem[2] {
				t.Logf("correct degree: %d\n", degree)
			} else {
				t.Errorf("invalid degree: expect %d but get %d\n", testItem[2], degree)
			}
		}
	}
}

var (
	benchmarkData = [][]int{
		{511, 35, 9},
		{512, 35, 10},
	}
	//benchmarkData = [][]int{
	//	{7098, 8, 13},
	//	{7099, 8, 14},
	//}
	//benchmarkData = [][]int{
	//	{150676185, 10, 33},
	//	{150676186, 10, 34},
	//}
	//benchmarkData = [][]int{
	//	{1073741822, 29, 30},
	//	{1073741823, 29, 31},
	//}
	//benchmarkData = [][]int{
	//	{8589928573, 29, 33},
	//	{8589928574, 29, 34},
	//}
)

func BenchmarkCalculateA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testItem := range benchmarkData {
			degree := Calculate(testItem[0], testItem[1], InnerCalculateA)
			if degree != testItem[2] {
				b.Errorf("invalid degree: expect %d but get %d\n", testItem[2], degree)
			}
		}
	}
}

func BenchmarkCalculateB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testItem := range benchmarkData {
			degree := Calculate(testItem[0], testItem[1], InnerCalculateB) // Calculate() is the mandatory caller for InnerCalculateB
			if degree != testItem[2] {
				b.Errorf("invalid degree: expect %d but get %d\n", testItem[2], degree)
			}
		}
	}
}

func BenchmarkCalculateC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testItem := range benchmarkData {
			degree := Calculate(testItem[0], testItem[1], InnerCalculateC)
			if degree != testItem[2] {
				b.Errorf("invalid degree: expect %d but get %d\n", testItem[2], degree)
			}
		}
	}
}

func BenchmarkInnerCalculateA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testItem := range benchmarkData {
			degree := InnerCalculateA(testItem[0], testItem[1])
			if degree != testItem[2] {
				b.Errorf("invalid degree: expect %d but get %d\n", testItem[2], degree)
			}
		}
	}
}

func BenchmarkInnerCalculateC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testItem := range benchmarkData {
			degree := InnerCalculateC(testItem[0], testItem[1])
			if degree != testItem[2] {
				b.Errorf("invalid degree: expect %d but get %d\n", testItem[2], degree)
			}
		}
	}
}
