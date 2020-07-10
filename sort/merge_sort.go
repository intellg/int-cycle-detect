package sort

func mergeSort(nums []int) ([]int, int) {
	if len(nums) == 1 {
		return nums, 0
	}

	mid := len(nums) / 2
	numsL := make([]int, mid)
	copy(numsL, nums[:mid])
	numsR := make([]int, len(nums)-mid)
	copy(numsR, nums[mid:])
	sortedL, counterL := mergeSort(numsL)
	sortedR, counterR := mergeSort(numsR)
	counter := counterL + counterR

	i := 0
	j := 0
	k := 0
	for i < len(sortedL) && j < len(sortedR) {
		if sortedL[i] <= sortedR[j] {
			nums[k] = sortedL[i]
			counter++
			i++
		} else {
			nums[k] = sortedR[j]
			counter++
			j++
		}
		k++
	}

	if i == len(sortedL) { // sortedR has remaining data
		counter += copy(nums[k:], sortedR[j:])
	} else { // sortedL has remaining data
		counter += copy(nums[k:], sortedL[i:])
	}

	return nums, counter
}
