package sort

func RadixSort(nums []int) int {
	max := getMax(nums)
	rng := 10
	counter := len(nums)
	arr := nums

	at := 1
	for max > 0 {
		arr = countingSort(arr, rng, at)
		max /= 10
		at *= 10
		counter += len(nums)*2 + rng
	}

	copy(nums, arr)
	return counter + len(nums)
}

func countingSort(nums []int, rng, at int) (result []int) {
	arrRange := make([]int, rng)
	for i := 0; i < len(nums); i++ {
		dgt := getDigitAt(nums[i], at)
		arrRange[dgt] = arrRange[dgt] + 1
	}
	// arrRange[i] now contains the number of elements equal to i

	for i := 1; i < rng; i++ {
		arrRange[i] = arrRange[i] + arrRange[i-1]
	}
	// arrRange[i] now contains the number of elements less than or equal to i

	result = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		dgt := getDigitAt(nums[i], at)
		result[arrRange[dgt]-1] = nums[i]
		arrRange[dgt]--
	}
	return
}

func getDigitAt(num, at int) int {
	return num % (at * 10) / at
}

func getMax(nums []int) (max int) {
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return
}
