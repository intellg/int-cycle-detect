package sort

const (
	weight = 16 // The weight value is estimated only by compare with other sort result
)

func BucketSort(nums []int) int {
	max := getMax(nums)
	array := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index := (len(nums) - 1) * nums[i] / max
		array[index] = append(array[index], nums[i])
	}
	counter := len(nums) * weight

	result := make([]int, 0)
	for i := 0; i < len(array); i++ {
		if len(array[i]) > 0 {
			counter += InsertionSort(array[i])
			result = append(result, array[i]...)
			counter += len(array[i])
		}
	}

	copy(nums, result)
	return counter
}
