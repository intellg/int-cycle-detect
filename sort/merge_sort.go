package sort

func mergeSort(nums []int) (counter int) {
	if len(nums) == 1 {
		return
	}

	mid := len(nums) / 2
	numsL := make([]int, mid)
	copy(numsL, nums[:mid])
	numsR := make([]int, len(nums)-mid)
	copy(numsR, nums[mid:])
	counterL := mergeSort(numsL)
	counterR := mergeSort(numsR)
	counter = counterL + counterR

	i := 0
	j := 0
	k := 0
	for i < len(numsL) && j < len(numsR) {
		if numsL[i] <= numsR[j] {
			nums[k] = numsL[i]
			counter++
			i++
		} else {
			nums[k] = numsR[j]
			counter++
			j++
		}
		k++
	}

	if i == len(numsL) { // sortedR has remaining data
		counter += copy(nums[k:], numsR[j:])
	} else { // sortedL has remaining data
		counter += copy(nums[k:], numsL[i:])
	}

	return counter
}
