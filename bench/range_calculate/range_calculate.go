package range_calculate

const length = 100

func byIndex() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func byRange() int {
	array := make([]int, length)
	sum := 0
	for _, item := range array {
		sum += item
	}
	return sum
}
