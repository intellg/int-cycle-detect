package bench

const length = 1000000

func test1() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += len(array)
	}
	return sum
}

func test2() int {
	array := make([]int, length)
	length := len(array)
	sum := 0
	for i := 0; i < length; i++ {
		sum += length
	}
	return sum
}

func test3() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func test4() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += array[i]
	}
	return sum
}

func test5() int {
	array := make([]int, length)
	sum := 0
	for _, item := range array {
		sum += item
	}
	return sum
}
