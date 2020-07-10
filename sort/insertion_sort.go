package sort

func insertionSort(nums []int) (counter int) {
	for j := 1; j < len(nums); j++ {
		key := nums[j]
		// Insert A[j] into the sorted sequence A[1..j-1]
		i := j - 1
		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			counter++
			i--
		}
		nums[i+1] = key
		counter++
	}
	return
}
