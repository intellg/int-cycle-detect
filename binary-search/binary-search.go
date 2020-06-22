package main

import "fmt"

const (
	length        = 10000000
	overNumber    = 0
	displayDetail = false
)

var (
	list []int
)

func init() {
	list = make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = i
	}
}

func search(list []int, val int, calculateMid func(int, int) int) (count int, isFound bool) {
	from := 0
	to := len(list)
	for from < to {
		count++
		mid := calculateMid(from, to)
		if mid == from {
			return
		}

		if list[mid] == val {
			isFound = true
			return
		} else if list[mid] > val {
			to = mid
		} else { // list[mid] < val
			from = mid
		}
	}
	return
}

func testingList(calculateMid func(int, int) int) {
	result := make(map[int]int, 0)
	for val := 0 - overNumber; val < length+overNumber; val++ {
		count, ok := search(list, val, calculateMid)
		displayResult(ok, val, count)
		result[count] ++
	}

	fmt.Println(result)

	sum := 0
	for k, v := range result {
		sum += k * v
	}
	fmt.Printf("Average %d\n", sum/(length+2*overNumber))
}

func displayResult(ok bool, val, count int) {
	if displayDetail {
		var status string
		if ok {
			status = "O"
		} else {
			status = "X"
		}
		fmt.Printf("%s for %d with search count %d\n", status, val, count)
	}
}

func half(from, to int) int {
	return (from + to) / 2
}

func oneThird(from, to int) int {
	if to-from == 2 {
		return to - 1
	}
	return (2*from + to) / 3
}

func main() {
	testingList(half)
	testingList(oneThird)
}
