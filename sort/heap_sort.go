package sort

func left(i int) int {
	i = i << 1
	i++
	return i
}

func right(i int) int {
	i++
	return i << 1
}

func parent(i int) int { // This function is not used in heapSort()
	i--
	i = i >> 1
	return i
}

func exchange(nums []int, from, to int) {
	temp := nums[from]
	nums[from] = nums[to]
	nums[to] = temp
}

func maxHeapify(nums []int, index int) (counter int) {
	l := left(index)
	r := right(index)
	largest := index

	if l < len(nums) && nums[l] > nums[largest] {
		largest = l
	}
	if r < len(nums) && nums[r] > nums[largest] {
		largest = r
	}

	if largest != index {
		exchange(nums, index, largest)
		counter += maxHeapify(nums, largest) + 1
	}
	return
}

func HeapSort(nums []int) (counter int) {
	length := len(nums)

	for i := len(nums) / 2; i >= 0; i-- {
		counter += maxHeapify(nums, i)
	}

	for i := len(nums) - 1; i > 0; i-- {
		exchange(nums, 0, i)
		nums = nums[:len(nums)-1]
		counter += maxHeapify(nums, 0) + 1
	}

	nums = nums[:length]
	return
}
