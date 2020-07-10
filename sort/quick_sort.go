package sort

func quickSort(A []int) int {
	return innerQuickSort(A, 0, len(A)-1)
}

func innerQuickSort(A []int, p, r int) (counter int) {
	if p < r {
		q, c := partition(A, p, r)
		counter = c
		counter += innerQuickSort(A, p, q-1)
		counter += innerQuickSort(A, q+1, r)
	}
	return
}

func partition(A []int, p, r int) (q, counter int) {
	x := A[r]
	q = p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			q++
			exchange(A, q, j)
			counter++
		}
	}
	q++
	exchange(A, q, r)
	counter++
	return
}
