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

func exchange(A []int, from, to int) {
	temp := A[from]
	A[from] = A[to]
	A[to] = temp
}

func maxHeapify(A []int, i int) (counter int) {
	l := left(i)
	r := right(i)
	largest := i

	if l < len(A) && A[l] > A[i] {
		largest = l
	}
	if r < len(A) && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		exchange(A, i, largest)
		counter++
		maxHeapify(A, largest)
	}
	return
}

func heapSort(A []int) (counter int) {
	length := len(A)

	for i := len(A) / 2; i >= 0; i-- {
		counter += maxHeapify(A, i)
	}

	for i := len(A) - 1; i > 0; i-- {
		exchange(A, 0, i)
		A = A[:len(A)-1]
		counter += maxHeapify(A, 0) + 1
	}

	A = A[:length]
	return
}
