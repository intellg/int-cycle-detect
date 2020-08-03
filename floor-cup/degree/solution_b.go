package degree

import "math"

// ∑n=1~degree(∑m=0~cup(C(n,m)))
func InnerCalculateB(floor, cup int) (degree int) { // this function relies on the caller function to filter out the log2(Floor)
	// 1.1 Initialize sum
	sum := int(math.Pow(2, float64(cup))) - 1

	// 1.2 Further calculate sum
	for degree = cup; sum < floor; degree++ {
		sum += sumCompose(degree, cup)
	}
	return
}

// ∑m=0~cup(C(n, m))
func sumCompose(n, cup int) int {
	sum := 0
	half := (n + 1) / 2
	if cup <= half {
		for i := 0; i < cup; i++ {
			sum += compose(n, i)
		}
	} else {
		floatHalf := float64(n+1) / 2
		mirror := n + 1 - cup
		for i := 0; float64(i) < floatHalf; i++ {
			if i < mirror || i == half {
				sum += compose(n, i)
			} else {
				sum += compose(n, i) * 2
			}
		}
	}
	return sum
}

// C(n, m)
func compose(n, m int) int {
	// FIXME: Comment out below 2 blocks because they are checked outside of the function
	//// filter out invalid values
	//if m < 0 || n < 0 || m > n {
	//	return 0
	//}
	//
	//// prepare
	//if m > n/2 {
	//	m = n - m
	//}

	// calculate
	result := 1
	j := n
	for i := m; i > 0; i-- {
		result *= j
		j--
	}
	for i := m; i > 1; i-- {
		result /= i
	}

	counter++
	return result
}
