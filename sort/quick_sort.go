package sort

func QuickSort(nums []int) int {
	return innerQuickSort(nums, 0, len(nums)-1)
}

func innerQuickSort(nums []int, from, to int) (counter int) {
	if from < to {
		var div int
		div, counter = partition(nums, from, to)
		counter += innerQuickSort(nums, from, div-1)
		counter += innerQuickSort(nums, div+1, to)
	}
	return
}

func partition(nums []int, from, to int) (div, counter int) {
	x := nums[to]
	div = from - 1
	for j := from; j < to; j++ {
		if nums[j] <= x {
			div++
			exchange(nums, div, j)
			counter++
		}
	}
	div++
	exchange(nums, div, to)
	counter++
	return
}
