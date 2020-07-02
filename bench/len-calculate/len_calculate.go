package len_calculate

const length = 100

func withPreparation() int {
	array := make([]int, length)
	length := len(array)
	sum := 0
	for i := 0; i < length; i++ {
		sum += length
	}
	return sum
}

func withoutPreparation() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += len(array)
	}
	return sum
}
