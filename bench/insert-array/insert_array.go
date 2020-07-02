package insert_array

func insertWithPrepare(arr []int, index, item int) []int {
	arr = append(arr, arr[len(arr)-1])
	copy(arr[index+1:], arr[index:len(arr)-1])
	arr[index] = item
	return arr
}

func insertWith0(arr []int, index, item int) []int {
	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = item
	return arr
}

func insertByIteration(arr []int, index, item int) []int {
	arr = append(arr, arr[len(arr)-1])
	for i := len(arr) - 2; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = item
	return arr
}
